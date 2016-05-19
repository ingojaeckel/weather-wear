#!/bin/sh
gcloud compute instances create weather-wear-container-vm-test-1 \
    --image gci-beta-51-8172-26-0 \
    --image-project google-containers \
    --zone us-central1-a \
    --machine-type f1-micro

echo "Stop it via:"
echo "gcloud compute instances stop weather-wear-container-vm-test-1 --zone us-central1-a"

