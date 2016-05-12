# Kubernetes introduction

This is a summary of the Kubernetes commands needed to deploy and run a simple app like this one to GKE as described on [http://kubernetes.io/docs/hellonode/](http://kubernetes.io/docs/hellonode/).

## Building and running app locally

    $ docker build -t gcr.io/weather-wea/hello-node:v1 .
    $ docker run -d -p 8080:8080 gcr.io/weather-wea/hello-node:v1
    $ docker ps
    $ docker stop <container id>

## Uploading container and deploying it to GKE

    $ gcloud docker push gcr.io/weather-wea/hello-node:v1
    $ gcloud container clusters get-credentials hello-world --zone us-east1-c
    $ kubectl run hello-node --image=gcr.io/weather-wea/hello-node:v1 --port=8080
    $ kubectl get deployments

## View cluster status

    $ kubectl get pods
    $ kubectl logs <POD-NAME>
    $ kubectl cluster-info
    $ kubectl get events
    $ kubectl config view

## Add load balancer

    $ kubectl expose deployment hello-node --type="LoadBalancer"

It will take a couple minutes until the external IP address shows up.

    $ kubectl get services hello-node
    $ curl http://EXTERNAL_IP:8080
    $ kubectl scale deployment hello-node --replicas=4

## Updating the existing app

    $ docker build -t gcr.io/weather-wea/hello-node:v2 .
    $ gcloud docker push gcr.io/weather-wea/hello-node:v2

# Update app config by increasing version of hello-node Docker container to v2

    $ kubectl edit deployment hello-node
    $ kubectl get deployments
    $ kubectl delete service,deployment hello-node
    $ gcloud container clusters delete hello-world --zone us-east1-c
    $ gsutil rm -r gs://artifacts.weather-wea.appspot.com/
