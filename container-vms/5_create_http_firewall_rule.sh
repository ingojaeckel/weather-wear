#!/bin/sh

gcloud compute firewall-rules create allow-http --description "Incoming http allowed." --allow tcp:80 --format json
