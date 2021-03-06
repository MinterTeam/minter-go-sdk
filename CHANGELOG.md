# Change Log

## [v2.3.0](https://github.com/MinterTeam/minter-go-sdk/tree/v2.3.0) (2021-05-20)

[Full Changelog](https://github.com/MinterTeam/minter-go-sdk/compare/v2.2.0...v2.3.0)

- DeepLink: Allow to set custom URL [#11](https://github.com/MinterTeam/minter-go-sdk/pull/11)

## [v2.2.0](https://github.com/MinterTeam/minter-go-sdk/tree/v2.2.0) (2021-04-12)

[Full Changelog](https://github.com/MinterTeam/minter-go-sdk/compare/v2.1.1...v2.2.0)

## [v2.1.1](https://github.com/MinterTeam/minter-go-sdk/tree/v2.1.1) (2020-12-24)

[Full Changelog](https://github.com/MinterTeam/minter-go-sdk/compare/v2.1.0...v2.1.1)

**Fixed:**

- Calculate commissions

## [v2.1.0](https://github.com/MinterTeam/minter-go-sdk/tree/v2.1.0) (2020-12-23)

[Full Changelog](https://github.com/MinterTeam/minter-go-sdk/compare/v2.0.3...v2.1.0)

**Added:**

- Autotests to Github Actions.
- More documentation and examples for API clients.
- Telegram channel with release notifications and discussions.
- Increase your commission multiplier to 10e16
- **WithDebug** method for *http_client*.
- **WithLogger** method for *http_client*.
- **WithHeaders** method for *http_client*.
- **AddressExtended** method with `delegated` flag support in request for *http_client* and *grpc_client*
- **AddressesExtended** method with `delegated` flag support in request for *http_client* and *grpc_client*
- **CandidateExtended** method with `not_show_stakes` flag support in request for *http_client* and *grpc_client*
- **CandidatesExtended** method with `not_show_stakes` flag support in request for *http_client* and *grpc_client*
- **CheckVersion** method for *http_client* and *grpc_client*

**Changed:**

- Remove `status` parameter from **Candidates** method for *http_client* and *grpc_client*, flag `not_show_stakes` sets
  as default.
