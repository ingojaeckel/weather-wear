#!/bin/sh

./buildForDocker.sh

# Build and run docker container
docker build -t docker-test .
docker run -i -t -p 8080:8080 docker-test
