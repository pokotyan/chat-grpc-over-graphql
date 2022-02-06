FROM golang:1.17.6

ENV GO111MODULE=on

WORKDIR /grpc-server

RUN apt-get -qq update && apt-get -qq install -y git

RUN go get github.com/pilu/fresh
