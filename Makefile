VERSION := $(shell git describe --tags --always)
GITREV := $(shell git rev-parse --short HEAD)
GITBRANCH := $(shell git rev-parse --abbrev-ref HEAD)
DATE := $(shell LANG=US date +"%a, %d %b %Y %X %z")

GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/dist
GOARCH := $(ARCH)
GOENVVARS := GOBIN=$(GOBIN) CGO_ENABLED=1 GOOS=$(OS) GOARCH=$(GOARCH)
GOBINARY := lagrange-cli
GOCMD := $(GOBASE)/cmd/

LDFLAGS += -X 'github.com/Lagrange-Labs/client-cli.Version=$(VERSION)'
LDFLAGS += -X 'github.com/Lagrange-Labs/client-cli.GitRev=$(GITREV)'
LDFLAGS += -X 'github.com/Lagrange-Labs/client-cli.GitBranch=$(GITBRANCH)'
LDFLAGS += -X 'github.com/Lagrange-Labs/client-cli.BuildDate=$(DATE)'

STOP := docker compose down --remove-orphans

# Building the docker image and the binary
build: ## Builds the binary locally into ./dist
	$(GOENVVARS) go build -ldflags "all=$(LDFLAGS)" -o $(GOBIN)/$(GOBINARY) $(GOCMD)
.PHONY: build

# Useful and Test Scripts
scgen: # Generate the go bindings for the smart contracts
	@ cd scinterface && sh generator.sh

# Linting, Teseting, Benchmarking
golangci_lint_cmd=github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2

install-linter:
	@echo "--> Installing linter"
	@go install $(golangci_lint_cmd)

lint:
	@echo "--> Running linter"
	@ $$(go env GOPATH)/bin/golangci-lint run --timeout=10m
.PHONY:	lint install-linter

test:
	trap '$(STOP)' EXIT; go test ./... --timeout=10m
.PHONY: test

docker-build: ## Builds a docker image with the cli binary
	docker build -t lagrange-cli -f ./Dockerfile .
.PHONY: docker-build