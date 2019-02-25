FROM golang:1.11 as builder

LABEL maintainer "ruben.vasconcelos3@mail.dcu.ie"

WORKDIR /go/src/example/

COPY main.go .
RUN go get \
    && CGO_ENABLED=0 go build -o app .

FROM alpine

RUN apk upgrade \
    && apk add --no-cache ca-certificates

WORKDIR /src
COPY --from=builder /go/src/example/app .

CMD ["./app"]
