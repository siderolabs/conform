<p align="center">
  <h1 align="center">Conform</h1>
  <p align="center">Policy enforcement for your pipelines.</p>
  <p align="center">
    <a href="https://conventionalcommits.org"><img alt="Conventional Commits" src="https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg?style=flat-square"></a>
    <a href="https://godoc.org/github.com/autonomy/conform"><img alt="GoDoc" src="http://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square"></a>
    <a href="https://travis-ci.org/autonomy/conform"><img alt="Travis" src="https://img.shields.io/travis/autonomy/conform.svg?style=flat-square"></a>
    <a href="https://codecov.io/gh/autonomy/conform"><img alt="Codecov" src="https://img.shields.io/codecov/c/github/autonomy/conform.svg?style=flat-square"></a>
    <a href="https://goreportcard.com/report/github.com/autonomy/conform"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/autonomy/conform?style=flat-square"></a>
    <a href="https://github.com/autonomy/conform/releases/latest"><img alt="Release" src="https://img.shields.io/github/release/autonomy/conform.svg?style=flat-square"></a>
    <a href="https://github.com/autonomy/conform/releases/latest"><img alt="GitHub (pre-)release" src="https://img.shields.io/github/release/autonomy/conform/all.svg?style=flat-square"></a>
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

Create a file named `.conform.yaml` with the following contents:

```yaml
policies:
  - type: commit
    spec:
      headerLength: 89
      dco: true
      gpg: false
      imperative: true
      maximumOfOneCommit: true
      requireCommitBody: true
      conventional:
        types:
          - "type"
        scopes:
          - "scope"
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
POLICY         CHECK                      STATUS        MESSAGE
commit         Header Length              PASS          <none>
commit         DCO                        PASS          <none>
commit         Imperative Mood            PASS          <none>
commit         Conventional Commit        PASS          <none>
commit         Number of Commits          PASS          <none>
commit         Commit Body                PASS          <none>
license        File Header                PASS          <none>
```

### License
[![license](https://img.shields.io/github/license/autonomy/conform.svg?style=flat-square)](https://github.com/autonomy/conform/blob/master/LICENSE)
