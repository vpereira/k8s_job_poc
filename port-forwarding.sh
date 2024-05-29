#!/bin/bash

# Get the name of the Redis pod
POD_NAME=$(kubectl get pods -n redis-namespace -l app=redis -o jsonpath="{.items[0].metadata.name}")

# Forward port 6379 from the Redis pod to localhost
kubectl port-forward -n redis-namespace pod/$POD_NAME 6379:6379
