#!/bin/bash

set -e

GOFILES=$(find . -type f -name '*.go' -not -path "./vendor/*")

echo "Linting packages"
gometalinter --vendor --disable=gas --disable=gotype --sort=path --deadline=600s ./...

echo "Formatting go files"
GOFMTFILES="$(gofmt -l -d -s ${GOFILES})"
if [ ! -z "${GOFMTFILES}" ]; then
  echo -e "Failed gofmt files:\n${GOFMTFILES}"
  exit 1
fi

exit 0
