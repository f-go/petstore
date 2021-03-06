#!/bin/bash

# Initialise Git
git init
git remote add origin git@github.com:f-go/petstore.git
git pull origin master

# Initialize Go Modules
go mod init github.com/f-go/petstore

# Create required directories
mkdir -p api/protobuf
mkdir -p api/openapi
mkdir -p pkg/service

# Install Go required packages
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u google.golang.org/grpc
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
