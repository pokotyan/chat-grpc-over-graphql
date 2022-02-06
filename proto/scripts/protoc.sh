#!/bin/sh

set -xe

SERVER_OUTPUT_DIR=server/pb

protoc --version
protoc --proto_path=proto $(find proto -iname "*.proto") \
  --go_out=plugins="grpc:${SERVER_OUTPUT_DIR}"
