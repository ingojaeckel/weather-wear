#!/bin/sh

# Builds the app for a Linux docker container
CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo

# Build and run docker container
docker build -t docker-test .
docker run -i -t -p 8080:8080 docker-test
