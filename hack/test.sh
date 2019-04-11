#!/bin/bash

set -e

CGO_ENABLED=1

lint_packages() {
  if [ "${lint}" = true ]; then
    echo "linting packages"
    curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $GOPATH/bin v1.16.0
    golangci-lint run --config "${BASH_SOURCE%/*}/golangci-lint.yaml"
  fi
}

go_test() {
  if [ "${short}" = true ]; then
    echo "performing short tests"
    go test -v -short ./...
  fi

  if [ "${tests}" = true ]; then
    echo "performing tests"
    go test -v -race -covermode=atomic -coverprofile=/coverage.txt ./...
  fi
}

lint=false
short=false
tests=false

case $1 in
  --lint)
  lint=true
  ;;
  --short)
  short=true
  ;;
  --integration)
  tests=true
  ;;
  --all)
  lint=true
  short=true
  tests=true
  ;;
  *)
  ;;
esac

go_test
lint_packages

exit 0
