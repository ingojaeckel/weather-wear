#!/bin/sh
VERSION=$1
gcloud docker push gcr.io/weather-wea/front-end:${VERSION} 
