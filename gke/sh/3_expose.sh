#!/bin/sh

kubectl expose pod weather-wear-pod --type="LoadBalancer"
kubectl get services weather-wear-pod
