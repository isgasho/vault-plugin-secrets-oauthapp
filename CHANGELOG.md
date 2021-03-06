# [1.7.0](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/compare/v1.6.0...v1.7.0) (2020-10-27)


### New

* Add minimum_seconds credential read option ([8050dc3ffdc0c774c37afe0f3038aa3863cfff66](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/commit/8050dc3ffdc0c774c37afe0f3038aa3863cfff66))
* Add tests for minimum_seconds ([52daf36215a3ea8e7dfa9f3ea5e884fc4020df4a](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/commit/52daf36215a3ea8e7dfa9f3ea5e884fc4020df4a))

### Update

* Make sure retrieved token meets all requirements ([f67616f15e955e5efc17580d3d23f9767fa2e1ad](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/commit/f67616f15e955e5efc17580d3d23f9767fa2e1ad))
* Rename tokenOk2Reuse to tokenValid ([ec8a7e45600fc79f7dc5e39de393cb81cb2ae253](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/commit/ec8a7e45600fc79f7dc5e39de393cb81cb2ae253))
* Treat zero time as never expired ([f93aff64d428a1281b7c367a0cb3eb567c28a4cd](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/commit/f93aff64d428a1281b7c367a0cb3eb567c28a4cd))

# [1.6.0](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/compare/v1.5.0...v1.6.0) (2020-09-28)


### Update

* expand characters allowed for creds path ([7264f66c69930a9c91c63a944673bc4bf9765ec9](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/commit/7264f66c69930a9c91c63a944673bc4bf9765ec9))
* use local Regex function instead of expanding framework function ([8d74a12dc388c9da1051c23c2ff55e0f2ec404a1](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/commit/8d74a12dc388c9da1051c23c2ff55e0f2ec404a1))

# [1.5.0](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/compare/v1.4.0...v1.5.0) (2020-09-28)


### Fix

* add keys to logger function calls ([0b2ea21a427c2c8b31d2ab394acdb3760112f59c](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/commit/0b2ea21a427c2c8b31d2ab394acdb3760112f59c))

### Update

* Give more complete oauth2 error messages ([31b0662c4b83afcde7508963a447790787f69d73](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/commit/31b0662c4b83afcde7508963a447790787f69d73))
* move oauth2 error messages to log ([48e1eccb92d0a413473a5f99fa7ea9a1d9ad1fb5](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/commit/48e1eccb92d0a413473a5f99fa7ea9a1d9ad1fb5))

# [1.4.0](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/compare/v1.3.0...v1.4.0) (2020-09-28)


### Chore

* Update CODEOWNERS ([92be6a42e595f54bc2d9b001f8a80d2325a61ff0](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/commit/92be6a42e595f54bc2d9b001f8a80d2325a61ff0))

### Fix

* Use OptionError when discovery_url fails ([83e100eaed65764ab13e0f3cd8d4844c30eec4b7](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/commit/83e100eaed65764ab13e0f3cd8d4844c30eec4b7))

### New

* Add discovery_url ([cb59513a2b4fc8c03833304dc62b1c7d877cae6b](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/commit/cb59513a2b4fc8c03833304dc62b1c7d877cae6b))

### Upgrade

* Bump lodash from 4.17.15 to 4.17.19 ([4c55418da05974170f67d274d01e4d003b1fc063](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/commit/4c55418da05974170f67d274d01e4d003b1fc063))
* Bump node-fetch from 2.6.0 to 2.6.1 ([e62e8942fb02204f9427fd9bf2a0112c7e003173](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/commit/e62e8942fb02204f9427fd9bf2a0112c7e003173))

# [1.3.0](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/compare/v1.2.0...v1.3.0) (2020-07-13)


### New

* add Google OAuth support ([194cfafeaff1ce246d80aa8f9bf783f909fb5179](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/commit/194cfafeaff1ce246d80aa8f9bf783f909fb5179))

### Upgrade

* Bump https-proxy-agent from 2.2.2 to 2.2.4 ([80ab6a27aa520f952632e91f6f26fa58b691e224](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/commit/80ab6a27aa520f952632e91f6f26fa58b691e224))
* Bump npm from 6.13.4 to 6.14.6 ([0a30d2a99804db00b6b580dbce34ad72bcf04d4c](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/commit/0a30d2a99804db00b6b580dbce34ad72bcf04d4c))

# [1.2.0](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/compare/v1.1.1...v1.2.0) (2020-04-06)


### New

* Add refresh_token option to creds write ([a74f7e484e20dacf21d9c9e742ea4d0f34ff7037](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/commit/a74f7e484e20dacf21d9c9e742ea4d0f34ff7037))

### Upgrade

* Bump handlebars from 4.2.0 to 4.5.3 ([188593baa9c83f4ec16f69d4f211b7c78c19def3](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/commit/188593baa9c83f4ec16f69d4f211b7c78c19def3))
* Bump npm from 6.11.3 to 6.13.4 ([e3aaa079b4fee4934f40f542fbff4ca554e544e4](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/commit/e3aaa079b4fee4934f40f542fbff4ca554e544e4))

## [1.1.1](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/compare/v1.1.0...v1.1.1) (2019-10-07)


### Fix

* Do not propagate API errors from providers when token refresh fails ([c3cf680a61a70775f9b2981da880fb335b7d2b3e](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/commit/c3cf680a61a70775f9b2981da880fb335b7d2b3e))

# [1.1.0](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/compare/v1.0.0...v1.1.0) (2019-09-19)


### Update

* Expose token type in credential response ([cf95ad17e69ccdeae97e1ddf07df26157fcce8a5](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/commit/cf95ad17e69ccdeae97e1ddf07df26157fcce8a5))

# 1.0.0 (2019-09-18)


### Build

* Remove unneeded image release ([ec02a426ffc5b900e48c07a91ffa0e1cb594a1f5](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/commit/ec02a426ffc5b900e48c07a91ffa0e1cb594a1f5))

### New

* Initial release ([1788c24609c036a6778b06883347c6fc5f5961a2](https://github.com/puppetlabs/vault-plugin-secrets-oauthapp/commit/1788c24609c036a6778b06883347c6fc5f5961a2))
