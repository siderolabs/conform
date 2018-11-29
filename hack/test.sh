#!/bin/bash

set -e

CGO_ENABLED=1
GOPACKAGES=$(go list ./...)

lint_packages() {
  if [ "${lint}" = true ]; then
    echo "linting packages"
    curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $GOPATH/bin v1.10.1
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
    go test -v ./...
  fi
}

coverage_tests() {
  if [ "${coverage}" = true ]; then
    echo "performing coverage tests"
    local coverage_report="../build/coverage.txt"
    local profile="../build/profile.out"
    if [[ -f ${coverage_report} ]]; then
      rm ${coverage_report}
    fi
    touch ${coverage_report}
    for package in ${GOPACKAGES[@]}; do
      go test -v -short -race -coverprofile=${profile} -covermode=atomic $package
      if [ -f ${profile} ]; then
        cat ${profile} >> ${coverage_report}
        rm ${profile}
      fi
    done
  fi
}

lint=false
short=false
tests=false
coverage=false

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
  --coverage)
  coverage=true
  ;;
  --all)
  lint=true
  short=true
  tests=true
  coverage=true
  ;;
  *)
  ;;
esac

go_test
coverage_tests
lint_packages

exit 0
