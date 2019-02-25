# dgoss
>A simple dgoss image for CI Jobs

Goss is tool for validating server’s configuration (avoid conf. drift). Dgoss is wrapper written on top of the goss for validating docker images.

https://github.com/aelsabbahy/goss/tree/master/extras/dgoss

docker build -t dgoss .

docker build -t app example

docker run --rm -p 8080:8080 app

docker run --rm -it \
  -v "$(pwd)/example":/src \
  -v /var/run/docker.sock:/var/run/docker.sock \
dgoss

dgoss edit app
goss a process app
goss a http http://localhost:8080
exit

dgoss run app

dgoss edit app
goss a user root
goss a user node

goss a file /app
goss a file /app/app
