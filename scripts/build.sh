#!/bin/sh

cd twirp/lambda

# Pre0build Cleanup
rm main main.zip

# Build
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main *.go
zip main.zip main

# Post-Build
cd -

