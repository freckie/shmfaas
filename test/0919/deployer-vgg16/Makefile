VERSION=0.3.4

.PHONY: build
build:
	docker build --platform linux/amd64 --tag ghcr.io/freckie/deployer-vgg16:$(VERSION) .
	docker push ghcr.io/freckie/deployer-vgg16:$(VERSION)