#!/bin/sh
gcloud compute instances create weather-wear-container-vm-test-1 \
    --image gci-beta-51-8172-26-0 \
    --image-project google-containers \
    --metadata-from-file "google-container-manifest=containers.yaml,user-data=cloud-config.yaml" \
    --zone us-central1-a \
    --machine-type f1-micro

