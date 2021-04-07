FROM golang:1.16

ENV GOPATH=/go
ENV GO111MODULE=on

RUN mkdir -p /go/src/github.com/tabakazu/golang-webapi-demo
