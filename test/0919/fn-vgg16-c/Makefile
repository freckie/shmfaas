VERSION=c-0.0.6

.PHONY: build
build:
	docker build --platform linux/amd64 --tag ghcr.io/freckie/fn-vgg16:$(VERSION) .
	docker push ghcr.io/freckie/fn-vgg16:$(VERSION)