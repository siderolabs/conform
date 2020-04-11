# 0.1.0-alpha.19 (2020-04-11)

### Code Refactoring

- rename summarizer to reporter ([f5826e5](https://github.com/andrewrynhard/conform/commit/f5826e55f29fb34194cee5233b5621501a4c2171))

### Features

- add spellcheck commit check ([b3b6e65](https://github.com/andrewrynhard/conform/commit/b3b6e657fbd4c94f57b81ac95885339ff4ee1b59))

### BREAKING CHANGES

- This renames the `--summary` flag to `--reporter`.

Signed-off-by: Andrew Rynhard <andrew@andrewrynhard.com>

# 0.1.0-alpha.18 (2020-04-10)

### Bug Fixes

- Correctly return errors from command run function ([59f365a](https://github.com/andrewrynhard/conform/commit/59f365afe2384aabeee5617f28ef56b7f1ed6e17))

### Features

- move body options to dedicated field ([b2f63c1](https://github.com/andrewrynhard/conform/commit/b2f63c178998589d9f6f0273842ca38cc425a84d))
- **policy:** add checks for header case and last character ([fa7df19](https://github.com/andrewrynhard/conform/commit/fa7df19996ece307285da44c73f210c6cbec9207))
- add compatibility with pre-commit.com ([ea7199a](https://github.com/andrewrynhard/conform/commit/ea7199a68e686a6b6aea9a2c7741556c89680bd1))
- **policy:** Make Conventional commit description configurable ([d97b22c](https://github.com/andrewrynhard/conform/commit/d97b22c1bb88575b9e7a19a33dd47f8885096db8))

### BREAKING CHANGES

- This moves `spec.requireCommitBody` to
  `spec.body.required`.

Signed-off-by: Andrew Rynhard <andrew@andrewrynhard.com>

# 0.1.0-alpha.17 (2019-12-08)

### Bug Fixes

- add action.yml ([4eb4beb](https://github.com/andrewrynhard/conform/commit/4eb4beb060332b18c585161728f77aee84b46783))
- address conform errors ([cf1213a](https://github.com/andrewrynhard/conform/commit/cf1213a49ab9e3c6682d6ad8bd12fb71a8b5effc))
- checks conv commit format in firstWord ([7bed912](https://github.com/andrewrynhard/conform/commit/7bed9129bc73b5a6fd2b6d8c12f7c024ea4b7107))
- remove autonomy references ([d9668e0](https://github.com/andrewrynhard/conform/commit/d9668e05ebde8f8c3bc96abf9711665c040efb6f))
- set pass to false if errors in policies ([365c592](https://github.com/andrewrynhard/conform/commit/365c592227c7ecadf52a8d22db22fca0dea015a0))

### Features

- add server mode ([5637edd](https://github.com/andrewrynhard/conform/commit/5637edd03905d2d7612b0a86c7ac51aee7ae00b6))

# 0.1.0-alpha.16 (2019-07-05)

### Bug Fixes

- print an error message ([3b208c1](https://github.com/andrewrynhard/conform/commit/3b208c1f665672ce64673b7a89a159de67d99f85))

# 0.1.0-alpha.15 (2019-07-04)

### Features

- add commit body check ([adce353](https://github.com/andrewrynhard/conform/commit/adce353bae54580a95a302681cba96b2f1706a36))
- add number of commits check ([e66888a](https://github.com/andrewrynhard/conform/commit/e66888aeba8179a34646806ec050f348ba4a7669))

# 0.1.0-alpha.14 (2019-07-03)

### Bug Fixes

- add file header check ([#131](https://github.com/andrewrynhard/conform/issues/131)) ([ef30db9](https://github.com/andrewrynhard/conform/commit/ef30db9b206a4f67cc867eff4a510bbfd2a30a2d))

# 0.1.0-alpha.13 (2019-07-02)

### Features

- add support for GH actions on forked repo PRs ([#130](https://github.com/andrewrynhard/conform/issues/130)) ([cc97536](https://github.com/andrewrynhard/conform/commit/cc975363a9254c4878cd9a20708daf523bb576cf))

# 0.1.0-alpha.12 (2019-06-29)

### Bug Fixes

- excludeSuffixes wasn't skipping any files ([#120](https://github.com/andrewrynhard/conform/issues/120)) ([c539351](https://github.com/andrewrynhard/conform/commit/c5393510751f9ba440e01845eae43d423970a16b))
- trim whitespace while validating DCO ([#126](https://github.com/andrewrynhard/conform/issues/126)) ([0af31f8](https://github.com/andrewrynhard/conform/commit/0af31f88a74b836c6c8da15504730785f7b15ee8))

### Features

- add checks interface ([#127](https://github.com/andrewrynhard/conform/issues/127)) ([6f8751c](https://github.com/andrewrynhard/conform/commit/6f8751cb0791a8aeeb00194db3dc9c11059c7922))
- add support for github status checks ([#128](https://github.com/andrewrynhard/conform/issues/128)) ([3be1319](https://github.com/andrewrynhard/conform/commit/3be1319605a0ee934a5fdf5a0e5a050f3a7e2579))
- implement `skipPaths` option for 'license' policy ([#121](https://github.com/andrewrynhard/conform/issues/121)) ([ebed4b3](https://github.com/andrewrynhard/conform/commit/ebed4b31cc2e6f10914850ffba9cd73fca803333))

# 0.1.0-alpha.11 (2019-03-12)

### Bug Fixes

- check empty commit-msg prior to parsing ([#118](https://github.com/andrewrynhard/conform/issues/118)) ([37e0e69](https://github.com/andrewrynhard/conform/commit/37e0e6973100e596f29bbf43675472c5d8236679))

### Features

- change the license header to a string ([#116](https://github.com/andrewrynhard/conform/issues/116)) ([1473b44](https://github.com/andrewrynhard/conform/commit/1473b4462de868edcca18bb2dfd5108f0545232e))

# 0.1.0-alpha.10 (2019-02-19)

### Bug Fixes

- ensure the imperative check is against lowercase word ([#112](https://github.com/andrewrynhard/conform/issues/112)) ([6ac7c2f](https://github.com/andrewrynhard/conform/commit/6ac7c2f640fbcf85419fb9914b8bdcccd71570c0))
- use mood instead of verb ([#114](https://github.com/andrewrynhard/conform/issues/114)) ([bd039e4](https://github.com/andrewrynhard/conform/commit/bd039e43fde1298a388c3c75e982b08a9610c98d))

# 0.1.0-alpha.9 (2019-01-23)

### Bug Fixes

- **policy:** use natural language processing for imperative check ([#109](https://github.com/andrewrynhard/conform/issues/109)) ([3f75846](https://github.com/andrewrynhard/conform/commit/3f758468cea5db94fdd897dca4bc8c98016a5089))

### Features

- **policy:** add imperative mood check ([#108](https://github.com/andrewrynhard/conform/issues/108)) ([5c6620a](https://github.com/andrewrynhard/conform/commit/5c6620a1f544d9f3dd11cf5092efd698dc260827))
- add license header policy ([#105](https://github.com/andrewrynhard/conform/issues/105)) ([f5ed717](https://github.com/andrewrynhard/conform/commit/f5ed7174d6d5019e322d1acd4e9de47e1064a4f9))

# 0.1.0-alpha.8 (2019-01-13)

### Bug Fixes

- **policy:** remove commit header length from conventional commit policy ([#102](https://github.com/andrewrynhard/conform/issues/102)) ([116a3bf](https://github.com/andrewrynhard/conform/commit/116a3bf1bd5cf6dc6026b5d3f5fd09640f67380e))

# 0.1.0-alpha.7 (2019-01-13)

### Bug Fixes

- **ci:** push built images images ([#83](https://github.com/andrewrynhard/conform/issues/83)) ([f27917e](https://github.com/andrewrynhard/conform/commit/f27917e150dc0b6f14d31a75f209457f8fa94889))
- **metadata:** keep original version string ([#82](https://github.com/andrewrynhard/conform/issues/82)) ([3b165b3](https://github.com/andrewrynhard/conform/commit/3b165b3f8d332bb9e0c16d6e43ffeb923f5acbf7))
- **policy:** unit test inline git config ([#77](https://github.com/andrewrynhard/conform/issues/77)) ([954c003](https://github.com/andrewrynhard/conform/commit/954c00327ea1cb6d16c646ba3a97b5bd4f47170a))

### Features

- Add generic git commit policy ([#92](https://github.com/andrewrynhard/conform/issues/92)) ([b59ae9c](https://github.com/andrewrynhard/conform/commit/b59ae9c6fd5482b2558d43186b2951d34b7c6c40))
- output status in tab format ([#98](https://github.com/andrewrynhard/conform/issues/98)) ([7646221](https://github.com/andrewrynhard/conform/commit/7646221dc581b026654efdc9617c6eac76bb09a9))
- remove artifacts before creating ([#84](https://github.com/andrewrynhard/conform/issues/84)) ([76217df](https://github.com/andrewrynhard/conform/commit/76217df7c54752b6805fade67a548ec45c502919))
- show git status ([#85](https://github.com/andrewrynhard/conform/issues/85)) ([0fdd552](https://github.com/andrewrynhard/conform/commit/0fdd552a27d573cf9ecbc5e75b51b2e71fabcb09))
- **metadata:** add git ref ([#81](https://github.com/andrewrynhard/conform/issues/81)) ([18d8905](https://github.com/andrewrynhard/conform/commit/18d8905104b7db6bd05d33cbcfed65c5412cf16e))
- **metadata:** add original semver string ([#79](https://github.com/andrewrynhard/conform/issues/79)) ([5caf3b5](https://github.com/andrewrynhard/conform/commit/5caf3b5c99fb0cd164b2a1f60d9eef86969bc384))
- **policy:** show valid types and scopes on error ([#78](https://github.com/andrewrynhard/conform/issues/78)) ([d65491a](https://github.com/andrewrynhard/conform/commit/d65491ab160a33ccf166bfe3e1d4ba2596792f57))

# 0.1.0-alpha.6 (2018-10-08)

### Bug Fixes

- **policy:** return error in conventional commit report ([#75](https://github.com/andrewrynhard/conform/issues/75)) ([2ac5059](https://github.com/andrewrynhard/conform/commit/2ac50599f3f5944c44fa0c45790481e8f9d532b5))

### Features

- adding command line flag for commit msg ([#73](https://github.com/andrewrynhard/conform/issues/73)) ([4a6cc1c](https://github.com/andrewrynhard/conform/commit/4a6cc1cb9529ad561372c09fd7f579694c212b1b))
- omit symbol and DWARF symbol tables ([#70](https://github.com/andrewrynhard/conform/issues/70)) ([05cfacb](https://github.com/andrewrynhard/conform/commit/05cfacb6edd0dd92bd28bbb60c91970a1dc46d3a))

# 0.1.0-alpha.5 (2018-10-04)

### Bug Fixes

- **pipeline:** nil pointer when no defined pipeline ([#60](https://github.com/andrewrynhard/conform/issues/60)) ([1933d19](https://github.com/andrewrynhard/conform/commit/1933d192d81bfb67150be634905179bf4fe183e1))
- **policy:** update regex to allow optional scope ([#61](https://github.com/andrewrynhard/conform/issues/61)) ([aed2c22](https://github.com/andrewrynhard/conform/commit/aed2c223188bb7ee04046976a4e1249e1451d8ca))

### Features

- add build command ([#62](https://github.com/andrewrynhard/conform/issues/62)) ([0a0cba3](https://github.com/andrewrynhard/conform/commit/0a0cba34137d39002bf762b52d19121290be9980))
- **metadata:** mark SHA as dirty ([#63](https://github.com/andrewrynhard/conform/issues/63)) ([155b036](https://github.com/andrewrynhard/conform/commit/155b0369ddf5bea3b162067b804a487fb3b634e2))

# 0.1.0-alpha.4 (2018-03-07)

### Bug Fixes

- **cli:** invalid variable message ([#53](https://github.com/andrewrynhard/conform/issues/53)) ([1e57715](https://github.com/andrewrynhard/conform/commit/1e577157c87a76461c8576a7f6b696ab3704c53f))

### Features

- **docker:** expose the image name and tag separately ([#58](https://github.com/andrewrynhard/conform/issues/58)) ([42c5d09](https://github.com/andrewrynhard/conform/commit/42c5d09d8189e222a5785232f98b000bc869228f))
- **fmt:** add fmt command ([#59](https://github.com/andrewrynhard/conform/issues/59)) ([7fd1e89](https://github.com/andrewrynhard/conform/commit/7fd1e89c567df89bbc16449182ec10a9ce2e5c18))
- **git:** recursively search for .git in parent directories ([#56](https://github.com/andrewrynhard/conform/issues/56)) ([5a73ea6](https://github.com/andrewrynhard/conform/commit/5a73ea6ed7490f91995e9fb916213226ac50fb01))
- **metadata:** allow users to specify variables ([#52](https://github.com/andrewrynhard/conform/issues/52)) ([72061b1](https://github.com/andrewrynhard/conform/commit/72061b11a121bd6ba1ad6e25a0e66059a09fc5d8))

# 0.1.0-alpha.3 (2018-01-11)

### Bug Fixes

- **pipeline:** don't show stdout of artifact extraction ([#49](https://github.com/andrewrynhard/conform/issues/49)) ([779bf93](https://github.com/andrewrynhard/conform/commit/779bf930fb4cdbe2d94b39c2218f9b2854efaf7a))
- **policy:** check the entire commit header length ([#31](https://github.com/andrewrynhard/conform/issues/31)) ([44d4ef0](https://github.com/andrewrynhard/conform/commit/44d4ef0256dd48965eabcbeade2838bb9ccdeaea))
- **policy:** strip leading newline from commit message ([#50](https://github.com/andrewrynhard/conform/issues/50)) ([4b7b903](https://github.com/andrewrynhard/conform/commit/4b7b903aac3671f839115f6240d4e12e5256e742))

### Features

- services, skip flag, and UX improvements ([#43](https://github.com/andrewrynhard/conform/issues/43)) ([0373fea](https://github.com/andrewrynhard/conform/commit/0373fea42e663c91e512fe1c144721b299bae457))
- **policy:** enforce 72 character limit on commit header ([#29](https://github.com/andrewrynhard/conform/issues/29)) ([9383d3e](https://github.com/andrewrynhard/conform/commit/9383d3ebab5b14c186fc84299c76cfa4d404a9c6))
- **renderer:** allow templates to be retrieved from URL ([#41](https://github.com/andrewrynhard/conform/issues/41)) ([c53f523](https://github.com/andrewrynhard/conform/commit/c53f52307152a6392835e23797f1432799a1e0cc))

# 0.1.0-alpha.2 (2017-07-22)

### Bug Fixes

- **metadata:** nil version struct ([#22](https://github.com/andrewrynhard/conform/issues/22)) ([a18332a](https://github.com/andrewrynhard/conform/commit/a18332af48f1e5ddb5308ac9ed3d1ff030e96a1b))

### Features

- **policy:** add policy enforcement; enforce git commit policy ([#17](https://github.com/andrewrynhard/conform/issues/17)) ([03caad0](https://github.com/andrewrynhard/conform/commit/03caad0cb1a02e8bf688557f21c3bd58b2c69fdc))

# 0.1.0-alpha.1 (2017-07-04)

# 0.1.0-alpha.0 (2017-06-05)
