build:
	docker-compose build
.PHONY: build

run: build
	docker-compose up app
.PHONY: run

test: build
	docker-compose run acceptance-test
	docker-compose down
.PHONY: test

clean:
	docker-compose down --rmi all --volumes --remove-orphans
.PHONY: clean
