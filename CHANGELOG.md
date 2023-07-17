## v0.3.1 [2023-07-17]

_What's new?_

- Update SDK to 5.5.0
- Update imroc/req to 3.37.2

## v0.3.0 [2023-01-25]

_What's new?_

- **Breaking Changes**
  - The configuration has been updated from `key` to `api_key`
  - The environment variable `FRESHPING_KEY` has been renamed to `FRESHPING_API_KEY`
- The requests are retried in case of rate limiting
- Update SDK to 5.1.2
- Documentation updates

## v0.2.1 [2023-01-09]

_What's new?_

- The default configuration token and subdomain are now commented

## v0.2.0 [2023-01-09]

_What's new?_

- Add *freshping_contact* table
- Add *alert_users* and *alert_contacts* columns to the freshping_check table
- Update SDK to 5.0.2

## v0.1.0 [2022-09-01]

_What's new?_

- Update SDK to 4.1.5
- Update to go 1.19

## v0.0.1 [2022-06-02]

_What's new?_

- Initial release with tables:

  - freshping_check
  - freshping_user
