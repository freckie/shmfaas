VERSION=a-0.0.1

.PHONY: build
build:
	docker build --platform linux/amd64 --tag ghcr.io/freckie/fn-mobilenetv2:$(VERSION) .
	docker push ghcr.io/freckie/fn-mobilenetv2:$(VERSION)

.PHONY: squash-build
squash-build:
	docker build --platform linux/amd64 --squash --tag ghcr.io/freckie/fn-mobilenetv2:$(VERSION)-squash .
	docker push ghcr.io/freckie/fn-mobilenetv2:$(VERSION)-squash
