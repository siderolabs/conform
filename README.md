<p align="center">
  <h1 align="center">Conform</h1>
  <p align="center">Policy enforcement for your pipelines.</p>
  <p align="center">
    <a href="https://conventionalcommits.org"><img alt="Conventional Commits" src="https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg?style=flat-square"></a>
    <a href="https://godoc.org/github.com/talos-systems/conform"><img alt="GoDoc" src="http://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square"></a>
    <a href="https://travis-ci.org/talos-systems/conform"><img alt="Travis" src="https://img.shields.io/travis/talos-systems/conform.svg?style=flat-square"></a>
    <a href="https://codecov.io/gh/talos-systems/conform"><img alt="Codecov" src="https://img.shields.io/codecov/c/github/talos-systems/conform.svg?style=flat-square"></a>
    <a href="https://goreportcard.com/report/github.com/talos-systems/conform"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/talos-systems/conform?style=flat-square"></a>
    <a href="https://github.com/talos-systems/conform/releases/latest"><img alt="Release" src="https://img.shields.io/github/release/talos-systems/conform.svg?style=flat-square"></a>
    <a href="https://github.com/talos-systems/conform/releases/latest"><img alt="GitHub (pre-)release" src="https://img.shields.io/github/release/talos-systems/conform/all.svg?style=flat-square"></a>
  </p>
</p>

---

**Conform** is a tool for enforcing policies on your build pipelines.

Some of the policies included are:

- **Commits**: Enforce commit policies including:
  - Commit message header length
  - Developer Certificate of Origin
  - GPG signature
  - [Conventional Commits](https://www.conventionalcommits.org)
  - Imperative mood
  - Maximum of one commit ahead of `master`
  - Require a commit body
- **License Headers**: Enforce license headers on source code files.

## Getting Started

To install conform you can download a [release](https://github.com/talos-systems/conform/releases), or build it locally (go must be installed):

```bash
go get github.com/talos-systems/conform
```

Now, create a file named `.conform.yaml` with the following contents:

```yaml
policies:
  - type: commit
    spec:
      header:
        length: 89
        imperative: true
        case: lower
        invalidLastCharacters: .
      body:
        required: true
      dco: true
      gpg: false
      maximumOfOneCommit: true
      conventional:
        types:
          - "type"
        scopes:
          - "scope"
        descriptionLength: 72
  - type: license
    spec:
      skipPaths:
        - .git/
        - .build*/
      includeSuffixes:
        - .ext
      excludeSuffixes:
        - .exclude-ext-prefix.ext
      header: |
        This is the contents of a license header.
```

In the same directory, run:

```bash
$ conform enforce
POLICY         CHECK                        STATUS        MESSAGE
commit         Header Length                PASS          <none>
commit         Imperative Mood              PASS          <none>
commit         Header Case                  PASS          <none>
commit         Header Last Character        PASS          <none>
commit         DCO                          PASS          <none>
commit         Conventional Commit          PASS          <none>
commit         Number of Commits            PASS          <none>
commit         Commit Body                  PASS          <none>
license        File Header                  PASS          <none>
```

To setup a `commit-msg` hook:

```bash
cat <<EOF | tee .git/hooks/commit-msg
#!/bin/sh

conform enforce --commit-msg-file \$1
EOF
chmod +x .git/hooks/commit-msg
```

We also provide a [Pre-Commit](https://pre-commit.com) hook that you can use as follows:

```yaml
# .pre-commit-config.yaml
repos:
  - repo: https://github.com/talos-systems/conform
    rev: master
    hooks:
      - id: conform
        stages:
          - commit-msg
```

### License

[![license](https://img.shields.io/github/license/talos-systems/conform.svg?style=flat-square)](https://github.com/talos-systems/conform/blob/master/LICENSE)
