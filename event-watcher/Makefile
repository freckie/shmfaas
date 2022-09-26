VERSION=0.2.0

.PHONY: build
build:
	docker build --platform linux/amd64 --tag ghcr.io/freckie/k8s-event-watcher:$(VERSION) .
	docker push ghcr.io/freckie/k8s-event-watcher:$(VERSION)