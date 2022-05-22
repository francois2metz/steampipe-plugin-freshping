# Freshping plugin for Steampipe

Use SQL to query checks and users from [Freshping][].

- **[Get started →](docs/index.md)**
- Documentation: [Table definitions & examples](docs/tables)

## Quick start

Install the plugin with [Steampipe][]:

    steampipe plugin install francois2metz/freshping

## Development

To build the plugin and install it in your `.steampipe` directory

    make

Copy the default config file:

    cp config/freshping.spc ~/.steampipe/config/freshping.spc

## License

Apache 2

[steampipe]: https://steampipe.io
[freshping]: https://www.freshworks.com/website-monitoring/
