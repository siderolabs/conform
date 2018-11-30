REPO ?= docker.io/autonomy
EXECUTOR ?= gcr.io/kaniko-project/executor
EXECUTOR_TAG ?= v0.6.0
WARMER ?= gcr.io/kaniko-project/warmer
WARMER_TAG ?= v0.6.0
GOLANG_IMAGE ?= golang:1.11.2
AUTH_CONFIG ?= $(HOME)/.kaniko/config.json

SHA := $(shell gitmeta git sha)
TAG := $(shell gitmeta image tag)
BUILT := $(shell gitmeta built)

EXECUTOR_ARGS := --context=/workspace --cache=true --cache-dir=/cache --cleanup
EXECUTOR_VOLUMES := --volume $(AUTH_CONFIG):/kaniko/.docker/config.json:ro --volume $(PWD)/cache:/cache --volume $(PWD)/build:/build

all: enforce clean conform

enforce:
	conform enforce

conform: cache
	docker run \
		--rm \
		$(EXECUTOR_VOLUMES) \
		--volume $(PWD):/workspace \
		$(EXECUTOR):$(EXECUTOR_TAG) \
			$(EXECUTOR_ARGS) \
			--dockerfile=Dockerfile \
			--cache-repo=$(REPO)/$@ \
			--destination=$(REPO)/$@:$(TAG) \
			--single-snapshot \
			--no-push \
			--build-arg GOLANG_IMAGE=$(GOLANG_IMAGE) \
			--build-arg SHA=$(SHA) \
			--build-arg TAG=$(TAG) \
			--build-arg BUILT="$(BUILT)"

.PHONY: cache
cache:
	docker run \
		--rm \
		$(EXECUTOR_VOLUMES) \
		$(WARMER):$(WARMER_TAG) \
			--cache-dir=/cache \
			--image=$(GOLANG_IMAGE)

debug:
	docker run \
		--rm \
		-it \
		$(EXECUTOR_VOLUMES) \
		--volume $(PWD):/workspace \
		--entrypoint=/busybox/sh \
		$(EXECUTOR):debug-${EXECUTOR_TAG}

clean:
	rm -rf ./build
