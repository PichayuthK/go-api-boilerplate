SHELL := /bin/bash

run:
	go run main.go

VERSION := 1.0

build_docker:
	docker build \
		-f dockerfile \
		-t app-boilerplate:$(VERSION) \
		--build-arg VCS_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.