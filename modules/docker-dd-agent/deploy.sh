#!/bin/sh
# Build the golang app
./build.sh
gcloud preview app deploy app.yaml --promote --version 1
