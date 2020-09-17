#!/bin/bash

## Generate Go code, OpenAPI/Swagger definition, and gRPC-Gateway code for REST/HTTP
protoc -I=./api/protobuf -I=./third_party \
  --go_out=plugins=grpc:pkg/service --go_opt=paths=source_relative \
  --grpc-gateway_out=logtostderr=true,paths=source_relative:pkg/service \
  --swagger_out=./api/openapi \
  api/protobuf/petstore.proto
