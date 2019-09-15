VERSION := 0.2.0
BIN_FILE_NAME := goexpand
DOCKER_CONTAINER_NAME := go-build

.PHONY: build
build:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/$(BIN_FILE_NAME) -ldflags "-X main.version=$(VERSION)" cmd/expand/main.go
	GOOS=linux GOARCH=amd64 go build -o ./bin/$(BIN_FILE_NAME)-linux -ldflags "-X main.version=$(VERSION)" cmd/expand/main.go
	chmod 755 bin/*

.PHONY: test
test:
	go test -v -cover

.PHONY: tag
tag:
	git tag $(VERSION)
	git push origin $(VERSION)

.PHONY: docker-build
docker-build:
	docker build -t $(DOCKER_CONTAINER_NAME) .

.PHONY: docker-run
docker-run:
	docker run -it -v $$(pwd)/bin:/go/src/app/bin --rm $(DOCKER_CONTAINER_NAME) make build