#!/bin/bash

set -e

GOPACKAGES=$(go list ./... | grep -v /vendor/)
COVERAGE_REPORT=coverage.txt
PROFILE=profile.out

echo "Running tests"
if [[ -f ${COVERAGE_REPORT} ]]; then
  rm ${COVERAGE_REPORT}
fi
touch ${COVERAGE_REPORT}
for package in ${GOPACKAGES[@]}; do
  go test -v -race -coverprofile=${PROFILE} -covermode=atomic $package
  if [ -f ${PROFILE} ]; then
    cat ${PROFILE} >> ${COVERAGE_REPORT}
    rm ${PROFILE}
  fi
done

mv ${COVERAGE_REPORT} /${COVERAGE_REPORT}

exit 0
