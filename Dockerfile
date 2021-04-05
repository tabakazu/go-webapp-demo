FROM golang:1.16-alpine

ENV GOPATH=/go
ENV GO111MODULE=on

RUN apk add --update --no-cache bash ca-certificates g++ gcc git mysql-client postgresql-client
RUN mkdir -p /go/src/github.com/tabakazu/golang-webapi-demo
