FROM golang:1.16

ENV GOPATH=/go
ENV GO111MODULE=on

RUN apt-get update -qq && apt-get install -yq default-mysql-client
RUN mkdir -p /go/src/github.com/tabakazu/golang-webapi-demo
