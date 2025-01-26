build:
	docker-compose build app
.PHONY: build

run: build
	docker-compose up app
.PHONY: run

test-build:
	docker-compose build test
.PHONY: test-build

test: test-build
	docker-compose up --abort-on-container-exit test
	docker-compose down
.PHONY: test
