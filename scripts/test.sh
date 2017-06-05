#!/bin/bash

set -e

GOPACKAGES=$(go list ./... | grep -v /vendor/ | grep -v /api)
GOFILES=$(find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./api/*")

COVERAGE_REPORT=coverage.txt
PROFILE=profile.out

if [[ -f ${COVERAGE_REPORT} ]]; then
  rm ${COVERAGE_REPORT}
fi

touch ${COVERAGE_REPORT}

echo "Running tests"
for package in ${GOPACKAGES[@]}; do
  go test -v -race -coverprofile=${PROFILE} -covermode=atomic $package
  if [ -f ${PROFILE} ]; then
    cat ${PROFILE} >> ${COVERAGE_REPORT}
    rm ${PROFILE}
  fi
done

echo "Vetting packages"
go vet ${GOPACKAGES}
if [ $? -eq 1 ]; then
  exit 1
fi

echo "Linting packages"
golint -set_exit_status ${GOPACKAGES}
if [ $? -eq 1 ]; then
  exit 1
fi

echo "Formatting go files"
if [ ! -z "$(gofmt -l -s ${GOFILES})" ]; then
  exit 1
fi

exit 0
