## [conform 0.1.0-alpha.25](https://github.com/talos-systems/conform/releases/tag/v0.1.0-alpha.25) (2021-12-21)

Welcome to the v0.1.0-alpha.25 release of conform!



Please try out the release binaries and report any issues at
https://github.com/talos-systems/conform/issues.

### Contributors

* Andrew Rynhard
* Andrey Smirnov
* Andrew Rynhard
* Andrey Smirnov
* Joey Espinosa
* Alexey Palazhchenko
* Andrey Smirnov
* Peter Ng
* Wouter Dullaert
* Alex Szakaly
* André Valentin
* Anton Isakov
* Brad Beam
* Danny Zhu
* Dave Forgac
* Don Bowman
* Per Abich

### Changes
<details><summary>157 commits</summary>
<p>

* [`18107a6`](https://github.com/talos-systems/conform/commit/18107a6d00cf30d420c6ef3a5c14f64efa4f385b) fix: checkout a branch only once
* [`f00081b`](https://github.com/talos-systems/conform/commit/f00081bdfeb35517095dbcbd053305e9c8ae6fc9) release(v0.1.0-alpha.24): prepare release
* [`e9bb88a`](https://github.com/talos-systems/conform/commit/e9bb88a607115a6ff0008aec919f672f981038f2) fix: don't try git checkout repository twice
* [`bb8186b`](https://github.com/talos-systems/conform/commit/bb8186be5f8d6688284621c00c04d7cec9885091) fix: update jira issue valid message
* [`785d27d`](https://github.com/talos-systems/conform/commit/785d27dc5487aa6c5b5709a898d4e982c7f88410) release(v0.1.0-alpha.23): prepare release
* [`6caaabc`](https://github.com/talos-systems/conform/commit/6caaabc0f9721d161e68593fe3ee776761bdc893) feat: allow comments before the license header
* [`9a470a1`](https://github.com/talos-systems/conform/commit/9a470a16fe05643de83db40bf3f8cb0c3452ee2f) docs: update README with the latest changes
* [`35fba60`](https://github.com/talos-systems/conform/commit/35fba6014b3ecdd155c82fd931051b90c2aa9c35) chore: use kres to setup build instructions
* [`fb43bd4`](https://github.com/talos-systems/conform/commit/fb43bd4b31cca4206f0469b1afffcb1666685ab2) feat: provide gpg identity signature check
* [`c23e2fc`](https://github.com/talos-systems/conform/commit/c23e2fc7c6159cdd3ff6b78eb7591c654171e1a8) feat: add conventional commit breaking change support
* [`0e3a28c`](https://github.com/talos-systems/conform/commit/0e3a28c994aa9a0727d48bcde91c39034a2056e1) fix: update action version
* [`0390165`](https://github.com/talos-systems/conform/commit/03901655927737cedf5c0e901f72832cde9ee03a) fix: change check_jira regexp
* [`73f334a`](https://github.com/talos-systems/conform/commit/73f334a4ac7e6a799bc71f0611674f26b9b9c3a9) fix: disallow 0 as valid jira issue id
* [`5b58100`](https://github.com/talos-systems/conform/commit/5b58100ec8619aa94eb855546ddbeb8f27a35a24) chore: bump golangci-lint and fix linting issues
* [`c5dc2e6`](https://github.com/talos-systems/conform/commit/c5dc2e6b86fa76eea7fbf37a064ca6b9962aee96) fix: change "jiraKeys" to "keys"
* [`efd7fbb`](https://github.com/talos-systems/conform/commit/efd7fbb0dbc0042f54af587ffa07a829fdc88f1c) feat: improve Jira check
* [`001de56`](https://github.com/talos-systems/conform/commit/001de5691ecce630808bc2716229d9ee700a1ab4) feat: add support for requiring a Jira issue in the header
* [`5a75e96`](https://github.com/talos-systems/conform/commit/5a75e96171ef6aaaebecfced193df379296e03d7) chore: remove gitmeta
* [`cea1ee9`](https://github.com/talos-systems/conform/commit/cea1ee90f257daa96aa09ef9ba8b1deafddf92b2) chore: bump golangci-lint and fix linting issues
* [`6e0c294`](https://github.com/talos-systems/conform/commit/6e0c294bc044bde8dc43d45e31dee3b41a97fd53) feat: implement full gitignore-style pattern matching
* [`72708f2`](https://github.com/talos-systems/conform/commit/72708f25b12786e10a7f49da584a638ac76bfe92) feat: support regex in conventional commit scope
* [`ec5e365`](https://github.com/talos-systems/conform/commit/ec5e3656494068721fa29a815ea8abf67a4922f6) fix: checkout pull requests
* [`6d1a620`](https://github.com/talos-systems/conform/commit/6d1a620acbe83af09a97468129d02a366aec9b24) fix: use text/template
* [`8212dc6`](https://github.com/talos-systems/conform/commit/8212dc6b2353af86d0f69fd0fe4458b23507a7ec) chore: update CHANGELOG
* [`b3b6e65`](https://github.com/talos-systems/conform/commit/b3b6e657fbd4c94f57b81ac95885339ff4ee1b59) feat: add spellcheck commit check
* [`8726189`](https://github.com/talos-systems/conform/commit/8726189af55c5278de18b464273b89fed024aa2a) chore: update go-git
* [`f5826e5`](https://github.com/talos-systems/conform/commit/f5826e55f29fb34194cee5233b5621501a4c2171) refactor: rename summarizer to reporter
* [`b2f63c1`](https://github.com/talos-systems/conform/commit/b2f63c178998589d9f6f0273842ca38cc425a84d) feat: move body options to dedicated field
* [`fa7df19`](https://github.com/talos-systems/conform/commit/fa7df19996ece307285da44c73f210c6cbec9207) feat(policy): add checks for header case and last character
* [`a55d411`](https://github.com/talos-systems/conform/commit/a55d411e8f9061b6ab5b8dacd96c03254f57a557) docs: update README to include pre-commit compatibility
* [`ea7199a`](https://github.com/talos-systems/conform/commit/ea7199a68e686a6b6aea9a2c7741556c89680bd1) feat: add compatibility with pre-commit.com
* [`d97b22c`](https://github.com/talos-systems/conform/commit/d97b22c1bb88575b9e7a19a33dd47f8885096db8) feat(policy): Make Conventional commit description configurable
* [`59f365a`](https://github.com/talos-systems/conform/commit/59f365afe2384aabeee5617f28ef56b7f1ed6e17) fix: Correctly return errors from command run function
* [`01f87b9`](https://github.com/talos-systems/conform/commit/01f87b956d3b317e834713a918acaa3726fc7980) docs: add installation guide to README
* [`4eb4beb`](https://github.com/talos-systems/conform/commit/4eb4beb060332b18c585161728f77aee84b46783) fix: add action.yml
* [`d9668e0`](https://github.com/talos-systems/conform/commit/d9668e05ebde8f8c3bc96abf9711665c040efb6f) fix: remove autonomy references
* [`cf1213a`](https://github.com/talos-systems/conform/commit/cf1213a49ab9e3c6682d6ad8bd12fb71a8b5effc) fix: address conform errors
* [`5637edd`](https://github.com/talos-systems/conform/commit/5637edd03905d2d7612b0a86c7ac51aee7ae00b6) feat: add server mode
* [`0076446`](https://github.com/talos-systems/conform/commit/007644608f2455688d194164d2d29a3fa58a9918) chore: Replace autonomy with talos-systems
* [`7bed912`](https://github.com/talos-systems/conform/commit/7bed9129bc73b5a6fd2b6d8c12f7c024ea4b7107) fix: checks conv commit format in firstWord
* [`365c592`](https://github.com/talos-systems/conform/commit/365c592227c7ecadf52a8d22db22fca0dea015a0) fix: set pass to false if errors in policies
* [`19dd2b8`](https://github.com/talos-systems/conform/commit/19dd2b82ab96a65ab739ff00c5d7331b1293e6ce) chore: set docker server entrypoint to dockerd to avoid TLS generation
* [`25d013c`](https://github.com/talos-systems/conform/commit/25d013c905306829da13ed14189e994bac5a8137) docs: update README
* [`3b208c1`](https://github.com/talos-systems/conform/commit/3b208c1f665672ce64673b7a89a159de67d99f85) fix: print an error message
* [`4716823`](https://github.com/talos-systems/conform/commit/4716823c79156870677e239cf229b9af3f28273d) chore: require a commit body
* [`7809e90`](https://github.com/talos-systems/conform/commit/7809e90da2a0c4a087ea83a4c2a684f77f1dfe76) chore: push images for all branches
* [`0f7ec4c`](https://github.com/talos-systems/conform/commit/0f7ec4cfec40ea0ee377c56e4d356898064df1a4) chore: fix container labels
* [`adce353`](https://github.com/talos-systems/conform/commit/adce353bae54580a95a302681cba96b2f1706a36) feat: add commit body check
* [`e66888a`](https://github.com/talos-systems/conform/commit/e66888aeba8179a34646806ec050f348ba4a7669) feat: add number of commits check
* [`b330410`](https://github.com/talos-systems/conform/commit/b33041034a6b857f140f9d1b8485da3dc4e0a939) chore: build image with host net to avoid apk hang (#133)
* [`9e94c43`](https://github.com/talos-systems/conform/commit/9e94c43cfe66aff074a8c106899104a60a64272f) chore: prepare release v0.1.0-alpha.14 (#132)
* [`ef30db9`](https://github.com/talos-systems/conform/commit/ef30db9b206a4f67cc867eff4a510bbfd2a30a2d) fix: add file header check (#131)
* [`cc97536`](https://github.com/talos-systems/conform/commit/cc975363a9254c4878cd9a20708daf523bb576cf) feat: add support for GH actions on forked repo PRs (#130)
* [`4447684`](https://github.com/talos-systems/conform/commit/4447684a4433d321f1586f90bc2d107e6e7bcb91) chore: prepare release v0.1.0-alpha.12 (#129)
* [`3be1319`](https://github.com/talos-systems/conform/commit/3be1319605a0ee934a5fdf5a0e5a050f3a7e2579) feat: add support for github status checks (#128)
* [`6f8751c`](https://github.com/talos-systems/conform/commit/6f8751cb0791a8aeeb00194db3dc9c11059c7922) feat: add checks interface (#127)
* [`0af31f8`](https://github.com/talos-systems/conform/commit/0af31f88a74b836c6c8da15504730785f7b15ee8) fix: trim whitespace while validating DCO (#126)
* [`57c9dbd`](https://github.com/talos-systems/conform/commit/57c9dbd056d607fa0be2c5e1eb4628900ef85c2c) chore: quote docker creds (#122)
* [`ebed4b3`](https://github.com/talos-systems/conform/commit/ebed4b31cc2e6f10914850ffba9cd73fca803333) feat: implement `skipPaths` option for 'license' policy (#121)
* [`c539351`](https://github.com/talos-systems/conform/commit/c5393510751f9ba440e01845eae43d423970a16b) fix: excludeSuffixes wasn't skipping any files (#120)
* [`37e0e69`](https://github.com/talos-systems/conform/commit/37e0e6973100e596f29bbf43675472c5d8236679) fix: check empty commit-msg prior to parsing (#118)
* [`1473b44`](https://github.com/talos-systems/conform/commit/1473b4462de868edcca18bb2dfd5108f0545232e) feat: change the license header to a string (#116)
* [`abfd427`](https://github.com/talos-systems/conform/commit/abfd427a40a1eb8f2ee953e2442afd353625e460) chore: prepare v0.1.0-alpha.10 release (#115)
* [`bd039e4`](https://github.com/talos-systems/conform/commit/bd039e43fde1298a388c3c75e982b08a9610c98d) fix: use mood instead of verb (#114)
* [`6ac7c2f`](https://github.com/talos-systems/conform/commit/6ac7c2f640fbcf85419fb9914b8bdcccd71570c0) fix: ensure the imperative check is against lowercase word (#112)
* [`fd6ad6c`](https://github.com/talos-systems/conform/commit/fd6ad6cdb0746bc5547b5bfc5d474c9ab84d68f7) chore: prepare v0.1.0-alpha.9 release (#111)
* [`286041a`](https://github.com/talos-systems/conform/commit/286041a7b48aa8b789bbab398de3a85f02ba87bd) docs: update README (#110)
* [`3f75846`](https://github.com/talos-systems/conform/commit/3f758468cea5db94fdd897dca4bc8c98016a5089) fix(policy): use natural language processing for imperative check (#109)
* [`5c6620a`](https://github.com/talos-systems/conform/commit/5c6620a1f544d9f3dd11cf5092efd698dc260827) feat(policy): add imperative mood check (#108)
* [`86a7d3e`](https://github.com/talos-systems/conform/commit/86a7d3e57de33b2800c5feb1deb4655d51fc151d) docs: fix code highlight in README (#107)
* [`eeb3d5c`](https://github.com/talos-systems/conform/commit/eeb3d5ce7ff602af855de37ec5f30a34f45e550c) docs: move LICENSE_HEADER to root of project (#106)
* [`f5ed717`](https://github.com/talos-systems/conform/commit/f5ed7174d6d5019e322d1acd4e9de47e1064a4f9) feat: add license header policy (#105)
* [`763d4d9`](https://github.com/talos-systems/conform/commit/763d4d94587a7000ebd67d7277d38275179184b0) docs: add conventional commits badge to README (#104)
* [`e4602b8`](https://github.com/talos-systems/conform/commit/e4602b810eace5fe01ef191d58026a60e448e9c6) chore: prepare v0.1.0-alpha.8 release (#103)
* [`116a3bf`](https://github.com/talos-systems/conform/commit/116a3bf1bd5cf6dc6026b5d3f5fd09640f67380e) fix(policy): remove commit header length from conventional commit policy (#102)
* [`2be1e1e`](https://github.com/talos-systems/conform/commit/2be1e1e7eaf5827c960d747f831a376c10a433ce) docs: update README (#101)
* [`22804ff`](https://github.com/talos-systems/conform/commit/22804ff48efae434dc3bfc7b1e775c747d5772da) chore: fix image tag (#100)
* [`e6664a9`](https://github.com/talos-systems/conform/commit/e6664a9705169e76673632bbb6c68b23b4b1194b) chore: prepare v0.1.0-alpha.7 release (#99)
* [`7646221`](https://github.com/talos-systems/conform/commit/7646221dc581b026654efdc9617c6eac76bb09a9) feat: output status in tab format (#98)
* [`e93a47e`](https://github.com/talos-systems/conform/commit/e93a47ee759d0345e9a2ad3526a64dcc1715f366) chore: push latest tag (#97)
* [`598b595`](https://github.com/talos-systems/conform/commit/598b595b47df3e87c7eae5fd997c3f01d2af859f) chore: revert base image to scratch (#96)
* [`eb6cc6d`](https://github.com/talos-systems/conform/commit/eb6cc6db46c88ecfe2099288c1aa5826185d1575) chore: add conform binary to /bin (#95)
* [`fd41a2f`](https://github.com/talos-systems/conform/commit/fd41a2f3619e35f9770dfddfc00c2f0eb94637bd) chore: use alpine:3.8 as base image (#94)
* [`36df035`](https://github.com/talos-systems/conform/commit/36df0355239e6fee6767b3b2e536d976ce76bfdd) chore: use buildkit for builds (#93)
* [`b59ae9c`](https://github.com/talos-systems/conform/commit/b59ae9c6fd5482b2558d43186b2951d34b7c6c40) feat: Add generic git commit policy (#92)
* [`76b6d7a`](https://github.com/talos-systems/conform/commit/76b6d7a2a7346c5d09de14bec6cb757f4a0ddb9d) chore: fix image push (#91)
* [`22e0f7b`](https://github.com/talos-systems/conform/commit/22e0f7bb7bf6ca5face6f75036fbe54407090699) chore: pin Kaniko to v0.6.0 (#90)
* [`9b72c17`](https://github.com/talos-systems/conform/commit/9b72c1797945a00cba940df4e47fa0fa8718f9ce) docs: fix README example (#89)
* [`5543c79`](https://github.com/talos-systems/conform/commit/5543c7908afe1d205234b3acefe883df2b359b4f) chore: fix typos (#88)
* [`349ba37`](https://github.com/talos-systems/conform/commit/349ba37ba4957efc1f99e24b5c34cb664dd0aa5e) docs: fix README (#87)
* [`8ced588`](https://github.com/talos-systems/conform/commit/8ced58895fb10576407e85edfadfbd568861e6ec) refactor: use Kaniko for builds (#86)
* [`0fdd552`](https://github.com/talos-systems/conform/commit/0fdd552a27d573cf9ecbc5e75b51b2e71fabcb09) feat: show git status (#85)
* [`76217df`](https://github.com/talos-systems/conform/commit/76217df7c54752b6805fade67a548ec45c502919) feat: remove artifacts before creating (#84)
* [`3b165b3`](https://github.com/talos-systems/conform/commit/3b165b3f8d332bb9e0c16d6e43ffeb923f5acbf7) fix(metadata): keep original version string (#82)
* [`f27917e`](https://github.com/talos-systems/conform/commit/f27917e150dc0b6f14d31a75f209457f8fa94889) fix(ci): push built images images (#83)
* [`e67dd43`](https://github.com/talos-systems/conform/commit/e67dd4363be3567a4094949ccf0d3263d8aa8dcc) chore(ci): add brigade configuration (#80)
* [`18d8905`](https://github.com/talos-systems/conform/commit/18d8905104b7db6bd05d33cbcfed65c5412cf16e) feat(metadata): add git ref (#81)
* [`5caf3b5`](https://github.com/talos-systems/conform/commit/5caf3b5c99fb0cd164b2a1f60d9eef86969bc384) feat(metadata): add original semver string (#79)
* [`d65491a`](https://github.com/talos-systems/conform/commit/d65491ab160a33ccf166bfe3e1d4ba2596792f57) feat(policy): show valid types and scopes on error (#78)
* [`954c003`](https://github.com/talos-systems/conform/commit/954c00327ea1cb6d16c646ba3a97b5bd4f47170a) fix(policy): unit test inline git config (#77)
* [`e439cd7`](https://github.com/talos-systems/conform/commit/e439cd7bee060a0c16bc6d9ac23f4e87bc46da9e) chore(ci): run go mod tidy (#76)
* [`2ac5059`](https://github.com/talos-systems/conform/commit/2ac50599f3f5944c44fa0c45790481e8f9d532b5) fix(policy): return error in conventional commit report (#75)
* [`7d19c82`](https://github.com/talos-systems/conform/commit/7d19c82835585768a7601a1745a344a347f0ee9c) chore(ci): show git status when dirty working tree (#74)
* [`4a6cc1c`](https://github.com/talos-systems/conform/commit/4a6cc1cb9529ad561372c09fd7f579694c212b1b) feat: adding command line flag for commit msg (#73)
* [`088e0a7`](https://github.com/talos-systems/conform/commit/088e0a76b6e04ca521c6b9e03e66acb03b8dfbe8) chore(*): output binaries (#72)
* [`4194aa5`](https://github.com/talos-systems/conform/commit/4194aa5004222589c36265d4649452b72fb744bd) chore(*): format .conform.yaml (#71)
* [`05cfacb`](https://github.com/talos-systems/conform/commit/05cfacb6edd0dd92bd28bbb60c91970a1dc46d3a) feat(*): omit symbol and DWARF symbol tables (#70)
* [`0f0ff02`](https://github.com/talos-systems/conform/commit/0f0ff02f6b134d2249f90c3698937b62222bfebd) refactor(policy): start policy error report at 1 (#68)
* [`4aaf049`](https://github.com/talos-systems/conform/commit/4aaf04924064dfcaaba3d984666f1fd4f4dbe4fe) chore(*): go module cleanup (#67)
* [`fbc195c`](https://github.com/talos-systems/conform/commit/fbc195c8055a61bc2f45cf95f29957aea2e429b3) chore(*): disable cgo (#66)
* [`155b036`](https://github.com/talos-systems/conform/commit/155b0369ddf5bea3b162067b804a487fb3b634e2) feat(metadata): mark SHA as dirty (#63)
* [`f39b434`](https://github.com/talos-systems/conform/commit/f39b4343fd0181dbb1ffca2c36573e6724768764) chore(*): always push latest (#65)
* [`1276371`](https://github.com/talos-systems/conform/commit/12763710cbe120102f5da965524a84b2b4d57bf1) chore(*): go modules and faster linting (#64)
* [`0a0cba3`](https://github.com/talos-systems/conform/commit/0a0cba34137d39002bf762b52d19121290be9980) feat(*): add build command (#62)
* [`aed2c22`](https://github.com/talos-systems/conform/commit/aed2c223188bb7ee04046976a4e1249e1451d8ca) fix(policy): update regex to allow optional scope (#61)
* [`1933d19`](https://github.com/talos-systems/conform/commit/1933d192d81bfb67150be634905179bf4fe183e1) fix(pipeline): nil pointer when no defined pipeline (#60)
* [`7fd1e89`](https://github.com/talos-systems/conform/commit/7fd1e89c567df89bbc16449182ec10a9ce2e5c18) feat(fmt): add fmt command (#59)
* [`42c5d09`](https://github.com/talos-systems/conform/commit/42c5d09d8189e222a5785232f98b000bc869228f) feat(docker): expose the image name and tag separately (#58)
* [`5a73ea6`](https://github.com/talos-systems/conform/commit/5a73ea6ed7490f91995e9fb916213226ac50fb01) feat(git): recursively search for .git in parent directories (#56)
* [`1e57715`](https://github.com/talos-systems/conform/commit/1e577157c87a76461c8576a7f6b696ab3704c53f) fix(cli): invalid variable message (#53)
* [`72061b1`](https://github.com/talos-systems/conform/commit/72061b11a121bd6ba1ad6e25a0e66059a09fc5d8) feat(metadata): allow users to specify variables (#52)
* [`4b7b903`](https://github.com/talos-systems/conform/commit/4b7b903aac3671f839115f6240d4e12e5256e742) fix(policy): strip leading newline from commit message (#50)
* [`779bf93`](https://github.com/talos-systems/conform/commit/779bf930fb4cdbe2d94b39c2218f9b2854efaf7a) fix(pipeline): don't show stdout of artifact extraction (#49)
* [`0373fea`](https://github.com/talos-systems/conform/commit/0373fea42e663c91e512fe1c144721b299bae457) feat(*): services, skip flag, and UX improvements (#43)
* [`c53f523`](https://github.com/talos-systems/conform/commit/c53f52307152a6392835e23797f1432799a1e0cc) feat(renderer): allow templates to be retrieved from URL (#41)
* [`c2cb181`](https://github.com/talos-systems/conform/commit/c2cb1818f44f6b5d0dd73b2e0de64c70972fd72e) refactor(*): minor changes (#37)
* [`19e4656`](https://github.com/talos-systems/conform/commit/19e4656da53bc707af536efa49ceae4edc82de8c) refactor(*): add enforcer package (#33)
* [`26ff570`](https://github.com/talos-systems/conform/commit/26ff570fbe76d82ac5e571eb709efa9e9459c99c) chore(ci): use autonomy/golang:1.8.3 for build and test tasks (#32)
* [`44d4ef0`](https://github.com/talos-systems/conform/commit/44d4ef0256dd48965eabcbeade2838bb9ccdeaea) fix(policy): check the entire commit header length (#31)
* [`bd3404a`](https://github.com/talos-systems/conform/commit/bd3404a593d9e891a2c70aadf25add80e7a71ae5) refactor(*): rename conform directory to pkg (#30)
* [`daa39f3`](https://github.com/talos-systems/conform/commit/daa39f36a0376e1e4767da5e43d13f0ec1b2c838) docs(readme): update example (#28)
* [`9383d3e`](https://github.com/talos-systems/conform/commit/9383d3ebab5b14c186fc84299c76cfa4d404a9c6) feat(policy): enforce 72 character limit on commit header (#29)
* [`6a115cf`](https://github.com/talos-systems/conform/commit/6a115cfbe5a750b8a4aecea0467adf258fa25e74) chore(ci): check if not a PR (#27)
* [`ec732ae`](https://github.com/talos-systems/conform/commit/ec732ae306598fd292c3b0e8ce38359d0da71aae) chore(ci): build on master or tag (#26)
* [`ba44e03`](https://github.com/talos-systems/conform/commit/ba44e03d35ce19ce0372970609110c46aa6be7ed) chore(ci): skip branch check in script (#25)
* [`6cece44`](https://github.com/talos-systems/conform/commit/6cece4467af854f47b5066a226ac43e4f8c60ecb) chore(ci): use Travis deploy (#24)
* [`31f79af`](https://github.com/talos-systems/conform/commit/31f79af8dc091d9bcb7ef9f2e3b6647ab12ef553) chore(ci): update script to print useful information (#23)
* [`a18332a`](https://github.com/talos-systems/conform/commit/a18332af48f1e5ddb5308ac9ed3d1ff030e96a1b) fix(metadata): nil version struct (#22)
* [`3ae8e5f`](https://github.com/talos-systems/conform/commit/3ae8e5f0f333d1bc44e31fb256448f36f4619332) refactor(*): make conform.yaml a dotfile (#21)
* [`936b64e`](https://github.com/talos-systems/conform/commit/936b64e1b6169a28ab254455df3457366d04a5fa) refactor(*): complete rewrite (#20)
* [`03caad0`](https://github.com/talos-systems/conform/commit/03caad0cb1a02e8bf688557f21c3bd58b2c69fdc) feat(policy): add policy enforcement; enforce git commit policy (#17)
* [`9927a05`](https://github.com/talos-systems/conform/commit/9927a05dc80f0ce8eb7a4cafaf49efc1db28cc57) refactor(docker): read Dockerfile from stdin (#16)
* [`dcc9fe5`](https://github.com/talos-systems/conform/commit/dcc9fe5417f820fe4c2b1673b41f7022d536c80d) chore(ci): use the stable Docker repository (#15)
* [`d37461a`](https://github.com/talos-systems/conform/commit/d37461ae38f87d0701dfe492ef497cef321b26f9) refactor(git): use go-git instead of shelled out commands (#14)
* [`71fa116`](https://github.com/talos-systems/conform/commit/71fa116a49757e7357467c7e21bac29ceab86beb) Add pre-release to git info (#12)
* [`fd5c627`](https://github.com/talos-systems/conform/commit/fd5c62790011ef68bf96aa7720e88839c61c1c11) Use gometalinter and fix linting errors (#11)
* [`0e66ba1`](https://github.com/talos-systems/conform/commit/0e66ba1b6fa9e60296fb0276d18c13d1bad88d10) Fix deploy on tags (#10)
* [`00fbfa8`](https://github.com/talos-systems/conform/commit/00fbfa845376925c51ead7c723ce4f5401da9df8) Use generic language in Travis build (#9)
* [`925dabf`](https://github.com/talos-systems/conform/commit/925dabfde80372ee952da4d4584f1e5fbf60c087) Remove 'version' from path of variables set at build time (#8)
* [`aa8ced7`](https://github.com/talos-systems/conform/commit/aa8ced7d725629190bba8e47f29bb24bb1bae3ad) Fix package path of variables set at build time (#7)
* [`ad8eef9`](https://github.com/talos-systems/conform/commit/ad8eef9901a2a1800751361453112ba5766cde6b) Fix copy of artifact in image template (#6)
* [`6311568`](https://github.com/talos-systems/conform/commit/6311568530d96447d9f6b1de673a777532fe6fd6) Set execution bit of deploy script (#5)
* [`92643d5`](https://github.com/talos-systems/conform/commit/92643d5fd3e0a46a4e937b6ec947b26f06f043d1) Stream script output and deploy on master or tags (#4)
* [`81055ed`](https://github.com/talos-systems/conform/commit/81055ed01af331bf8d5c5596cf8dfe26f608e532) Return script output on error (#3)
* [`426abe1`](https://github.com/talos-systems/conform/commit/426abe1aea53e396d5a2da260d71f75372344cc1) Fix bad tag detection and setup CI (#2)
* [`0c55035`](https://github.com/talos-systems/conform/commit/0c55035ee0f2129a4cb63d07da834dc2210684f1) Initial implementation (#1)
* [`994ba0b`](https://github.com/talos-systems/conform/commit/994ba0b98618d07e29e21092f6929ac9399e6fd5) Initial commit
</p>
</details>

### Changes since v0.1.0-alpha.24
<details><summary>1 commit</summary>
<p>

* [`18107a6`](https://github.com/talos-systems/conform/commit/18107a6d00cf30d420c6ef3a5c14f64efa4f385b) fix: checkout a branch only once
</p>
</details>

### Dependency Changes

This release has no dependency changes

## [conform 0.1.0-alpha.24](https://github.com/talos-systems/conform/releases/tag/v0.1.0-alpha.24) (2021-12-17)

Welcome to the v0.1.0-alpha.24 release of conform!



Please try out the release binaries and report any issues at
https://github.com/talos-systems/conform/issues.

### Contributors

* Andrew Rynhard
* Andrew Rynhard
* Andrey Smirnov
* Andrey Smirnov
* Joey Espinosa
* Alexey Palazhchenko
* Andrey Smirnov
* Peter Ng
* Wouter Dullaert
* Alex Szakaly
* André Valentin
* Anton Isakov
* Brad Beam
* Danny Zhu
* Dave Forgac
* Don Bowman
* Per Abich

### Changes
<details><summary>155 commits</summary>
<p>

* [`e9bb88a`](https://github.com/talos-systems/conform/commit/e9bb88a607115a6ff0008aec919f672f981038f2) fix: don't try git checkout repository twice
* [`bb8186b`](https://github.com/talos-systems/conform/commit/bb8186be5f8d6688284621c00c04d7cec9885091) fix: update jira issue valid message
* [`785d27d`](https://github.com/talos-systems/conform/commit/785d27dc5487aa6c5b5709a898d4e982c7f88410) release(v0.1.0-alpha.23): prepare release
* [`6caaabc`](https://github.com/talos-systems/conform/commit/6caaabc0f9721d161e68593fe3ee776761bdc893) feat: allow comments before the license header
* [`9a470a1`](https://github.com/talos-systems/conform/commit/9a470a16fe05643de83db40bf3f8cb0c3452ee2f) docs: update README with the latest changes
* [`35fba60`](https://github.com/talos-systems/conform/commit/35fba6014b3ecdd155c82fd931051b90c2aa9c35) chore: use kres to setup build instructions
* [`fb43bd4`](https://github.com/talos-systems/conform/commit/fb43bd4b31cca4206f0469b1afffcb1666685ab2) feat: provide gpg identity signature check
* [`c23e2fc`](https://github.com/talos-systems/conform/commit/c23e2fc7c6159cdd3ff6b78eb7591c654171e1a8) feat: add conventional commit breaking change support
* [`0e3a28c`](https://github.com/talos-systems/conform/commit/0e3a28c994aa9a0727d48bcde91c39034a2056e1) fix: update action version
* [`0390165`](https://github.com/talos-systems/conform/commit/03901655927737cedf5c0e901f72832cde9ee03a) fix: change check_jira regexp
* [`73f334a`](https://github.com/talos-systems/conform/commit/73f334a4ac7e6a799bc71f0611674f26b9b9c3a9) fix: disallow 0 as valid jira issue id
* [`5b58100`](https://github.com/talos-systems/conform/commit/5b58100ec8619aa94eb855546ddbeb8f27a35a24) chore: bump golangci-lint and fix linting issues
* [`c5dc2e6`](https://github.com/talos-systems/conform/commit/c5dc2e6b86fa76eea7fbf37a064ca6b9962aee96) fix: change "jiraKeys" to "keys"
* [`efd7fbb`](https://github.com/talos-systems/conform/commit/efd7fbb0dbc0042f54af587ffa07a829fdc88f1c) feat: improve Jira check
* [`001de56`](https://github.com/talos-systems/conform/commit/001de5691ecce630808bc2716229d9ee700a1ab4) feat: add support for requiring a Jira issue in the header
* [`5a75e96`](https://github.com/talos-systems/conform/commit/5a75e96171ef6aaaebecfced193df379296e03d7) chore: remove gitmeta
* [`cea1ee9`](https://github.com/talos-systems/conform/commit/cea1ee90f257daa96aa09ef9ba8b1deafddf92b2) chore: bump golangci-lint and fix linting issues
* [`6e0c294`](https://github.com/talos-systems/conform/commit/6e0c294bc044bde8dc43d45e31dee3b41a97fd53) feat: implement full gitignore-style pattern matching
* [`72708f2`](https://github.com/talos-systems/conform/commit/72708f25b12786e10a7f49da584a638ac76bfe92) feat: support regex in conventional commit scope
* [`ec5e365`](https://github.com/talos-systems/conform/commit/ec5e3656494068721fa29a815ea8abf67a4922f6) fix: checkout pull requests
* [`6d1a620`](https://github.com/talos-systems/conform/commit/6d1a620acbe83af09a97468129d02a366aec9b24) fix: use text/template
* [`8212dc6`](https://github.com/talos-systems/conform/commit/8212dc6b2353af86d0f69fd0fe4458b23507a7ec) chore: update CHANGELOG
* [`b3b6e65`](https://github.com/talos-systems/conform/commit/b3b6e657fbd4c94f57b81ac95885339ff4ee1b59) feat: add spellcheck commit check
* [`8726189`](https://github.com/talos-systems/conform/commit/8726189af55c5278de18b464273b89fed024aa2a) chore: update go-git
* [`f5826e5`](https://github.com/talos-systems/conform/commit/f5826e55f29fb34194cee5233b5621501a4c2171) refactor: rename summarizer to reporter
* [`b2f63c1`](https://github.com/talos-systems/conform/commit/b2f63c178998589d9f6f0273842ca38cc425a84d) feat: move body options to dedicated field
* [`fa7df19`](https://github.com/talos-systems/conform/commit/fa7df19996ece307285da44c73f210c6cbec9207) feat(policy): add checks for header case and last character
* [`a55d411`](https://github.com/talos-systems/conform/commit/a55d411e8f9061b6ab5b8dacd96c03254f57a557) docs: update README to include pre-commit compatibility
* [`ea7199a`](https://github.com/talos-systems/conform/commit/ea7199a68e686a6b6aea9a2c7741556c89680bd1) feat: add compatibility with pre-commit.com
* [`d97b22c`](https://github.com/talos-systems/conform/commit/d97b22c1bb88575b9e7a19a33dd47f8885096db8) feat(policy): Make Conventional commit description configurable
* [`59f365a`](https://github.com/talos-systems/conform/commit/59f365afe2384aabeee5617f28ef56b7f1ed6e17) fix: Correctly return errors from command run function
* [`01f87b9`](https://github.com/talos-systems/conform/commit/01f87b956d3b317e834713a918acaa3726fc7980) docs: add installation guide to README
* [`4eb4beb`](https://github.com/talos-systems/conform/commit/4eb4beb060332b18c585161728f77aee84b46783) fix: add action.yml
* [`d9668e0`](https://github.com/talos-systems/conform/commit/d9668e05ebde8f8c3bc96abf9711665c040efb6f) fix: remove autonomy references
* [`cf1213a`](https://github.com/talos-systems/conform/commit/cf1213a49ab9e3c6682d6ad8bd12fb71a8b5effc) fix: address conform errors
* [`5637edd`](https://github.com/talos-systems/conform/commit/5637edd03905d2d7612b0a86c7ac51aee7ae00b6) feat: add server mode
* [`0076446`](https://github.com/talos-systems/conform/commit/007644608f2455688d194164d2d29a3fa58a9918) chore: Replace autonomy with talos-systems
* [`7bed912`](https://github.com/talos-systems/conform/commit/7bed9129bc73b5a6fd2b6d8c12f7c024ea4b7107) fix: checks conv commit format in firstWord
* [`365c592`](https://github.com/talos-systems/conform/commit/365c592227c7ecadf52a8d22db22fca0dea015a0) fix: set pass to false if errors in policies
* [`19dd2b8`](https://github.com/talos-systems/conform/commit/19dd2b82ab96a65ab739ff00c5d7331b1293e6ce) chore: set docker server entrypoint to dockerd to avoid TLS generation
* [`25d013c`](https://github.com/talos-systems/conform/commit/25d013c905306829da13ed14189e994bac5a8137) docs: update README
* [`3b208c1`](https://github.com/talos-systems/conform/commit/3b208c1f665672ce64673b7a89a159de67d99f85) fix: print an error message
* [`4716823`](https://github.com/talos-systems/conform/commit/4716823c79156870677e239cf229b9af3f28273d) chore: require a commit body
* [`7809e90`](https://github.com/talos-systems/conform/commit/7809e90da2a0c4a087ea83a4c2a684f77f1dfe76) chore: push images for all branches
* [`0f7ec4c`](https://github.com/talos-systems/conform/commit/0f7ec4cfec40ea0ee377c56e4d356898064df1a4) chore: fix container labels
* [`adce353`](https://github.com/talos-systems/conform/commit/adce353bae54580a95a302681cba96b2f1706a36) feat: add commit body check
* [`e66888a`](https://github.com/talos-systems/conform/commit/e66888aeba8179a34646806ec050f348ba4a7669) feat: add number of commits check
* [`b330410`](https://github.com/talos-systems/conform/commit/b33041034a6b857f140f9d1b8485da3dc4e0a939) chore: build image with host net to avoid apk hang (#133)
* [`9e94c43`](https://github.com/talos-systems/conform/commit/9e94c43cfe66aff074a8c106899104a60a64272f) chore: prepare release v0.1.0-alpha.14 (#132)
* [`ef30db9`](https://github.com/talos-systems/conform/commit/ef30db9b206a4f67cc867eff4a510bbfd2a30a2d) fix: add file header check (#131)
* [`cc97536`](https://github.com/talos-systems/conform/commit/cc975363a9254c4878cd9a20708daf523bb576cf) feat: add support for GH actions on forked repo PRs (#130)
* [`4447684`](https://github.com/talos-systems/conform/commit/4447684a4433d321f1586f90bc2d107e6e7bcb91) chore: prepare release v0.1.0-alpha.12 (#129)
* [`3be1319`](https://github.com/talos-systems/conform/commit/3be1319605a0ee934a5fdf5a0e5a050f3a7e2579) feat: add support for github status checks (#128)
* [`6f8751c`](https://github.com/talos-systems/conform/commit/6f8751cb0791a8aeeb00194db3dc9c11059c7922) feat: add checks interface (#127)
* [`0af31f8`](https://github.com/talos-systems/conform/commit/0af31f88a74b836c6c8da15504730785f7b15ee8) fix: trim whitespace while validating DCO (#126)
* [`57c9dbd`](https://github.com/talos-systems/conform/commit/57c9dbd056d607fa0be2c5e1eb4628900ef85c2c) chore: quote docker creds (#122)
* [`ebed4b3`](https://github.com/talos-systems/conform/commit/ebed4b31cc2e6f10914850ffba9cd73fca803333) feat: implement `skipPaths` option for 'license' policy (#121)
* [`c539351`](https://github.com/talos-systems/conform/commit/c5393510751f9ba440e01845eae43d423970a16b) fix: excludeSuffixes wasn't skipping any files (#120)
* [`37e0e69`](https://github.com/talos-systems/conform/commit/37e0e6973100e596f29bbf43675472c5d8236679) fix: check empty commit-msg prior to parsing (#118)
* [`1473b44`](https://github.com/talos-systems/conform/commit/1473b4462de868edcca18bb2dfd5108f0545232e) feat: change the license header to a string (#116)
* [`abfd427`](https://github.com/talos-systems/conform/commit/abfd427a40a1eb8f2ee953e2442afd353625e460) chore: prepare v0.1.0-alpha.10 release (#115)
* [`bd039e4`](https://github.com/talos-systems/conform/commit/bd039e43fde1298a388c3c75e982b08a9610c98d) fix: use mood instead of verb (#114)
* [`6ac7c2f`](https://github.com/talos-systems/conform/commit/6ac7c2f640fbcf85419fb9914b8bdcccd71570c0) fix: ensure the imperative check is against lowercase word (#112)
* [`fd6ad6c`](https://github.com/talos-systems/conform/commit/fd6ad6cdb0746bc5547b5bfc5d474c9ab84d68f7) chore: prepare v0.1.0-alpha.9 release (#111)
* [`286041a`](https://github.com/talos-systems/conform/commit/286041a7b48aa8b789bbab398de3a85f02ba87bd) docs: update README (#110)
* [`3f75846`](https://github.com/talos-systems/conform/commit/3f758468cea5db94fdd897dca4bc8c98016a5089) fix(policy): use natural language processing for imperative check (#109)
* [`5c6620a`](https://github.com/talos-systems/conform/commit/5c6620a1f544d9f3dd11cf5092efd698dc260827) feat(policy): add imperative mood check (#108)
* [`86a7d3e`](https://github.com/talos-systems/conform/commit/86a7d3e57de33b2800c5feb1deb4655d51fc151d) docs: fix code highlight in README (#107)
* [`eeb3d5c`](https://github.com/talos-systems/conform/commit/eeb3d5ce7ff602af855de37ec5f30a34f45e550c) docs: move LICENSE_HEADER to root of project (#106)
* [`f5ed717`](https://github.com/talos-systems/conform/commit/f5ed7174d6d5019e322d1acd4e9de47e1064a4f9) feat: add license header policy (#105)
* [`763d4d9`](https://github.com/talos-systems/conform/commit/763d4d94587a7000ebd67d7277d38275179184b0) docs: add conventional commits badge to README (#104)
* [`e4602b8`](https://github.com/talos-systems/conform/commit/e4602b810eace5fe01ef191d58026a60e448e9c6) chore: prepare v0.1.0-alpha.8 release (#103)
* [`116a3bf`](https://github.com/talos-systems/conform/commit/116a3bf1bd5cf6dc6026b5d3f5fd09640f67380e) fix(policy): remove commit header length from conventional commit policy (#102)
* [`2be1e1e`](https://github.com/talos-systems/conform/commit/2be1e1e7eaf5827c960d747f831a376c10a433ce) docs: update README (#101)
* [`22804ff`](https://github.com/talos-systems/conform/commit/22804ff48efae434dc3bfc7b1e775c747d5772da) chore: fix image tag (#100)
* [`e6664a9`](https://github.com/talos-systems/conform/commit/e6664a9705169e76673632bbb6c68b23b4b1194b) chore: prepare v0.1.0-alpha.7 release (#99)
* [`7646221`](https://github.com/talos-systems/conform/commit/7646221dc581b026654efdc9617c6eac76bb09a9) feat: output status in tab format (#98)
* [`e93a47e`](https://github.com/talos-systems/conform/commit/e93a47ee759d0345e9a2ad3526a64dcc1715f366) chore: push latest tag (#97)
* [`598b595`](https://github.com/talos-systems/conform/commit/598b595b47df3e87c7eae5fd997c3f01d2af859f) chore: revert base image to scratch (#96)
* [`eb6cc6d`](https://github.com/talos-systems/conform/commit/eb6cc6db46c88ecfe2099288c1aa5826185d1575) chore: add conform binary to /bin (#95)
* [`fd41a2f`](https://github.com/talos-systems/conform/commit/fd41a2f3619e35f9770dfddfc00c2f0eb94637bd) chore: use alpine:3.8 as base image (#94)
* [`36df035`](https://github.com/talos-systems/conform/commit/36df0355239e6fee6767b3b2e536d976ce76bfdd) chore: use buildkit for builds (#93)
* [`b59ae9c`](https://github.com/talos-systems/conform/commit/b59ae9c6fd5482b2558d43186b2951d34b7c6c40) feat: Add generic git commit policy (#92)
* [`76b6d7a`](https://github.com/talos-systems/conform/commit/76b6d7a2a7346c5d09de14bec6cb757f4a0ddb9d) chore: fix image push (#91)
* [`22e0f7b`](https://github.com/talos-systems/conform/commit/22e0f7bb7bf6ca5face6f75036fbe54407090699) chore: pin Kaniko to v0.6.0 (#90)
* [`9b72c17`](https://github.com/talos-systems/conform/commit/9b72c1797945a00cba940df4e47fa0fa8718f9ce) docs: fix README example (#89)
* [`5543c79`](https://github.com/talos-systems/conform/commit/5543c7908afe1d205234b3acefe883df2b359b4f) chore: fix typos (#88)
* [`349ba37`](https://github.com/talos-systems/conform/commit/349ba37ba4957efc1f99e24b5c34cb664dd0aa5e) docs: fix README (#87)
* [`8ced588`](https://github.com/talos-systems/conform/commit/8ced58895fb10576407e85edfadfbd568861e6ec) refactor: use Kaniko for builds (#86)
* [`0fdd552`](https://github.com/talos-systems/conform/commit/0fdd552a27d573cf9ecbc5e75b51b2e71fabcb09) feat: show git status (#85)
* [`76217df`](https://github.com/talos-systems/conform/commit/76217df7c54752b6805fade67a548ec45c502919) feat: remove artifacts before creating (#84)
* [`3b165b3`](https://github.com/talos-systems/conform/commit/3b165b3f8d332bb9e0c16d6e43ffeb923f5acbf7) fix(metadata): keep original version string (#82)
* [`f27917e`](https://github.com/talos-systems/conform/commit/f27917e150dc0b6f14d31a75f209457f8fa94889) fix(ci): push built images images (#83)
* [`e67dd43`](https://github.com/talos-systems/conform/commit/e67dd4363be3567a4094949ccf0d3263d8aa8dcc) chore(ci): add brigade configuration (#80)
* [`18d8905`](https://github.com/talos-systems/conform/commit/18d8905104b7db6bd05d33cbcfed65c5412cf16e) feat(metadata): add git ref (#81)
* [`5caf3b5`](https://github.com/talos-systems/conform/commit/5caf3b5c99fb0cd164b2a1f60d9eef86969bc384) feat(metadata): add original semver string (#79)
* [`d65491a`](https://github.com/talos-systems/conform/commit/d65491ab160a33ccf166bfe3e1d4ba2596792f57) feat(policy): show valid types and scopes on error (#78)
* [`954c003`](https://github.com/talos-systems/conform/commit/954c00327ea1cb6d16c646ba3a97b5bd4f47170a) fix(policy): unit test inline git config (#77)
* [`e439cd7`](https://github.com/talos-systems/conform/commit/e439cd7bee060a0c16bc6d9ac23f4e87bc46da9e) chore(ci): run go mod tidy (#76)
* [`2ac5059`](https://github.com/talos-systems/conform/commit/2ac50599f3f5944c44fa0c45790481e8f9d532b5) fix(policy): return error in conventional commit report (#75)
* [`7d19c82`](https://github.com/talos-systems/conform/commit/7d19c82835585768a7601a1745a344a347f0ee9c) chore(ci): show git status when dirty working tree (#74)
* [`4a6cc1c`](https://github.com/talos-systems/conform/commit/4a6cc1cb9529ad561372c09fd7f579694c212b1b) feat: adding command line flag for commit msg (#73)
* [`088e0a7`](https://github.com/talos-systems/conform/commit/088e0a76b6e04ca521c6b9e03e66acb03b8dfbe8) chore(*): output binaries (#72)
* [`4194aa5`](https://github.com/talos-systems/conform/commit/4194aa5004222589c36265d4649452b72fb744bd) chore(*): format .conform.yaml (#71)
* [`05cfacb`](https://github.com/talos-systems/conform/commit/05cfacb6edd0dd92bd28bbb60c91970a1dc46d3a) feat(*): omit symbol and DWARF symbol tables (#70)
* [`0f0ff02`](https://github.com/talos-systems/conform/commit/0f0ff02f6b134d2249f90c3698937b62222bfebd) refactor(policy): start policy error report at 1 (#68)
* [`4aaf049`](https://github.com/talos-systems/conform/commit/4aaf04924064dfcaaba3d984666f1fd4f4dbe4fe) chore(*): go module cleanup (#67)
* [`fbc195c`](https://github.com/talos-systems/conform/commit/fbc195c8055a61bc2f45cf95f29957aea2e429b3) chore(*): disable cgo (#66)
* [`155b036`](https://github.com/talos-systems/conform/commit/155b0369ddf5bea3b162067b804a487fb3b634e2) feat(metadata): mark SHA as dirty (#63)
* [`f39b434`](https://github.com/talos-systems/conform/commit/f39b4343fd0181dbb1ffca2c36573e6724768764) chore(*): always push latest (#65)
* [`1276371`](https://github.com/talos-systems/conform/commit/12763710cbe120102f5da965524a84b2b4d57bf1) chore(*): go modules and faster linting (#64)
* [`0a0cba3`](https://github.com/talos-systems/conform/commit/0a0cba34137d39002bf762b52d19121290be9980) feat(*): add build command (#62)
* [`aed2c22`](https://github.com/talos-systems/conform/commit/aed2c223188bb7ee04046976a4e1249e1451d8ca) fix(policy): update regex to allow optional scope (#61)
* [`1933d19`](https://github.com/talos-systems/conform/commit/1933d192d81bfb67150be634905179bf4fe183e1) fix(pipeline): nil pointer when no defined pipeline (#60)
* [`7fd1e89`](https://github.com/talos-systems/conform/commit/7fd1e89c567df89bbc16449182ec10a9ce2e5c18) feat(fmt): add fmt command (#59)
* [`42c5d09`](https://github.com/talos-systems/conform/commit/42c5d09d8189e222a5785232f98b000bc869228f) feat(docker): expose the image name and tag separately (#58)
* [`5a73ea6`](https://github.com/talos-systems/conform/commit/5a73ea6ed7490f91995e9fb916213226ac50fb01) feat(git): recursively search for .git in parent directories (#56)
* [`1e57715`](https://github.com/talos-systems/conform/commit/1e577157c87a76461c8576a7f6b696ab3704c53f) fix(cli): invalid variable message (#53)
* [`72061b1`](https://github.com/talos-systems/conform/commit/72061b11a121bd6ba1ad6e25a0e66059a09fc5d8) feat(metadata): allow users to specify variables (#52)
* [`4b7b903`](https://github.com/talos-systems/conform/commit/4b7b903aac3671f839115f6240d4e12e5256e742) fix(policy): strip leading newline from commit message (#50)
* [`779bf93`](https://github.com/talos-systems/conform/commit/779bf930fb4cdbe2d94b39c2218f9b2854efaf7a) fix(pipeline): don't show stdout of artifact extraction (#49)
* [`0373fea`](https://github.com/talos-systems/conform/commit/0373fea42e663c91e512fe1c144721b299bae457) feat(*): services, skip flag, and UX improvements (#43)
* [`c53f523`](https://github.com/talos-systems/conform/commit/c53f52307152a6392835e23797f1432799a1e0cc) feat(renderer): allow templates to be retrieved from URL (#41)
* [`c2cb181`](https://github.com/talos-systems/conform/commit/c2cb1818f44f6b5d0dd73b2e0de64c70972fd72e) refactor(*): minor changes (#37)
* [`19e4656`](https://github.com/talos-systems/conform/commit/19e4656da53bc707af536efa49ceae4edc82de8c) refactor(*): add enforcer package (#33)
* [`26ff570`](https://github.com/talos-systems/conform/commit/26ff570fbe76d82ac5e571eb709efa9e9459c99c) chore(ci): use autonomy/golang:1.8.3 for build and test tasks (#32)
* [`44d4ef0`](https://github.com/talos-systems/conform/commit/44d4ef0256dd48965eabcbeade2838bb9ccdeaea) fix(policy): check the entire commit header length (#31)
* [`bd3404a`](https://github.com/talos-systems/conform/commit/bd3404a593d9e891a2c70aadf25add80e7a71ae5) refactor(*): rename conform directory to pkg (#30)
* [`daa39f3`](https://github.com/talos-systems/conform/commit/daa39f36a0376e1e4767da5e43d13f0ec1b2c838) docs(readme): update example (#28)
* [`9383d3e`](https://github.com/talos-systems/conform/commit/9383d3ebab5b14c186fc84299c76cfa4d404a9c6) feat(policy): enforce 72 character limit on commit header (#29)
* [`6a115cf`](https://github.com/talos-systems/conform/commit/6a115cfbe5a750b8a4aecea0467adf258fa25e74) chore(ci): check if not a PR (#27)
* [`ec732ae`](https://github.com/talos-systems/conform/commit/ec732ae306598fd292c3b0e8ce38359d0da71aae) chore(ci): build on master or tag (#26)
* [`ba44e03`](https://github.com/talos-systems/conform/commit/ba44e03d35ce19ce0372970609110c46aa6be7ed) chore(ci): skip branch check in script (#25)
* [`6cece44`](https://github.com/talos-systems/conform/commit/6cece4467af854f47b5066a226ac43e4f8c60ecb) chore(ci): use Travis deploy (#24)
* [`31f79af`](https://github.com/talos-systems/conform/commit/31f79af8dc091d9bcb7ef9f2e3b6647ab12ef553) chore(ci): update script to print useful information (#23)
* [`a18332a`](https://github.com/talos-systems/conform/commit/a18332af48f1e5ddb5308ac9ed3d1ff030e96a1b) fix(metadata): nil version struct (#22)
* [`3ae8e5f`](https://github.com/talos-systems/conform/commit/3ae8e5f0f333d1bc44e31fb256448f36f4619332) refactor(*): make conform.yaml a dotfile (#21)
* [`936b64e`](https://github.com/talos-systems/conform/commit/936b64e1b6169a28ab254455df3457366d04a5fa) refactor(*): complete rewrite (#20)
* [`03caad0`](https://github.com/talos-systems/conform/commit/03caad0cb1a02e8bf688557f21c3bd58b2c69fdc) feat(policy): add policy enforcement; enforce git commit policy (#17)
* [`9927a05`](https://github.com/talos-systems/conform/commit/9927a05dc80f0ce8eb7a4cafaf49efc1db28cc57) refactor(docker): read Dockerfile from stdin (#16)
* [`dcc9fe5`](https://github.com/talos-systems/conform/commit/dcc9fe5417f820fe4c2b1673b41f7022d536c80d) chore(ci): use the stable Docker repository (#15)
* [`d37461a`](https://github.com/talos-systems/conform/commit/d37461ae38f87d0701dfe492ef497cef321b26f9) refactor(git): use go-git instead of shelled out commands (#14)
* [`71fa116`](https://github.com/talos-systems/conform/commit/71fa116a49757e7357467c7e21bac29ceab86beb) Add pre-release to git info (#12)
* [`fd5c627`](https://github.com/talos-systems/conform/commit/fd5c62790011ef68bf96aa7720e88839c61c1c11) Use gometalinter and fix linting errors (#11)
* [`0e66ba1`](https://github.com/talos-systems/conform/commit/0e66ba1b6fa9e60296fb0276d18c13d1bad88d10) Fix deploy on tags (#10)
* [`00fbfa8`](https://github.com/talos-systems/conform/commit/00fbfa845376925c51ead7c723ce4f5401da9df8) Use generic language in Travis build (#9)
* [`925dabf`](https://github.com/talos-systems/conform/commit/925dabfde80372ee952da4d4584f1e5fbf60c087) Remove 'version' from path of variables set at build time (#8)
* [`aa8ced7`](https://github.com/talos-systems/conform/commit/aa8ced7d725629190bba8e47f29bb24bb1bae3ad) Fix package path of variables set at build time (#7)
* [`ad8eef9`](https://github.com/talos-systems/conform/commit/ad8eef9901a2a1800751361453112ba5766cde6b) Fix copy of artifact in image template (#6)
* [`6311568`](https://github.com/talos-systems/conform/commit/6311568530d96447d9f6b1de673a777532fe6fd6) Set execution bit of deploy script (#5)
* [`92643d5`](https://github.com/talos-systems/conform/commit/92643d5fd3e0a46a4e937b6ec947b26f06f043d1) Stream script output and deploy on master or tags (#4)
* [`81055ed`](https://github.com/talos-systems/conform/commit/81055ed01af331bf8d5c5596cf8dfe26f608e532) Return script output on error (#3)
* [`426abe1`](https://github.com/talos-systems/conform/commit/426abe1aea53e396d5a2da260d71f75372344cc1) Fix bad tag detection and setup CI (#2)
* [`0c55035`](https://github.com/talos-systems/conform/commit/0c55035ee0f2129a4cb63d07da834dc2210684f1) Initial implementation (#1)
* [`994ba0b`](https://github.com/talos-systems/conform/commit/994ba0b98618d07e29e21092f6929ac9399e6fd5) Initial commit
</p>
</details>

### Changes since v0.1.0-alpha.23
<details><summary>2 commits</summary>
<p>

* [`e9bb88a`](https://github.com/talos-systems/conform/commit/e9bb88a607115a6ff0008aec919f672f981038f2) fix: don't try git checkout repository twice
* [`bb8186b`](https://github.com/talos-systems/conform/commit/bb8186be5f8d6688284621c00c04d7cec9885091) fix: update jira issue valid message
</p>
</details>

### Dependency Changes

This release has no dependency changes

## [conform 0.1.0-alpha.23](https://github.com/talos-systems/conform/releases/tag/v0.1.0-alpha.23) (2021-09-06)

Welcome to the v0.1.0-alpha.23 release of conform!

Two new features were added since alpha.19:
* GPG signature identity check.
* Ability to ignore preceding comments in the license header check.

Please try out the release binaries and report any issues at
https://github.com/talos-systems/conform/issues.

### Contributors

* Andrew Rynhard
* Andrey Smirnov
* Joey Espinosa
* Peter Ng
* Wouter Dullaert
* Alex Szakaly
* Alexey Palazhchenko
* André Valentin
* Anton Isakov
* Brad Beam
* Danny Zhu
* Don Bowman
* Per Abich

### Changes
<details><summary>152 commits</summary>
<p>

* [`6caaabc`](https://github.com/talos-systems/conform/commit/6caaabc0f9721d161e68593fe3ee776761bdc893) feat: allow comments before the license header
* [`9a470a1`](https://github.com/talos-systems/conform/commit/9a470a16fe05643de83db40bf3f8cb0c3452ee2f) docs: update README with the latest changes
* [`35fba60`](https://github.com/talos-systems/conform/commit/35fba6014b3ecdd155c82fd931051b90c2aa9c35) chore: use kres to setup build instructions
* [`fb43bd4`](https://github.com/talos-systems/conform/commit/fb43bd4b31cca4206f0469b1afffcb1666685ab2) feat: provide gpg identity signature check
* [`c23e2fc`](https://github.com/talos-systems/conform/commit/c23e2fc7c6159cdd3ff6b78eb7591c654171e1a8) feat: add conventional commit breaking change support
* [`0e3a28c`](https://github.com/talos-systems/conform/commit/0e3a28c994aa9a0727d48bcde91c39034a2056e1) fix: update action version
* [`0390165`](https://github.com/talos-systems/conform/commit/03901655927737cedf5c0e901f72832cde9ee03a) fix: change check_jira regexp
* [`73f334a`](https://github.com/talos-systems/conform/commit/73f334a4ac7e6a799bc71f0611674f26b9b9c3a9) fix: disallow 0 as valid jira issue id
* [`5b58100`](https://github.com/talos-systems/conform/commit/5b58100ec8619aa94eb855546ddbeb8f27a35a24) chore: bump golangci-lint and fix linting issues
* [`c5dc2e6`](https://github.com/talos-systems/conform/commit/c5dc2e6b86fa76eea7fbf37a064ca6b9962aee96) fix: change "jiraKeys" to "keys"
* [`efd7fbb`](https://github.com/talos-systems/conform/commit/efd7fbb0dbc0042f54af587ffa07a829fdc88f1c) feat: improve Jira check
* [`001de56`](https://github.com/talos-systems/conform/commit/001de5691ecce630808bc2716229d9ee700a1ab4) feat: add support for requiring a Jira issue in the header
* [`5a75e96`](https://github.com/talos-systems/conform/commit/5a75e96171ef6aaaebecfced193df379296e03d7) chore: remove gitmeta
* [`cea1ee9`](https://github.com/talos-systems/conform/commit/cea1ee90f257daa96aa09ef9ba8b1deafddf92b2) chore: bump golangci-lint and fix linting issues
* [`6e0c294`](https://github.com/talos-systems/conform/commit/6e0c294bc044bde8dc43d45e31dee3b41a97fd53) feat: implement full gitignore-style pattern matching
* [`72708f2`](https://github.com/talos-systems/conform/commit/72708f25b12786e10a7f49da584a638ac76bfe92) feat: support regex in conventional commit scope
* [`ec5e365`](https://github.com/talos-systems/conform/commit/ec5e3656494068721fa29a815ea8abf67a4922f6) fix: checkout pull requests
* [`6d1a620`](https://github.com/talos-systems/conform/commit/6d1a620acbe83af09a97468129d02a366aec9b24) fix: use text/template
* [`8212dc6`](https://github.com/talos-systems/conform/commit/8212dc6b2353af86d0f69fd0fe4458b23507a7ec) chore: update CHANGELOG
* [`b3b6e65`](https://github.com/talos-systems/conform/commit/b3b6e657fbd4c94f57b81ac95885339ff4ee1b59) feat: add spellcheck commit check
* [`8726189`](https://github.com/talos-systems/conform/commit/8726189af55c5278de18b464273b89fed024aa2a) chore: update go-git
* [`f5826e5`](https://github.com/talos-systems/conform/commit/f5826e55f29fb34194cee5233b5621501a4c2171) refactor: rename summarizer to reporter
* [`b2f63c1`](https://github.com/talos-systems/conform/commit/b2f63c178998589d9f6f0273842ca38cc425a84d) feat: move body options to dedicated field
* [`fa7df19`](https://github.com/talos-systems/conform/commit/fa7df19996ece307285da44c73f210c6cbec9207) feat(policy): add checks for header case and last character
* [`a55d411`](https://github.com/talos-systems/conform/commit/a55d411e8f9061b6ab5b8dacd96c03254f57a557) docs: update README to include pre-commit compatibility
* [`ea7199a`](https://github.com/talos-systems/conform/commit/ea7199a68e686a6b6aea9a2c7741556c89680bd1) feat: add compatibility with pre-commit.com
* [`d97b22c`](https://github.com/talos-systems/conform/commit/d97b22c1bb88575b9e7a19a33dd47f8885096db8) feat(policy): Make Conventional commit description configurable
* [`59f365a`](https://github.com/talos-systems/conform/commit/59f365afe2384aabeee5617f28ef56b7f1ed6e17) fix: Correctly return errors from command run function
* [`01f87b9`](https://github.com/talos-systems/conform/commit/01f87b956d3b317e834713a918acaa3726fc7980) docs: add installation guide to README
* [`4eb4beb`](https://github.com/talos-systems/conform/commit/4eb4beb060332b18c585161728f77aee84b46783) fix: add action.yml
* [`d9668e0`](https://github.com/talos-systems/conform/commit/d9668e05ebde8f8c3bc96abf9711665c040efb6f) fix: remove autonomy references
* [`cf1213a`](https://github.com/talos-systems/conform/commit/cf1213a49ab9e3c6682d6ad8bd12fb71a8b5effc) fix: address conform errors
* [`5637edd`](https://github.com/talos-systems/conform/commit/5637edd03905d2d7612b0a86c7ac51aee7ae00b6) feat: add server mode
* [`0076446`](https://github.com/talos-systems/conform/commit/007644608f2455688d194164d2d29a3fa58a9918) chore: Replace autonomy with talos-systems
* [`7bed912`](https://github.com/talos-systems/conform/commit/7bed9129bc73b5a6fd2b6d8c12f7c024ea4b7107) fix: checks conv commit format in firstWord
* [`365c592`](https://github.com/talos-systems/conform/commit/365c592227c7ecadf52a8d22db22fca0dea015a0) fix: set pass to false if errors in policies
* [`19dd2b8`](https://github.com/talos-systems/conform/commit/19dd2b82ab96a65ab739ff00c5d7331b1293e6ce) chore: set docker server entrypoint to dockerd to avoid TLS generation
* [`25d013c`](https://github.com/talos-systems/conform/commit/25d013c905306829da13ed14189e994bac5a8137) docs: update README
* [`3b208c1`](https://github.com/talos-systems/conform/commit/3b208c1f665672ce64673b7a89a159de67d99f85) fix: print an error message
* [`4716823`](https://github.com/talos-systems/conform/commit/4716823c79156870677e239cf229b9af3f28273d) chore: require a commit body
* [`7809e90`](https://github.com/talos-systems/conform/commit/7809e90da2a0c4a087ea83a4c2a684f77f1dfe76) chore: push images for all branches
* [`0f7ec4c`](https://github.com/talos-systems/conform/commit/0f7ec4cfec40ea0ee377c56e4d356898064df1a4) chore: fix container labels
* [`adce353`](https://github.com/talos-systems/conform/commit/adce353bae54580a95a302681cba96b2f1706a36) feat: add commit body check
* [`e66888a`](https://github.com/talos-systems/conform/commit/e66888aeba8179a34646806ec050f348ba4a7669) feat: add number of commits check
* [`b330410`](https://github.com/talos-systems/conform/commit/b33041034a6b857f140f9d1b8485da3dc4e0a939) chore: build image with host net to avoid apk hang (#133)
* [`9e94c43`](https://github.com/talos-systems/conform/commit/9e94c43cfe66aff074a8c106899104a60a64272f) chore: prepare release v0.1.0-alpha.14 (#132)
* [`ef30db9`](https://github.com/talos-systems/conform/commit/ef30db9b206a4f67cc867eff4a510bbfd2a30a2d) fix: add file header check (#131)
* [`cc97536`](https://github.com/talos-systems/conform/commit/cc975363a9254c4878cd9a20708daf523bb576cf) feat: add support for GH actions on forked repo PRs (#130)
* [`4447684`](https://github.com/talos-systems/conform/commit/4447684a4433d321f1586f90bc2d107e6e7bcb91) chore: prepare release v0.1.0-alpha.12 (#129)
* [`3be1319`](https://github.com/talos-systems/conform/commit/3be1319605a0ee934a5fdf5a0e5a050f3a7e2579) feat: add support for github status checks (#128)
* [`6f8751c`](https://github.com/talos-systems/conform/commit/6f8751cb0791a8aeeb00194db3dc9c11059c7922) feat: add checks interface (#127)
* [`0af31f8`](https://github.com/talos-systems/conform/commit/0af31f88a74b836c6c8da15504730785f7b15ee8) fix: trim whitespace while validating DCO (#126)
* [`57c9dbd`](https://github.com/talos-systems/conform/commit/57c9dbd056d607fa0be2c5e1eb4628900ef85c2c) chore: quote docker creds (#122)
* [`ebed4b3`](https://github.com/talos-systems/conform/commit/ebed4b31cc2e6f10914850ffba9cd73fca803333) feat: implement `skipPaths` option for 'license' policy (#121)
* [`c539351`](https://github.com/talos-systems/conform/commit/c5393510751f9ba440e01845eae43d423970a16b) fix: excludeSuffixes wasn't skipping any files (#120)
* [`37e0e69`](https://github.com/talos-systems/conform/commit/37e0e6973100e596f29bbf43675472c5d8236679) fix: check empty commit-msg prior to parsing (#118)
* [`1473b44`](https://github.com/talos-systems/conform/commit/1473b4462de868edcca18bb2dfd5108f0545232e) feat: change the license header to a string (#116)
* [`abfd427`](https://github.com/talos-systems/conform/commit/abfd427a40a1eb8f2ee953e2442afd353625e460) chore: prepare v0.1.0-alpha.10 release (#115)
* [`bd039e4`](https://github.com/talos-systems/conform/commit/bd039e43fde1298a388c3c75e982b08a9610c98d) fix: use mood instead of verb (#114)
* [`6ac7c2f`](https://github.com/talos-systems/conform/commit/6ac7c2f640fbcf85419fb9914b8bdcccd71570c0) fix: ensure the imperative check is against lowercase word (#112)
* [`fd6ad6c`](https://github.com/talos-systems/conform/commit/fd6ad6cdb0746bc5547b5bfc5d474c9ab84d68f7) chore: prepare v0.1.0-alpha.9 release (#111)
* [`286041a`](https://github.com/talos-systems/conform/commit/286041a7b48aa8b789bbab398de3a85f02ba87bd) docs: update README (#110)
* [`3f75846`](https://github.com/talos-systems/conform/commit/3f758468cea5db94fdd897dca4bc8c98016a5089) fix(policy): use natural language processing for imperative check (#109)
* [`5c6620a`](https://github.com/talos-systems/conform/commit/5c6620a1f544d9f3dd11cf5092efd698dc260827) feat(policy): add imperative mood check (#108)
* [`86a7d3e`](https://github.com/talos-systems/conform/commit/86a7d3e57de33b2800c5feb1deb4655d51fc151d) docs: fix code highlight in README (#107)
* [`eeb3d5c`](https://github.com/talos-systems/conform/commit/eeb3d5ce7ff602af855de37ec5f30a34f45e550c) docs: move LICENSE_HEADER to root of project (#106)
* [`f5ed717`](https://github.com/talos-systems/conform/commit/f5ed7174d6d5019e322d1acd4e9de47e1064a4f9) feat: add license header policy (#105)
* [`763d4d9`](https://github.com/talos-systems/conform/commit/763d4d94587a7000ebd67d7277d38275179184b0) docs: add conventional commits badge to README (#104)
* [`e4602b8`](https://github.com/talos-systems/conform/commit/e4602b810eace5fe01ef191d58026a60e448e9c6) chore: prepare v0.1.0-alpha.8 release (#103)
* [`116a3bf`](https://github.com/talos-systems/conform/commit/116a3bf1bd5cf6dc6026b5d3f5fd09640f67380e) fix(policy): remove commit header length from conventional commit policy (#102)
* [`2be1e1e`](https://github.com/talos-systems/conform/commit/2be1e1e7eaf5827c960d747f831a376c10a433ce) docs: update README (#101)
* [`22804ff`](https://github.com/talos-systems/conform/commit/22804ff48efae434dc3bfc7b1e775c747d5772da) chore: fix image tag (#100)
* [`e6664a9`](https://github.com/talos-systems/conform/commit/e6664a9705169e76673632bbb6c68b23b4b1194b) chore: prepare v0.1.0-alpha.7 release (#99)
* [`7646221`](https://github.com/talos-systems/conform/commit/7646221dc581b026654efdc9617c6eac76bb09a9) feat: output status in tab format (#98)
* [`e93a47e`](https://github.com/talos-systems/conform/commit/e93a47ee759d0345e9a2ad3526a64dcc1715f366) chore: push latest tag (#97)
* [`598b595`](https://github.com/talos-systems/conform/commit/598b595b47df3e87c7eae5fd997c3f01d2af859f) chore: revert base image to scratch (#96)
* [`eb6cc6d`](https://github.com/talos-systems/conform/commit/eb6cc6db46c88ecfe2099288c1aa5826185d1575) chore: add conform binary to /bin (#95)
* [`fd41a2f`](https://github.com/talos-systems/conform/commit/fd41a2f3619e35f9770dfddfc00c2f0eb94637bd) chore: use alpine:3.8 as base image (#94)
* [`36df035`](https://github.com/talos-systems/conform/commit/36df0355239e6fee6767b3b2e536d976ce76bfdd) chore: use buildkit for builds (#93)
* [`b59ae9c`](https://github.com/talos-systems/conform/commit/b59ae9c6fd5482b2558d43186b2951d34b7c6c40) feat: Add generic git commit policy (#92)
* [`76b6d7a`](https://github.com/talos-systems/conform/commit/76b6d7a2a7346c5d09de14bec6cb757f4a0ddb9d) chore: fix image push (#91)
* [`22e0f7b`](https://github.com/talos-systems/conform/commit/22e0f7bb7bf6ca5face6f75036fbe54407090699) chore: pin Kaniko to v0.6.0 (#90)
* [`9b72c17`](https://github.com/talos-systems/conform/commit/9b72c1797945a00cba940df4e47fa0fa8718f9ce) docs: fix README example (#89)
* [`5543c79`](https://github.com/talos-systems/conform/commit/5543c7908afe1d205234b3acefe883df2b359b4f) chore: fix typos (#88)
* [`349ba37`](https://github.com/talos-systems/conform/commit/349ba37ba4957efc1f99e24b5c34cb664dd0aa5e) docs: fix README (#87)
* [`8ced588`](https://github.com/talos-systems/conform/commit/8ced58895fb10576407e85edfadfbd568861e6ec) refactor: use Kaniko for builds (#86)
* [`0fdd552`](https://github.com/talos-systems/conform/commit/0fdd552a27d573cf9ecbc5e75b51b2e71fabcb09) feat: show git status (#85)
* [`76217df`](https://github.com/talos-systems/conform/commit/76217df7c54752b6805fade67a548ec45c502919) feat: remove artifacts before creating (#84)
* [`3b165b3`](https://github.com/talos-systems/conform/commit/3b165b3f8d332bb9e0c16d6e43ffeb923f5acbf7) fix(metadata): keep original version string (#82)
* [`f27917e`](https://github.com/talos-systems/conform/commit/f27917e150dc0b6f14d31a75f209457f8fa94889) fix(ci): push built images images (#83)
* [`e67dd43`](https://github.com/talos-systems/conform/commit/e67dd4363be3567a4094949ccf0d3263d8aa8dcc) chore(ci): add brigade configuration (#80)
* [`18d8905`](https://github.com/talos-systems/conform/commit/18d8905104b7db6bd05d33cbcfed65c5412cf16e) feat(metadata): add git ref (#81)
* [`5caf3b5`](https://github.com/talos-systems/conform/commit/5caf3b5c99fb0cd164b2a1f60d9eef86969bc384) feat(metadata): add original semver string (#79)
* [`d65491a`](https://github.com/talos-systems/conform/commit/d65491ab160a33ccf166bfe3e1d4ba2596792f57) feat(policy): show valid types and scopes on error (#78)
* [`954c003`](https://github.com/talos-systems/conform/commit/954c00327ea1cb6d16c646ba3a97b5bd4f47170a) fix(policy): unit test inline git config (#77)
* [`e439cd7`](https://github.com/talos-systems/conform/commit/e439cd7bee060a0c16bc6d9ac23f4e87bc46da9e) chore(ci): run go mod tidy (#76)
* [`2ac5059`](https://github.com/talos-systems/conform/commit/2ac50599f3f5944c44fa0c45790481e8f9d532b5) fix(policy): return error in conventional commit report (#75)
* [`7d19c82`](https://github.com/talos-systems/conform/commit/7d19c82835585768a7601a1745a344a347f0ee9c) chore(ci): show git status when dirty working tree (#74)
* [`4a6cc1c`](https://github.com/talos-systems/conform/commit/4a6cc1cb9529ad561372c09fd7f579694c212b1b) feat: adding command line flag for commit msg (#73)
* [`088e0a7`](https://github.com/talos-systems/conform/commit/088e0a76b6e04ca521c6b9e03e66acb03b8dfbe8) chore(*): output binaries (#72)
* [`4194aa5`](https://github.com/talos-systems/conform/commit/4194aa5004222589c36265d4649452b72fb744bd) chore(*): format .conform.yaml (#71)
* [`05cfacb`](https://github.com/talos-systems/conform/commit/05cfacb6edd0dd92bd28bbb60c91970a1dc46d3a) feat(*): omit symbol and DWARF symbol tables (#70)
* [`0f0ff02`](https://github.com/talos-systems/conform/commit/0f0ff02f6b134d2249f90c3698937b62222bfebd) refactor(policy): start policy error report at 1 (#68)
* [`4aaf049`](https://github.com/talos-systems/conform/commit/4aaf04924064dfcaaba3d984666f1fd4f4dbe4fe) chore(*): go module cleanup (#67)
* [`fbc195c`](https://github.com/talos-systems/conform/commit/fbc195c8055a61bc2f45cf95f29957aea2e429b3) chore(*): disable cgo (#66)
* [`155b036`](https://github.com/talos-systems/conform/commit/155b0369ddf5bea3b162067b804a487fb3b634e2) feat(metadata): mark SHA as dirty (#63)
* [`f39b434`](https://github.com/talos-systems/conform/commit/f39b4343fd0181dbb1ffca2c36573e6724768764) chore(*): always push latest (#65)
* [`1276371`](https://github.com/talos-systems/conform/commit/12763710cbe120102f5da965524a84b2b4d57bf1) chore(*): go modules and faster linting (#64)
* [`0a0cba3`](https://github.com/talos-systems/conform/commit/0a0cba34137d39002bf762b52d19121290be9980) feat(*): add build command (#62)
* [`aed2c22`](https://github.com/talos-systems/conform/commit/aed2c223188bb7ee04046976a4e1249e1451d8ca) fix(policy): update regex to allow optional scope (#61)
* [`1933d19`](https://github.com/talos-systems/conform/commit/1933d192d81bfb67150be634905179bf4fe183e1) fix(pipeline): nil pointer when no defined pipeline (#60)
* [`7fd1e89`](https://github.com/talos-systems/conform/commit/7fd1e89c567df89bbc16449182ec10a9ce2e5c18) feat(fmt): add fmt command (#59)
* [`42c5d09`](https://github.com/talos-systems/conform/commit/42c5d09d8189e222a5785232f98b000bc869228f) feat(docker): expose the image name and tag separately (#58)
* [`5a73ea6`](https://github.com/talos-systems/conform/commit/5a73ea6ed7490f91995e9fb916213226ac50fb01) feat(git): recursively search for .git in parent directories (#56)
* [`1e57715`](https://github.com/talos-systems/conform/commit/1e577157c87a76461c8576a7f6b696ab3704c53f) fix(cli): invalid variable message (#53)
* [`72061b1`](https://github.com/talos-systems/conform/commit/72061b11a121bd6ba1ad6e25a0e66059a09fc5d8) feat(metadata): allow users to specify variables (#52)
* [`4b7b903`](https://github.com/talos-systems/conform/commit/4b7b903aac3671f839115f6240d4e12e5256e742) fix(policy): strip leading newline from commit message (#50)
* [`779bf93`](https://github.com/talos-systems/conform/commit/779bf930fb4cdbe2d94b39c2218f9b2854efaf7a) fix(pipeline): don't show stdout of artifact extraction (#49)
* [`0373fea`](https://github.com/talos-systems/conform/commit/0373fea42e663c91e512fe1c144721b299bae457) feat(*): services, skip flag, and UX improvements (#43)
* [`c53f523`](https://github.com/talos-systems/conform/commit/c53f52307152a6392835e23797f1432799a1e0cc) feat(renderer): allow templates to be retrieved from URL (#41)
* [`c2cb181`](https://github.com/talos-systems/conform/commit/c2cb1818f44f6b5d0dd73b2e0de64c70972fd72e) refactor(*): minor changes (#37)
* [`19e4656`](https://github.com/talos-systems/conform/commit/19e4656da53bc707af536efa49ceae4edc82de8c) refactor(*): add enforcer package (#33)
* [`26ff570`](https://github.com/talos-systems/conform/commit/26ff570fbe76d82ac5e571eb709efa9e9459c99c) chore(ci): use autonomy/golang:1.8.3 for build and test tasks (#32)
* [`44d4ef0`](https://github.com/talos-systems/conform/commit/44d4ef0256dd48965eabcbeade2838bb9ccdeaea) fix(policy): check the entire commit header length (#31)
* [`bd3404a`](https://github.com/talos-systems/conform/commit/bd3404a593d9e891a2c70aadf25add80e7a71ae5) refactor(*): rename conform directory to pkg (#30)
* [`daa39f3`](https://github.com/talos-systems/conform/commit/daa39f36a0376e1e4767da5e43d13f0ec1b2c838) docs(readme): update example (#28)
* [`9383d3e`](https://github.com/talos-systems/conform/commit/9383d3ebab5b14c186fc84299c76cfa4d404a9c6) feat(policy): enforce 72 character limit on commit header (#29)
* [`6a115cf`](https://github.com/talos-systems/conform/commit/6a115cfbe5a750b8a4aecea0467adf258fa25e74) chore(ci): check if not a PR (#27)
* [`ec732ae`](https://github.com/talos-systems/conform/commit/ec732ae306598fd292c3b0e8ce38359d0da71aae) chore(ci): build on master or tag (#26)
* [`ba44e03`](https://github.com/talos-systems/conform/commit/ba44e03d35ce19ce0372970609110c46aa6be7ed) chore(ci): skip branch check in script (#25)
* [`6cece44`](https://github.com/talos-systems/conform/commit/6cece4467af854f47b5066a226ac43e4f8c60ecb) chore(ci): use Travis deploy (#24)
* [`31f79af`](https://github.com/talos-systems/conform/commit/31f79af8dc091d9bcb7ef9f2e3b6647ab12ef553) chore(ci): update script to print useful information (#23)
* [`a18332a`](https://github.com/talos-systems/conform/commit/a18332af48f1e5ddb5308ac9ed3d1ff030e96a1b) fix(metadata): nil version struct (#22)
* [`3ae8e5f`](https://github.com/talos-systems/conform/commit/3ae8e5f0f333d1bc44e31fb256448f36f4619332) refactor(*): make conform.yaml a dotfile (#21)
* [`936b64e`](https://github.com/talos-systems/conform/commit/936b64e1b6169a28ab254455df3457366d04a5fa) refactor(*): complete rewrite (#20)
* [`03caad0`](https://github.com/talos-systems/conform/commit/03caad0cb1a02e8bf688557f21c3bd58b2c69fdc) feat(policy): add policy enforcement; enforce git commit policy (#17)
* [`9927a05`](https://github.com/talos-systems/conform/commit/9927a05dc80f0ce8eb7a4cafaf49efc1db28cc57) refactor(docker): read Dockerfile from stdin (#16)
* [`dcc9fe5`](https://github.com/talos-systems/conform/commit/dcc9fe5417f820fe4c2b1673b41f7022d536c80d) chore(ci): use the stable Docker repository (#15)
* [`d37461a`](https://github.com/talos-systems/conform/commit/d37461ae38f87d0701dfe492ef497cef321b26f9) refactor(git): use go-git instead of shelled out commands (#14)
* [`71fa116`](https://github.com/talos-systems/conform/commit/71fa116a49757e7357467c7e21bac29ceab86beb) Add pre-release to git info (#12)
* [`fd5c627`](https://github.com/talos-systems/conform/commit/fd5c62790011ef68bf96aa7720e88839c61c1c11) Use gometalinter and fix linting errors (#11)
* [`0e66ba1`](https://github.com/talos-systems/conform/commit/0e66ba1b6fa9e60296fb0276d18c13d1bad88d10) Fix deploy on tags (#10)
* [`00fbfa8`](https://github.com/talos-systems/conform/commit/00fbfa845376925c51ead7c723ce4f5401da9df8) Use generic language in Travis build (#9)
* [`925dabf`](https://github.com/talos-systems/conform/commit/925dabfde80372ee952da4d4584f1e5fbf60c087) Remove 'version' from path of variables set at build time (#8)
* [`aa8ced7`](https://github.com/talos-systems/conform/commit/aa8ced7d725629190bba8e47f29bb24bb1bae3ad) Fix package path of variables set at build time (#7)
* [`ad8eef9`](https://github.com/talos-systems/conform/commit/ad8eef9901a2a1800751361453112ba5766cde6b) Fix copy of artifact in image template (#6)
* [`6311568`](https://github.com/talos-systems/conform/commit/6311568530d96447d9f6b1de673a777532fe6fd6) Set execution bit of deploy script (#5)
* [`92643d5`](https://github.com/talos-systems/conform/commit/92643d5fd3e0a46a4e937b6ec947b26f06f043d1) Stream script output and deploy on master or tags (#4)
* [`81055ed`](https://github.com/talos-systems/conform/commit/81055ed01af331bf8d5c5596cf8dfe26f608e532) Return script output on error (#3)
* [`426abe1`](https://github.com/talos-systems/conform/commit/426abe1aea53e396d5a2da260d71f75372344cc1) Fix bad tag detection and setup CI (#2)
* [`0c55035`](https://github.com/talos-systems/conform/commit/0c55035ee0f2129a4cb63d07da834dc2210684f1) Initial implementation (#1)
* [`994ba0b`](https://github.com/talos-systems/conform/commit/994ba0b98618d07e29e21092f6929ac9399e6fd5) Initial commit
</p>
</details>

### Changes since v0.1.0-alpha.22
<details><summary>2 commits</summary>
<p>

* [`6caaabc`](https://github.com/talos-systems/conform/commit/6caaabc0f9721d161e68593fe3ee776761bdc893) feat: allow comments before the license header
* [`9a470a1`](https://github.com/talos-systems/conform/commit/9a470a16fe05643de83db40bf3f8cb0c3452ee2f) docs: update README with the latest changes
</p>
</details>

### Dependency Changes

This release has no dependency changes

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
