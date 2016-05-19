#!/bin/sh

./buildForDocker.sh
gcloud preview app deploy app.yaml --promote --version 1
