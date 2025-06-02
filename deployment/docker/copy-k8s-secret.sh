#!/bin/bash

# copy k8s secret example files to actual files
cp ../k8s/secret/broker-secret.example.yml ../k8s/secret/broker-secret.yml
cp ../k8s/secret/product-service-secret.example.yml ../k8s/secret/product-service-secret.yml
cp ../k8s/secret/score-service-secret.example.yml ../k8s/secret/score-service-secret.yml
cp ../k8s/secret/widget-service-secret.example.yml ../k8s/secret/widget-service-secret.yml