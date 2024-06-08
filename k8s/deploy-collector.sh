#!/bin/sh
envsubst < k8s/collector.yml | kubectl apply -n qualitytrace -f -
