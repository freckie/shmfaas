VERSION=0.1.0

.PHONY: build
build:
	docker build --platform linux/amd64 --tag ghcr.io/freckie/deployer-resnet50:$(VERSION) .
	docker push ghcr.io/freckie/deployer-resnet50:$(VERSION)