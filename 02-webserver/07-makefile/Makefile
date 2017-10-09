PROJECT_PACKAGE := github.com/3fs/go-academy/02-webserver/07-makefile
DOCKER_IMAGE    := go-academy/02-webserver

# list of available packages
PKG_LIST_CMD := go list ./... | grep -v '/vendor/'
SOURCE_FILES := $(shell /usr/bin/find . -type f -name '*.go' -not -path './vendor/*')

.DEFAULT_GOAL := help
.PHONY: help
help: ## Print this text
	@grep -E '^[a-zA-Z_/%\-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

clean: ## Clean artifacts 
	rm -rf build


build: build/server build/cli ## Build all executables

build/server: ## Builds server executable
	go build -o build/server $(PROJECT_PACKAGE)/cmd/server

build/cli: ## Builds cli executable
	go build -o build/cli $(PROJECT_PACKAGE)/cmd/cli


run/server: build/server ## Run server executable
	./build/server

run/cli: build/cli ## Run cli executable
	./build/cli


qa: test/unit test/static ## Run entire QA suite

test/unit: ## Run unit tests
	go test -v $(shell $(PKG_LIST_CMD))

test/static: test/format test/lint test/vet ## Perform static analysis of the code

test/format: ## Test code formatting
	test -z "$(shell gofmt -l $(SOURCE_FILES))"

test/lint: ## Lint the source code
	@$(foreach pkg,$(shell $(PKG_LIST_CMD)),golint -set_exit_status $(pkg) || exit 1;)

test/vet: ## Vet the source code
	go vet $(shell $(PKG_LIST_CMD))


docker/build/builder: ## Build a builder Docker image
	docker build -t $(DOCKER_IMAGE):builder .

docker/publish/builder: ## Publish a builder Docker image
	docker push $(DOCKER_IMAGE):builder

builder/%:: ## Run make target in builder container
	docker run -it \
		-v "$(shell pwd)":/go/src/$(PROJECT_PACKAGE) \
		-w /go/src/$(PROJECT_PACKAGE) \
		-p 8080:8080 \
		$(DOCKER_IMAGE):builder make $*;

console: ## Run bash shell, i.e. builder/console
	@bash
