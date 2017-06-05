<p align="center">
  <h1 align="center">Conform</h1>
  <p align="center">DRY, hygienic, fast builds.</p>
  <p align="center">
    <a href="https://gitter.im/autonomy/conform"><img alt="Gitter" src="https://img.shields.io/gitter/room/autonomy/conform.svg?style=flat-square"></a>
    <a href="https://godoc.org/github.com/autonomy/conform"><img alt="GoDoc" src="http://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square"></a>
  </p>
  <p align="center">
    <a href="https://travis-ci.org/autonomy/conform"><img alt="Travis" src="https://img.shields.io/travis/autonomy/conform.svg?style=flat-square"></a>
    <a href="https://codecov.io/gh/autonomy/conform"><img alt="Codecov" src="https://img.shields.io/codecov/c/github/autonomy/conform.svg?style=flat-square"></a>
    <a href="https://goreportcard.com/report/github.com/autonomy/conform"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/autonomy/conform?style=flat-square"></a>
  </p>
  <p align="center">
    <a href="https://github.com/autonomy/conform/releases/latest"><img alt="Release" src="https://img.shields.io/github/release/autonomy/conform.svg?style=flat-square"></a>
    <a href="https://github.com/autonomy/conform/releases/latest"><img alt="GitHub (pre-)release" src="https://img.shields.io/github/release/autonomy/conform/all.svg?style=flat-square"></a>
  </p>
</p>

---

**Conform** is a tool for building projects in a flexible and reliabale manner.

The key features of Conform are:
-   **DRY**: Templatized multi-stage Docker builds.
-   **Hygienic**: Builds run in Docker.
-   **Fast**: Leverages Docker caching, building only what has changed.

Getting Started
---------------
Create a file named `conform.yaml` with the following contents:
```yaml
metadata:
  repository: example
scripts:
  init : |
    #!/bin/bash

    set -e

    echo "Initialize any dependencies here."
  deploy : |
    #!/bin/bash

    set -e

    echo "Deploy you image here."
templates:
  build: |
    FROM alpine:latest as build
    RUN echo "Run your build here."
    RUN touch artifact
  test: |
    FROM alpine:latest as test
    COPY --from=build artifact .
    RUN echo "Run your tests here."
  image: |
    FROM scratch as image
    RUN echo "Prepare your final image here."
    COPY --from=build artifact .
rules:
  image:
    before:
      - init
    templates:
      - build
      - test
      - image
    after:
      - deploy
```

In the same directory, run:
```
$ conform enforce image
```

Devloping Conform
----------------

### License
[![license](https://img.shields.io/github/license/autonomy/conform.svg?style=flat-square)](https://github.com/autonomy/conform/blob/master/LICENSE)
