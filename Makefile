# THIS FILE WAS AUTOMATICALLY GENERATED, PLEASE DO NOT EDIT.
#
# Generated on 2022-09-19T13:49:50Z by kres 9ea8a33.

# common variables

SHA := $(shell git describe --match=none --always --abbrev=8 --dirty)
TAG := $(shell git describe --tag --always --dirty)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
ARTIFACTS := _out
REGISTRY ?= ghcr.io
USERNAME ?= siderolabs
REGISTRY_AND_USERNAME ?= $(REGISTRY)/$(USERNAME)
GOLANGCILINT_VERSION ?= v1.49.0
GOFUMPT_VERSION ?= v0.3.1
GO_VERSION ?= 1.19
GOIMPORTS_VERSION ?= v0.1.12
PROTOBUF_GO_VERSION ?= 1.28.1
GRPC_GO_VERSION ?= 1.2.0
GRPC_GATEWAY_VERSION ?= 2.11.3
VTPROTOBUF_VERSION ?= 0.3.0
DEEPCOPY_VERSION ?= v0.5.5
TESTPKGS ?= ./...
KRES_IMAGE ?= ghcr.io/siderolabs/kres:latest
CONFORMANCE_IMAGE ?= ghcr.io/siderolabs/conform:latest

# docker build settings

BUILD := docker buildx build
PLATFORM ?= linux/amd64
PROGRESS ?= auto
PUSH ?= false
CI_ARGS ?=
COMMON_ARGS = --file=Dockerfile
COMMON_ARGS += --progress=$(PROGRESS)
COMMON_ARGS += --platform=$(PLATFORM)
COMMON_ARGS += --push=$(PUSH)
COMMON_ARGS += --build-arg=ARTIFACTS=$(ARTIFACTS)
COMMON_ARGS += --build-arg=SHA=$(SHA)
COMMON_ARGS += --build-arg=TAG=$(TAG)
COMMON_ARGS += --build-arg=USERNAME=$(USERNAME)
COMMON_ARGS += --build-arg=REGISTRY=$(REGISTRY)
COMMON_ARGS += --build-arg=TOOLCHAIN=$(TOOLCHAIN)
COMMON_ARGS += --build-arg=GOLANGCILINT_VERSION=$(GOLANGCILINT_VERSION)
COMMON_ARGS += --build-arg=GOFUMPT_VERSION=$(GOFUMPT_VERSION)
COMMON_ARGS += --build-arg=GOIMPORTS_VERSION=$(GOIMPORTS_VERSION)
COMMON_ARGS += --build-arg=PROTOBUF_GO_VERSION=$(PROTOBUF_GO_VERSION)
COMMON_ARGS += --build-arg=GRPC_GO_VERSION=$(GRPC_GO_VERSION)
COMMON_ARGS += --build-arg=GRPC_GATEWAY_VERSION=$(GRPC_GATEWAY_VERSION)
COMMON_ARGS += --build-arg=VTPROTOBUF_VERSION=$(VTPROTOBUF_VERSION)
COMMON_ARGS += --build-arg=DEEPCOPY_VERSION=$(DEEPCOPY_VERSION)
COMMON_ARGS += --build-arg=TESTPKGS=$(TESTPKGS)
TOOLCHAIN ?= docker.io/golang:1.19-alpine

# help menu

export define HELP_MENU_HEADER
# Getting Started

To build this project, you must have the following installed:

- git
- make
- docker (19.03 or higher)

## Creating a Builder Instance

The build process makes use of experimental Docker features (buildx).
To enable experimental features, add 'experimental: "true"' to '/etc/docker/daemon.json' on
Linux or enable experimental features in Docker GUI for Windows or Mac.

To create a builder instance, run:

	docker buildx create --name local --use


If you already have a compatible builder instance, you may use that instead.

## Artifacts

All artifacts will be output to ./$(ARTIFACTS). Images will be tagged with the
registry "$(REGISTRY)", username "$(USERNAME)", and a dynamic tag (e.g. $(IMAGE):$(TAG)).
The registry and username can be overridden by exporting REGISTRY, and USERNAME
respectively.

endef

all: unit-tests conform image-conform lint

.PHONY: clean
clean:  ## Cleans up all artifacts.
	@rm -rf $(ARTIFACTS)

target-%:  ## Builds the specified target defined in the Dockerfile. The build result will only remain in the build cache.
	@$(BUILD) --target=$* $(COMMON_ARGS) $(TARGET_ARGS) $(CI_ARGS) .

local-%:  ## Builds the specified target defined in the Dockerfile using the local output type. The build result will be output to the specified local destination.
	@$(MAKE) target-$* TARGET_ARGS="--output=type=local,dest=$(DEST) $(TARGET_ARGS)"

lint-golangci-lint:  ## Runs golangci-lint linter.
	@$(MAKE) target-$@

lint-gofumpt:  ## Runs gofumpt linter.
	@$(MAKE) target-$@

.PHONY: fmt
fmt:  ## Formats the source code
	@docker run --rm -it -v $(PWD):/src -w /src golang:$(GO_VERSION) \
		bash -c "export GO111MODULE=on; export GOPROXY=https://proxy.golang.org; \
		go install mvdan.cc/gofumpt@$(GOFUMPT_VERSION) && \
		gofumpt -w ."

lint-goimports:  ## Runs goimports linter.
	@$(MAKE) target-$@

.PHONY: base
base:  ## Prepare base toolchain
	@$(MAKE) target-$@

.PHONY: unit-tests
unit-tests:  ## Performs unit tests
	@$(MAKE) local-$@ DEST=$(ARTIFACTS)

.PHONY: unit-tests-race
unit-tests-race:  ## Performs unit tests with race detection enabled.
	@$(MAKE) target-$@

.PHONY: coverage
coverage:  ## Upload coverage data to codecov.io.
	bash -c "bash <(curl -s https://codecov.io/bash) -f $(ARTIFACTS)/coverage.txt -X fix"

.PHONY: $(ARTIFACTS)/conform-darwin-amd64
$(ARTIFACTS)/conform-darwin-amd64:
	@$(MAKE) local-conform-darwin-amd64 DEST=$(ARTIFACTS)

.PHONY: conform-darwin-amd64
conform-darwin-amd64: $(ARTIFACTS)/conform-darwin-amd64  ## Builds executable for conform-darwin-amd64.

.PHONY: $(ARTIFACTS)/conform-darwin-arm64
$(ARTIFACTS)/conform-darwin-arm64:
	@$(MAKE) local-conform-darwin-arm64 DEST=$(ARTIFACTS)

.PHONY: conform-darwin-arm64
conform-darwin-arm64: $(ARTIFACTS)/conform-darwin-arm64  ## Builds executable for conform-darwin-arm64.

.PHONY: $(ARTIFACTS)/conform-linux-amd64
$(ARTIFACTS)/conform-linux-amd64:
	@$(MAKE) local-conform-linux-amd64 DEST=$(ARTIFACTS)

.PHONY: conform-linux-amd64
conform-linux-amd64: $(ARTIFACTS)/conform-linux-amd64  ## Builds executable for conform-linux-amd64.

.PHONY: $(ARTIFACTS)/conform-linux-arm64
$(ARTIFACTS)/conform-linux-arm64:
	@$(MAKE) local-conform-linux-arm64 DEST=$(ARTIFACTS)

.PHONY: conform-linux-arm64
conform-linux-arm64: $(ARTIFACTS)/conform-linux-arm64  ## Builds executable for conform-linux-arm64.

.PHONY: conform
conform: conform-darwin-amd64 conform-darwin-arm64 conform-linux-amd64 conform-linux-arm64  ## Builds executables for conform.

.PHONY: lint-markdown
lint-markdown:  ## Runs markdownlint.
	@$(MAKE) target-$@

.PHONY: lint
lint: lint-golangci-lint lint-gofumpt lint-goimports lint-markdown  ## Run all linters for the project.

.PHONY: image-conform
image-conform:  ## Builds image for conform.
	@$(MAKE) target-$@ TARGET_ARGS="--tag=$(REGISTRY)/$(USERNAME)/conform:$(TAG)"

.PHONY: rekres
rekres:
	@docker pull $(KRES_IMAGE)
	@docker run --rm -v $(PWD):/src -w /src -e GITHUB_TOKEN $(KRES_IMAGE)

.PHONY: help
help:  ## This help menu.
	@echo "$$HELP_MENU_HEADER"
	@grep -E '^[a-zA-Z%_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: release-notes
release-notes:
	mkdir -p $(ARTIFACTS)
	@ARTIFACTS=$(ARTIFACTS) ./hack/release.sh $@ $(ARTIFACTS)/RELEASE_NOTES.md $(TAG)

.PHONY: conformance
conformance:
	@docker pull $(CONFORMANCE_IMAGE)
	@docker run --rm -it -v $(PWD):/src -w /src $(CONFORMANCE_IMAGE) enforce

