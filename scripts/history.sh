#!/bin/bash

# Initialise Git
git init
git remote add origin git@github.com:f-go/petstore.git
git pull origin master

# Initialize Go Modules
go mod init github.com/f-go/petstore
