IMAGE_NAME=docker-test

build:
	docker build -t $(IMAGE_NAME) .
.PHONY: build

run: build
	docker run -p 8080:8080 $(IMAGE_NAME)
.PHONY: run
