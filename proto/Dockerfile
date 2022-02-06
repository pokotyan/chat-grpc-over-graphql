FROM golang:1.17.6

ARG GRPCWEB_VERSION=1.3.1

WORKDIR /proto

RUN apt-get -qq update && apt-get -qq install -y \
  unzip protobuf-compiler

RUN curl -sSL https://github.com/grpc/grpc-web/releases/download/${GRPCWEB_VERSION}/protoc-gen-grpc-web-${GRPCWEB_VERSION}-linux-x86_64 -o /usr/local/bin/protoc-gen-grpc-web && \
  chmod +x /usr/local/bin/protoc-gen-grpc-web

RUN go get -u github.com/golang/protobuf/protoc-gen-go