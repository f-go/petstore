#!/bin/bash

## Generate Go code
protoc --proto_path=api/protobuf --proto_path=third_party \
  --go_out=plugins=grpc:pkg/petstore --go_opt=paths=source_relative \
  api/protobuf/petstore.proto
  
## Generate OpenAPI/Swagger definition
protoc --proto_path=api/protobuf --proto_path=third_party \
  --swagger_out=./api/openapi \
  api/protobuf/petstore.proto
