#!/bin/bash

# Get the name of the Redis pod
POD_NAME=$(kubectl get pods -n redis-namespace -l app=redis -o jsonpath="{.items[0].metadata.name}")

# Forward port 6379 from the Redis pod to localhost
kubectl port-forward -n redis-namespace pod/$POD_NAME 6379:6379 &



# Get the name of the registry pod
POD_NAME_R=$(kubectl get pods -n redis-namespace -l app=registry -o jsonpath="{.items[0].metadata.name}")

# Forward port 5000 from the registry pod to localhost
kubectl port-forward -n redis-namespace pod/$POD_NAME_R 5000:5000


# Get the name of the webui pod
POD_NAME_W=$(kubectl get pods -n redis-namespace -l app=webui -o jsonpath="{.items[0].metadata.name}")

# Forward port 6379 from the Redis pod to localhost
kubectl port-forward -n redis-namespace pod/$POD_NAME_W 8080:8080