# add a build step and a run step for the docker container
build:
	docker build -t docker-test .
.PHONY: build

run: build
	docker run -p 8080:8080 docker-test
.PHONY: run
