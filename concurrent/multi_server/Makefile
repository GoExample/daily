PROJECT_NAME := "multi_server"
PKG := "$(PROJECT_NAME)"
GO_FILES := $(shell find . -name '*.go' | grep -v _test.go)

.PHONY: all lint fmt vet compile build clean help

all: build

lint: ## Lint Golang files
	@golint -set_exit_status ${GO_FILES}

fmt: ## Format Golang files
	echo ${GO_FILES}
	@go fmt ${GO_FILES}

vet: ## Run go vet
	@go vet ${GO_FILES}

compile: ## Compile the binary file
	@for var in ${GO_FILES}; do \
        go build $$var; \
    done

build: compile ## Build and upx executable binary
	@find . -type f -executable | xargs upx

clean: ## Remove executable binary
	@find . -type f -executable | xargs rm -f

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'