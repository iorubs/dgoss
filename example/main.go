package main

import (
  "context"
  "flag"
  "fmt"
  "log"
  "net/http"
  "os"
  "os/signal"
  "time"
)

var listenAddr string

func main() {
  flag.StringVar(
    &listenAddr,
    "listen-addr",
    "0.0.0.0:8080",
    "server listen address",
  )
  flag.Parse()

  logger := log.New(os.Stdout, "http: ", log.LstdFlags)
  logger.Println("Starting server...")

  router := http.NewServeMux()
  router.Handle("/", index())

  server := &http.Server{
    Addr: listenAddr,
    Handler: logging(logger)(router),
    ErrorLog: logger,
    ReadTimeout: 5 * time.Second,
    WriteTimeout: 10 * time.Second,
    IdleTimeout: 15 * time.Second,
  }

  done := make(chan bool)
  quit := make(chan os.Signal, 1)
  signal.Notify(quit, os.Interrupt)

  go func() {
    <-quit
    logger.Println("Shutting down server...")

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    server.SetKeepAlivesEnabled(false)
    if err := server.Shutdown(ctx); err != nil {
      logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
    }
    close(done)
  }()

  logger.Println("Server listening on", listenAddr)
  if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
    logger.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
  }

  <-done
  logger.Println("Stopped server.")
}

func index() http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
      http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
      return
    }
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    w.Header().Set("X-Content-Type-Options", "nosniff")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "Hello, World!")
  })
}

func logging(logger *log.Logger) func(http.Handler) http.Handler {
  return func(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
      defer func() {
        logger.Println(
          r.Method,
          r.URL.Path,
          r.RemoteAddr,
          r.UserAgent(),
        )
      }()
      next.ServeHTTP(w, r)
    })
  }
}
