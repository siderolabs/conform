# syntax = docker/dockerfile-upstream:1.2.0-labs

# THIS FILE WAS AUTOMATICALLY GENERATED, PLEASE DO NOT EDIT.
#
# Generated on 2022-06-03T19:43:01Z by kres 65530e7.

ARG TOOLCHAIN

# cleaned up specs and compiled versions
FROM scratch AS generate

FROM ghcr.io/siderolabs/ca-certificates:v1.0.0 AS image-ca-certificates

FROM ghcr.io/siderolabs/fhs:v1.0.0 AS image-fhs

# runs markdownlint
FROM node:18.2.0-alpine AS lint-markdown
WORKDIR /src
RUN npm i -g markdownlint-cli@0.31.1
RUN npm i sentences-per-line@0.2.1
COPY .markdownlint.json .
COPY ./CHANGELOG.md ./CHANGELOG.md
COPY ./README.md ./README.md
RUN markdownlint --ignore "CHANGELOG.md" --ignore "**/node_modules/**" --ignore '**/hack/chglog/**' --rules node_modules/sentences-per-line/index.js .

# base toolchain image
FROM ${TOOLCHAIN} AS toolchain
RUN apk --update --no-cache add bash curl build-base protoc protobuf-dev git

# build tools
FROM toolchain AS tools
ENV GO111MODULE on
ENV CGO_ENABLED 0
ENV GOPATH /go
ARG GOLANGCILINT_VERSION
RUN curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/${GOLANGCILINT_VERSION}/install.sh | bash -s -- -b /bin ${GOLANGCILINT_VERSION}
ARG GOFUMPT_VERSION
RUN go install mvdan.cc/gofumpt@${GOFUMPT_VERSION} \
	&& mv /go/bin/gofumpt /bin/gofumpt
ARG GOIMPORTS_VERSION
RUN go install golang.org/x/tools/cmd/goimports@${GOIMPORTS_VERSION} \
	&& mv /go/bin/goimports /bin/goimports
ARG DEEPCOPY_VERSION
RUN go install github.com/siderolabs/deep-copy@${DEEPCOPY_VERSION} \
	&& mv /go/bin/deep-copy /bin/deep-copy

# tools and sources
FROM tools AS base
WORKDIR /src
COPY ./go.mod .
COPY ./go.sum .
RUN --mount=type=cache,target=/go/pkg go mod download
RUN --mount=type=cache,target=/go/pkg go mod verify
COPY ./cmd ./cmd
COPY ./internal ./internal
RUN --mount=type=cache,target=/go/pkg go list -mod=readonly all >/dev/null

# builds conform-darwin-amd64
FROM base AS conform-darwin-amd64-build
COPY --from=generate / /
WORKDIR /src/cmd/conform
ARG VERSION_PKG="github.com/siderolabs/conform/internal/version"
ARG SHA
ARG TAG
RUN --mount=type=cache,target=/root/.cache/go-build --mount=type=cache,target=/go/pkg GOARCH=amd64 GOOS=darwin go build -ldflags "-s -w -X ${VERSION_PKG}.Name=conform -X ${VERSION_PKG}.SHA=${SHA} -X ${VERSION_PKG}.Tag=${TAG}" -o /conform-darwin-amd64

# builds conform-darwin-arm64
FROM base AS conform-darwin-arm64-build
COPY --from=generate / /
WORKDIR /src/cmd/conform
ARG VERSION_PKG="github.com/siderolabs/conform/internal/version"
ARG SHA
ARG TAG
RUN --mount=type=cache,target=/root/.cache/go-build --mount=type=cache,target=/go/pkg GOARCH=arm64 GOOS=darwin go build -ldflags "-s -w -X ${VERSION_PKG}.Name=conform -X ${VERSION_PKG}.SHA=${SHA} -X ${VERSION_PKG}.Tag=${TAG}" -o /conform-darwin-arm64

# builds conform-linux-amd64
FROM base AS conform-linux-amd64-build
COPY --from=generate / /
WORKDIR /src/cmd/conform
ARG VERSION_PKG="github.com/siderolabs/conform/internal/version"
ARG SHA
ARG TAG
RUN --mount=type=cache,target=/root/.cache/go-build --mount=type=cache,target=/go/pkg GOARCH=amd64 GOOS=linux go build -ldflags "-s -w -X ${VERSION_PKG}.Name=conform -X ${VERSION_PKG}.SHA=${SHA} -X ${VERSION_PKG}.Tag=${TAG}" -o /conform-linux-amd64

# builds conform-linux-arm64
FROM base AS conform-linux-arm64-build
COPY --from=generate / /
WORKDIR /src/cmd/conform
ARG VERSION_PKG="github.com/siderolabs/conform/internal/version"
ARG SHA
ARG TAG
RUN --mount=type=cache,target=/root/.cache/go-build --mount=type=cache,target=/go/pkg GOARCH=arm64 GOOS=linux go build -ldflags "-s -w -X ${VERSION_PKG}.Name=conform -X ${VERSION_PKG}.SHA=${SHA} -X ${VERSION_PKG}.Tag=${TAG}" -o /conform-linux-arm64

# runs gofumpt
FROM base AS lint-gofumpt
RUN FILES="$(gofumpt -l .)" && test -z "${FILES}" || (echo -e "Source code is not formatted with 'gofumpt -w .':\n${FILES}"; exit 1)

# runs goimports
FROM base AS lint-goimports
RUN FILES="$(goimports -l -local github.com/siderolabs/conform .)" && test -z "${FILES}" || (echo -e "Source code is not formatted with 'goimports -w -local github.com/siderolabs/conform .':\n${FILES}"; exit 1)

# runs golangci-lint
FROM base AS lint-golangci-lint
COPY .golangci.yml .
ENV GOGC 50
RUN --mount=type=cache,target=/root/.cache/go-build --mount=type=cache,target=/root/.cache/golangci-lint --mount=type=cache,target=/go/pkg golangci-lint run --config .golangci.yml

# runs unit-tests with race detector
FROM base AS unit-tests-race
ARG TESTPKGS
RUN --mount=type=cache,target=/root/.cache/go-build --mount=type=cache,target=/go/pkg --mount=type=cache,target=/tmp CGO_ENABLED=1 go test -v -race -count 1 ${TESTPKGS}

# runs unit-tests
FROM base AS unit-tests-run
ARG TESTPKGS
RUN --mount=type=cache,target=/root/.cache/go-build --mount=type=cache,target=/go/pkg --mount=type=cache,target=/tmp go test -v -covermode=atomic -coverprofile=coverage.txt -coverpkg=${TESTPKGS} -count 1 ${TESTPKGS}

FROM scratch AS conform-darwin-amd64
COPY --from=conform-darwin-amd64-build /conform-darwin-amd64 /conform-darwin-amd64

FROM scratch AS conform-darwin-arm64
COPY --from=conform-darwin-arm64-build /conform-darwin-arm64 /conform-darwin-arm64

FROM scratch AS conform-linux-amd64
COPY --from=conform-linux-amd64-build /conform-linux-amd64 /conform-linux-amd64

FROM scratch AS conform-linux-arm64
COPY --from=conform-linux-arm64-build /conform-linux-arm64 /conform-linux-arm64

FROM scratch AS unit-tests
COPY --from=unit-tests-run /src/coverage.txt /coverage.txt

FROM conform-linux-${TARGETARCH} AS conform

FROM scratch AS image-conform
ARG TARGETARCH
COPY --from=conform conform-linux-${TARGETARCH} /conform
COPY --from=image-fhs / /
COPY --from=image-ca-certificates / /
LABEL org.opencontainers.image.source https://github.com/siderolabs/conform
ENTRYPOINT ["/conform"]

