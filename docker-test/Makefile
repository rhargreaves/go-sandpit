build:
	docker-compose build
.PHONY: build

run: build
	docker-compose up app
.PHONY: run

test: build
	docker-compose up --abort-on-container-exit \
		acceptance-test
	docker-compose down
.PHONY: test
