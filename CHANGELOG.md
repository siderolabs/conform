<a name="0.1.0-alpha.9"></a>
# [0.1.0-alpha.9](https://github.com/autonomy/conform/compare/v0.1.0-alpha.8...v0.1.0-alpha.9) (2019-01-21)


### Bug Fixes

* **policy:** use natural language processing for imperative check ([#109](https://github.com/autonomy/conform/issues/109)) ([3f75846](https://github.com/autonomy/conform/commit/3f75846))


### Features

* add license header policy ([#105](https://github.com/autonomy/conform/issues/105)) ([f5ed717](https://github.com/autonomy/conform/commit/f5ed717))
* **policy:** add imperative mood check ([#108](https://github.com/autonomy/conform/issues/108)) ([5c6620a](https://github.com/autonomy/conform/commit/5c6620a))



<a name="0.1.0-alpha.8"></a>
# [0.1.0-alpha.8](https://github.com/autonomy/conform/compare/v0.1.0-alpha.7...v0.1.0-alpha.8) (2019-01-13)


### Bug Fixes

* **policy:** remove commit header length from conventional commit policy ([#102](https://github.com/autonomy/conform/issues/102)) ([116a3bf](https://github.com/autonomy/conform/commit/116a3bf))



<a name="0.1.0-alpha.7"></a>
# [0.1.0-alpha.7](https://github.com/autonomy/conform/compare/v0.1.0-alpha.6...v0.1.0-alpha.7) (2019-01-13)


### Bug Fixes

* **ci:** push built images images ([#83](https://github.com/autonomy/conform/issues/83)) ([f27917e](https://github.com/autonomy/conform/commit/f27917e))
* **metadata:** keep original version string ([#82](https://github.com/autonomy/conform/issues/82)) ([3b165b3](https://github.com/autonomy/conform/commit/3b165b3))
* **policy:** unit test inline git config ([#77](https://github.com/autonomy/conform/issues/77)) ([954c003](https://github.com/autonomy/conform/commit/954c003))


### Features

* **metadata:** add git ref ([#81](https://github.com/autonomy/conform/issues/81)) ([18d8905](https://github.com/autonomy/conform/commit/18d8905))
* **metadata:** add original semver string ([#79](https://github.com/autonomy/conform/issues/79)) ([5caf3b5](https://github.com/autonomy/conform/commit/5caf3b5))
* **policy:** show valid types and scopes on error ([#78](https://github.com/autonomy/conform/issues/78)) ([d65491a](https://github.com/autonomy/conform/commit/d65491a))
* Add generic git commit policy ([#92](https://github.com/autonomy/conform/issues/92)) ([b59ae9c](https://github.com/autonomy/conform/commit/b59ae9c))
* output status in tab format ([#98](https://github.com/autonomy/conform/issues/98)) ([7646221](https://github.com/autonomy/conform/commit/7646221))
* remove artifacts before creating ([#84](https://github.com/autonomy/conform/issues/84)) ([76217df](https://github.com/autonomy/conform/commit/76217df))
* show git status ([#85](https://github.com/autonomy/conform/issues/85)) ([0fdd552](https://github.com/autonomy/conform/commit/0fdd552))



<a name="0.1.0-alpha.6"></a>
# [0.1.0-alpha.6](https://github.com/autonomy/conform/compare/v0.1.0-alpha.5...v0.1.0-alpha.6) (2018-10-08)


### Bug Fixes

* **policy:** return error in conventional commit report ([#75](https://github.com/autonomy/conform/issues/75)) ([2ac5059](https://github.com/autonomy/conform/commit/2ac5059))


### Features

* omit symbol and DWARF symbol tables ([#70](https://github.com/autonomy/conform/issues/70)) ([05cfacb](https://github.com/autonomy/conform/commit/05cfacb))
* adding command line flag for commit msg ([#73](https://github.com/autonomy/conform/issues/73)) ([4a6cc1c](https://github.com/autonomy/conform/commit/4a6cc1c))



<a name="0.1.0-alpha.5"></a>
# [0.1.0-alpha.5](https://github.com/autonomy/conform/compare/v0.1.0-alpha.4...v0.1.0-alpha.5) (2018-10-04)


### Bug Fixes

* **pipeline:** nil pointer when no defined pipeline ([#60](https://github.com/autonomy/conform/issues/60)) ([1933d19](https://github.com/autonomy/conform/commit/1933d19))
* **policy:** update regex to allow optional scope ([#61](https://github.com/autonomy/conform/issues/61)) ([aed2c22](https://github.com/autonomy/conform/commit/aed2c22))


### Features

* add build command ([#62](https://github.com/autonomy/conform/issues/62)) ([0a0cba3](https://github.com/autonomy/conform/commit/0a0cba3))
* **metadata:** mark SHA as dirty ([#63](https://github.com/autonomy/conform/issues/63)) ([155b036](https://github.com/autonomy/conform/commit/155b036))



<a name="0.1.0-alpha.4"></a>
# [0.1.0-alpha.4](https://github.com/autonomy/conform/compare/v0.1.0-alpha.3...v0.1.0-alpha.4) (2018-03-07)


### Bug Fixes

* **cli:** invalid variable message ([#53](https://github.com/autonomy/conform/issues/53)) ([1e57715](https://github.com/autonomy/conform/commit/1e57715))


### Features

* **docker:** expose the image name and tag separately ([#58](https://github.com/autonomy/conform/issues/58)) ([42c5d09](https://github.com/autonomy/conform/commit/42c5d09))
* **fmt:** add fmt command ([#59](https://github.com/autonomy/conform/issues/59)) ([7fd1e89](https://github.com/autonomy/conform/commit/7fd1e89))
* **git:** recursively search for .git in parent directories ([#56](https://github.com/autonomy/conform/issues/56)) ([5a73ea6](https://github.com/autonomy/conform/commit/5a73ea6))
* **metadata:** allow users to specify variables ([#52](https://github.com/autonomy/conform/issues/52)) ([72061b1](https://github.com/autonomy/conform/commit/72061b1))



<a name="0.1.0-alpha.3"></a>
# [0.1.0-alpha.3](https://github.com/autonomy/conform/compare/v0.1.0-alpha.2...v0.1.0-alpha.3) (2018-01-11)


### Bug Fixes

* **pipeline:** don't show stdout of artifact extraction ([#49](https://github.com/autonomy/conform/issues/49)) ([779bf93](https://github.com/autonomy/conform/commit/779bf93))
* **policy:** check the entire commit header length ([#31](https://github.com/autonomy/conform/issues/31)) ([44d4ef0](https://github.com/autonomy/conform/commit/44d4ef0))
* **policy:** strip leading newline from commit message ([#50](https://github.com/autonomy/conform/issues/50)) ([4b7b903](https://github.com/autonomy/conform/commit/4b7b903))


### Features

* services, skip flag, and UX improvements ([#43](https://github.com/autonomy/conform/issues/43)) ([0373fea](https://github.com/autonomy/conform/commit/0373fea))
* **policy:** enforce 72 character limit on commit header ([#29](https://github.com/autonomy/conform/issues/29)) ([9383d3e](https://github.com/autonomy/conform/commit/9383d3e))
* **renderer:** allow templates to be retrieved from URL ([#41](https://github.com/autonomy/conform/issues/41)) ([c53f523](https://github.com/autonomy/conform/commit/c53f523))



<a name="0.1.0-alpha.2"></a>
# [0.1.0-alpha.2](https://github.com/autonomy/conform/compare/v0.1.0-alpha.1...v0.1.0-alpha.2) (2017-07-22)


### Bug Fixes

* **metadata:** nil version struct ([#22](https://github.com/autonomy/conform/issues/22)) ([a18332a](https://github.com/autonomy/conform/commit/a18332a))


### Features

* **policy:** add policy enforcement; enforce git commit policy ([#17](https://github.com/autonomy/conform/issues/17)) ([03caad0](https://github.com/autonomy/conform/commit/03caad0))



<a name="0.1.0-alpha.1"></a>
# [0.1.0-alpha.1](https://github.com/autonomy/conform/compare/v0.1.0-alpha.0...v0.1.0-alpha.1) (2017-07-04)



<a name="0.1.0-alpha.0"></a>
# 0.1.0-alpha.0 (2017-06-05)



