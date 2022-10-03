VERSION=c-0.0.1

.PHONY: build
build:
	docker build --platform linux/amd64 --tag ghcr.io/freckie/fn-resnet50:$(VERSION) .
	docker push ghcr.io/freckie/fn-resnet50:$(VERSION)

.PHONY: squash-build
squash-build:
	docker build --platform linux/amd64 --squash --tag ghcr.io/freckie/fn-resnet50:$(VERSION)-squash .
	docker push ghcr.io/freckie/fn-resnet50:$(VERSION)-squash
