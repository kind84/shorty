GOVERSION ?= 1.15
default: build

build:
	go build ./...

build-dependencies:
	docker build -t shorty/dependencies --build-arg GOVERSION=$(GOVERSION) -f ./dependencies.Dockerfile .

docker: build-dependencies
	docker build -t shorty .

docker-build: build-dependencies
	docker-compose build

docker-run: docker-build
	docker-compose up
