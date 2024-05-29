# Define variables for image names and registry
REGISTRY := localhost:5000
PRODUCER_IMAGE := k8s_job_producer
CONSUMER_IMAGE := k8s_job_consumer
WEBUI_IMAGE := k8s_job_webui

# Target to build Docker images
build-images:
	docker build -t $(PRODUCER_IMAGE) -f Dockerfile.producer .
	docker build -t $(CONSUMER_IMAGE) -f Dockerfile.consumer .
	docker build -t $(WEBUI_IMAGE) -f Dockerfile.webui .

# Target to tag and push Docker images to local registry
push-images: build-images
	docker tag $(PRODUCER_IMAGE) $(REGISTRY)/$(PRODUCER_IMAGE)
	docker tag $(CONSUMER_IMAGE) $(REGISTRY)/$(CONSUMER_IMAGE)
	docker tag $(WEBUI_IMAGE) $(REGISTRY)/$(WEBUI_IMAGE)
	docker push $(REGISTRY)/$(PRODUCER_IMAGE)
	docker push $(REGISTRY)/$(CONSUMER_IMAGE)
	docker push $(REGISTRY)/$(WEBUI_IMAGE)

# Optional: clean up local images
clean:
	docker rmi $(PRODUCER_IMAGE) $(CONSUMER_IMAGE) $(WEBUI_IMAGE)
	docker rmi $(REGISTRY)/$(PRODUCER_IMAGE) $(REGISTRY)/$(CONSUMER_IMAGE) $(REGISTRY)/$(WEBUI_IMAGE)
