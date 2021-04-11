FROM golang:1.16

ENV GOPATH=/go
ENV GO111MODULE=on

RUN apt-get update -qq && apt-get install -yq default-mysql-client && \
  go get bitbucket.org/liamstask/goose/cmd/goose && \
  mkdir -p /go/src/github.com/tabakazu/golang-webapi-demo/db

WORKDIR /go/src/github.com/tabakazu/golang-webapi-demo
