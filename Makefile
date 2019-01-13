SHA := $(shell gitmeta git sha)
TAG := $(shell gitmeta image tag)
BUILT := $(shell gitmeta built)

GOLANG_IMAGE ?= golang:1.11.4

COMMON_ARGS := -f ./Dockerfile --build-arg GOLANG_IMAGE=$(GOLANG_IMAGE) --build-arg SHA=$(SHA) --build-arg TAG=$(TAG) --build-arg BUILT="$(BUILT)" .

export DOCKER_BUILDKIT := 1

all: enforce build test image

enforce:
	@conform enforce

.PHONY: build
build:
	@docker build \
		-t conform/$@:$(SHA) \
		--target=$@ \
		$(COMMON_ARGS)
	@docker run --rm -it -v $(PWD)/build:/build conform/$@:$(SHA) cp /conform-linux-amd64 /build
	@docker run --rm -it -v $(PWD)/build:/build conform/$@:$(SHA) cp /conform-darwin-amd64 /build

test:
	@docker build \
		-t conform/$@:$(SHA) \
		--target=$@ \
		$(COMMON_ARGS)
	@docker run --rm -it -v $(PWD)/build:/build conform/$@:$(SHA) cp /coverage.txt /build

image: build
	@docker build \
		-t autonomy/conform:$(SHA) \
		--target=$@ \
		$(COMMON_ARGS)

push: image
	@docker tag autonomy/conform:$(SHA) autonomy/conform:latest
	@docker push autonomy/conform:$(SHA)
	@docker push autonomy/conform:latest

deps:
	@GO111MODULES=on CGO_ENABLED=0 go get -u github.com/autonomy/gitmeta
	@GO111MODULES=on CGO_ENABLED=0 go get -u github.com/autonomy/conform

clean:
	go clean -modcache
	rm -rf build vendor
