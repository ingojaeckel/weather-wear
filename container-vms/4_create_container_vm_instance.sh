#!/bin/sh
gcloud compute instances create weather-wear-container-vm-test-1 \
    --image container-vm \
    --metadata-from-file "google-container-manifest=containers.yaml" \
    --zone us-central1-a \
    --machine-type f1-micro

