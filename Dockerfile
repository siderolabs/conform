# syntax = docker/dockerfile-upstream:1.23.0-labs

# THIS FILE WAS AUTOMATICALLY GENERATED, PLEASE DO NOT EDIT.
#
# Generated on 2026-04-28T10:33:12Z by kres e4dc583.

ARG TOOLCHAIN=scratch

FROM ghcr.io/siderolabs/ca-certificates:v1.13.0 AS image-ca-certificates

FROM ghcr.io/siderolabs/fhs:v1.13.0 AS image-fhs

# runs markdownlint
FROM docker.io/oven/bun:1.3.13-alpine AS lint-markdown
WORKDIR /src
RUN bun i markdownlint-cli@0.48.0 sentences-per-line@0.5.2
COPY .markdownlint.json .
COPY ./CHANGELOG.md ./CHANGELOG.md
COPY ./README.md ./README.md
RUN bunx markdownlint --ignore "CHANGELOG.md" --ignore "**/node_modules/**" --ignore '**/hack/chglog/**' --rules markdownlint-sentences-per-line .

# base toolchain image
FROM --platform=${BUILDPLATFORM} ${TOOLCHAIN} AS toolchain
RUN apk --update --no-cache add bash build-base curl jq protoc protobuf-dev git

# build tools
FROM --platform=${BUILDPLATFORM} toolchain AS tools
ENV GO111MODULE=on
ARG CGO_ENABLED
ENV CGO_ENABLED=${CGO_ENABLED}
ARG GOTOOLCHAIN
ENV GOTOOLCHAIN=${GOTOOLCHAIN}
ARG GOEXPERIMENT
ENV GOEXPERIMENT=${GOEXPERIMENT}
ENV GOPATH=/go
ARG DEEPCOPY_VERSION
RUN --mount=type=cache,target=/root/.cache/go-build,id=conform/root/.cache/go-build --mount=type=cache,target=/go/pkg,id=conform/go/pkg go install github.com/siderolabs/deep-copy@${DEEPCOPY_VERSION} \
	&& mv /go/bin/deep-copy /bin/deep-copy
ARG GOLANGCILINT_VERSION
RUN --mount=type=cache,target=/root/.cache/go-build,id=conform/root/.cache/go-build --mount=type=cache,target=/go/pkg,id=conform/go/pkg go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@${GOLANGCILINT_VERSION} \
	&& mv /go/bin/golangci-lint /bin/golangci-lint
RUN --mount=type=cache,target=/root/.cache/go-build,id=conform/root/.cache/go-build --mount=type=cache,target=/go/pkg,id=conform/go/pkg go install golang.org/x/vuln/cmd/govulncheck@latest \
	&& mv /go/bin/govulncheck /bin/govulncheck
ARG DIS_VULNCHECK_VERSION
RUN --mount=type=cache,target=/root/.cache/go-build,id=conform/root/.cache/go-build --mount=type=cache,target=/go/pkg,id=conform/go/pkg go install github.com/shanduur/dis-vulncheck@${DIS_VULNCHECK_VERSION} \
	&& mv /go/bin/dis-vulncheck /bin/dis-vulncheck
ARG GOFUMPT_VERSION
RUN go install mvdan.cc/gofumpt@${GOFUMPT_VERSION} \
	&& mv /go/bin/gofumpt /bin/gofumpt

# tools and sources
FROM tools AS base
WORKDIR /src
COPY go.mod go.mod
COPY go.sum go.sum
RUN cd .
RUN --mount=type=cache,target=/go/pkg,id=conform/go/pkg go mod download
RUN --mount=type=cache,target=/go/pkg,id=conform/go/pkg go mod verify
COPY ./cmd ./cmd
COPY ./internal ./internal
RUN --mount=type=cache,target=/go/pkg,id=conform/go/pkg go list -mod=readonly all >/dev/null

FROM tools AS embed-generate
ARG SHA
ARG TAG
WORKDIR /src
RUN mkdir -p internal/version/data && \
    echo -n ${SHA} > internal/version/data/sha && \
    echo -n ${TAG} > internal/version/data/tag

# runs gofumpt
FROM base AS lint-gofumpt
RUN FILES="$(gofumpt -l .)" && test -z "${FILES}" || (echo -e "Source code is not formatted with 'gofumpt -w .':\n${FILES}"; exit 1)

# runs golangci-lint
FROM base AS lint-golangci-lint
WORKDIR /src
COPY .golangci.yml .
ENV GOGC=50
RUN --mount=type=cache,target=/root/.cache/go-build,id=conform/root/.cache/go-build --mount=type=cache,target=/root/.cache/golangci-lint,id=conform/root/.cache/golangci-lint,sharing=locked --mount=type=cache,target=/go/pkg,id=conform/go/pkg golangci-lint run --config .golangci.yml

# runs golangci-lint fmt
FROM base AS lint-golangci-lint-fmt-run
WORKDIR /src
COPY .golangci.yml .
ENV GOGC=50
RUN --mount=type=cache,target=/root/.cache/go-build,id=conform/root/.cache/go-build --mount=type=cache,target=/root/.cache/golangci-lint,id=conform/root/.cache/golangci-lint,sharing=locked --mount=type=cache,target=/go/pkg,id=conform/go/pkg golangci-lint fmt --config .golangci.yml
RUN --mount=type=cache,target=/root/.cache/go-build,id=conform/root/.cache/go-build --mount=type=cache,target=/root/.cache/golangci-lint,id=conform/root/.cache/golangci-lint,sharing=locked --mount=type=cache,target=/go/pkg,id=conform/go/pkg golangci-lint run --fix --issues-exit-code 0 --config .golangci.yml

# runs govulncheck
FROM base AS lint-govulncheck
WORKDIR /src
RUN --mount=type=cache,target=/root/.cache/go-build,id=conform/root/.cache/go-build --mount=type=cache,target=/go/pkg,id=conform/go/pkg dis-vulncheck -tool=false ./...

# runs unit-tests with race detector
FROM base AS unit-tests-race
WORKDIR /src
ARG TESTPKGS
RUN --mount=type=cache,target=/root/.cache/go-build,id=conform/root/.cache/go-build --mount=type=cache,target=/go/pkg,id=conform/go/pkg --mount=type=cache,target=/tmp,id=conform/tmp CGO_ENABLED=1 go test -race ${TESTPKGS}

# runs unit-tests
FROM base AS unit-tests-run
WORKDIR /src
ARG TESTPKGS
RUN --mount=type=cache,target=/root/.cache/go-build,id=conform/root/.cache/go-build --mount=type=cache,target=/go/pkg,id=conform/go/pkg --mount=type=cache,target=/tmp,id=conform/tmp go test -covermode=atomic -coverprofile=coverage.txt -coverpkg=${TESTPKGS} ${TESTPKGS}

FROM embed-generate AS embed-abbrev-generate
WORKDIR /src
ARG ABBREV_TAG
RUN echo -n 'undefined' > internal/version/data/sha && \
    echo -n ${ABBREV_TAG} > internal/version/data/tag

# clean golangci-lint fmt output
FROM scratch AS lint-golangci-lint-fmt
COPY --from=lint-golangci-lint-fmt-run /src .

FROM scratch AS unit-tests
COPY --from=unit-tests-run /src/coverage.txt /coverage-unit-tests.txt

# cleaned up specs and compiled versions
FROM scratch AS generate
COPY --from=embed-abbrev-generate /src/internal/version internal/version

# builds conform-darwin-amd64
FROM base AS conform-darwin-amd64-build
COPY --from=generate / /
COPY --from=embed-generate / /
WORKDIR /src/cmd/conform
ARG GO_BUILDFLAGS
ARG GO_LDFLAGS
ARG VERSION_PKG="internal/version"
ARG SHA
ARG TAG
RUN --mount=type=cache,target=/root/.cache/go-build,id=conform/root/.cache/go-build --mount=type=cache,target=/go/pkg,id=conform/go/pkg GOARCH=amd64 GOOS=darwin go build ${GO_BUILDFLAGS} -ldflags "${GO_LDFLAGS} -X ${VERSION_PKG}.Name=conform -X ${VERSION_PKG}.SHA=${SHA} -X ${VERSION_PKG}.Tag=${TAG}" -o /conform-darwin-amd64

# builds conform-darwin-arm64
FROM base AS conform-darwin-arm64-build
COPY --from=generate / /
COPY --from=embed-generate / /
WORKDIR /src/cmd/conform
ARG GO_BUILDFLAGS
ARG GO_LDFLAGS
ARG VERSION_PKG="internal/version"
ARG SHA
ARG TAG
RUN --mount=type=cache,target=/root/.cache/go-build,id=conform/root/.cache/go-build --mount=type=cache,target=/go/pkg,id=conform/go/pkg GOARCH=arm64 GOOS=darwin go build ${GO_BUILDFLAGS} -ldflags "${GO_LDFLAGS} -X ${VERSION_PKG}.Name=conform -X ${VERSION_PKG}.SHA=${SHA} -X ${VERSION_PKG}.Tag=${TAG}" -o /conform-darwin-arm64

# builds conform-linux-amd64
FROM base AS conform-linux-amd64-build
COPY --from=generate / /
COPY --from=embed-generate / /
WORKDIR /src/cmd/conform
ARG GO_BUILDFLAGS
ARG GO_LDFLAGS
ARG VERSION_PKG="internal/version"
ARG SHA
ARG TAG
RUN --mount=type=cache,target=/root/.cache/go-build,id=conform/root/.cache/go-build --mount=type=cache,target=/go/pkg,id=conform/go/pkg GOARCH=amd64 GOOS=linux go build ${GO_BUILDFLAGS} -ldflags "${GO_LDFLAGS} -X ${VERSION_PKG}.Name=conform -X ${VERSION_PKG}.SHA=${SHA} -X ${VERSION_PKG}.Tag=${TAG}" -o /conform-linux-amd64

# builds conform-linux-arm64
FROM base AS conform-linux-arm64-build
COPY --from=generate / /
COPY --from=embed-generate / /
WORKDIR /src/cmd/conform
ARG GO_BUILDFLAGS
ARG GO_LDFLAGS
ARG VERSION_PKG="internal/version"
ARG SHA
ARG TAG
RUN --mount=type=cache,target=/root/.cache/go-build,id=conform/root/.cache/go-build --mount=type=cache,target=/go/pkg,id=conform/go/pkg GOARCH=arm64 GOOS=linux go build ${GO_BUILDFLAGS} -ldflags "${GO_LDFLAGS} -X ${VERSION_PKG}.Name=conform -X ${VERSION_PKG}.SHA=${SHA} -X ${VERSION_PKG}.Tag=${TAG}" -o /conform-linux-arm64

FROM scratch AS conform-darwin-amd64
COPY --from=conform-darwin-amd64-build /conform-darwin-amd64 /conform-darwin-amd64

FROM scratch AS conform-darwin-arm64
COPY --from=conform-darwin-arm64-build /conform-darwin-arm64 /conform-darwin-arm64

FROM scratch AS conform-linux-amd64
COPY --from=conform-linux-amd64-build /conform-linux-amd64 /conform-linux-amd64

FROM scratch AS conform-linux-arm64
COPY --from=conform-linux-arm64-build /conform-linux-arm64 /conform-linux-arm64

FROM conform-linux-${TARGETARCH} AS conform

FROM scratch AS conform-all
COPY --from=conform-darwin-amd64 / /
COPY --from=conform-darwin-arm64 / /
COPY --from=conform-linux-amd64 / /
COPY --from=conform-linux-arm64 / /

FROM scratch AS image-conform
ARG TARGETARCH
COPY --from=conform conform-linux-${TARGETARCH} /conform
COPY --from=image-fhs / /
COPY --from=image-ca-certificates / /
LABEL org.opencontainers.image.source=https://github.com/siderolabs/conform
ENTRYPOINT ["/conform"]

