#!/bin/sh
# Builds the app for a Linux docker container
CGO_ENABLED=0 GOOS=linux go build -ldflags "-s" -a -installsuffix cgo
