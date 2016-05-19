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

## Update app config by increasing version of hello-node Docker container to v2

    $ kubectl edit deployment hello-node
    $ kubectl get deployments
    $ kubectl delete service,deployment hello-node
    $ gcloud container clusters delete hello-world --zone us-east1-c
    $ gsutil rm -r gs://artifacts.weather-wea.appspot.com/

## Auto scaling

Details on [http://kubernetes.io/docs/user-guide/horizontal-pod-autoscaling/](http://kubernetes.io/docs/user-guide/horizontal-pod-autoscaling/).

Example: *Running*

    $ kubectl autoscale rc foo --min=2 --max=5 --cpu-percent=80

*will create an autoscaler for replication controller foo, with target CPU utilization set to 80% and the number of replicas between 2 and 5*

See [http://kubernetes.io/docs/user-guide/kubectl/kubectl_autoscale/](http://kubernetes.io/docs/user-guide/kubectl/kubectl_autoscale/) for more information.

    $ kubectl get hpa
    $ kubectl describe hpa
    $ kubectl delete hpa

## Replication controllers

Details are on [http://kubernetes.io/docs/user-guide/replication-controller/](http://kubernetes.io/docs/user-guide/replication-controller/): *A replication controller ensures that a specified number of pod “replicas” are running at any one time.*
