#!/bin/bash

## Generate Go code
protoc -I=./api/protobuf -I=./third_party \
  --go_out=plugins=grpc:pkg/petstore --go_opt=paths=source_relative \
  api/protobuf/petstore.proto

## Generate OpenAPI/Swagger definition
protoc -I=./api/protobuf -I=./third_party \
  --swagger_out=./api/openapi \
  api/protobuf/petstore.proto
  
## Generate gRPC-Gateway code for REST/HTTP
protoc -I=./api/protobuf -I=./third_party \
  --grpc-gateway_out=logtostderr=true,paths=source_relative:pkg/petstore \
  api/protobuf/petstore.proto
