VERSION=0.0.1

.PHONY: build
build:
	docker build --platform linux/amd64 --tag ghcr.io/freckie/deployer-mobilenetv2:$(VERSION) .
	docker push ghcr.io/freckie/deployer-mobilenetv2:$(VERSION)