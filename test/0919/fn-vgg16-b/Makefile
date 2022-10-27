VERSION=b-0.1.0

.PHONY: build
build:
	docker build --platform linux/amd64 --tag ghcr.io/freckie/fn-vgg16:$(VERSION) .
	docker push ghcr.io/freckie/fn-vgg16:$(VERSION)

.PHONY: squash-build
squash-build:
	docker build --platform linux/amd64 --squash --tag ghcr.io/freckie/fn-vgg16:$(VERSION)-squash .
	docker push ghcr.io/freckie/fn-vgg16:$(VERSION)-squash
