---
organization: francois2metz
category: ["saas"]
brand_color: "#1f9999"
display_name: "Freshping"
short_name: "freshping"
description: "Steampipe plugin for querying Freshping."
og_description: "Query Freshping with SQL! Open source CLI. No DB required."
icon_url: "/images/plugins/francois2metz/freshping.svg"
og_image: "/images/plugins/francois2metz/freshping-social-graphic.png"
---

# Freshping + Steampipe

[Freshping](https://www.freshworks.com/website-monitoring/) is a monitoring service that is part of Freshworks.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select
  id,
  name,
  url
from
  freshping_check
```

```
+----+-----------------+------------------------------+
| id | name            | url                          |
+----+-----------------+------------------------------+
| 42 | Ca reste ouvert | https://caresteouvert.fr     |
| 43 | API             | https://api.caresteouvert.fr |
+----+-----------------+------------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/francois2metz/freshping/tables)**

## Get started

### Install

Download and install the latest Freshping plugin:

```bash
steampipe plugin install francois2metz/freshping
```

### Configuration

Installing the latest freshping plugin will create a config file (`~/.steampipe/config/freshping.spc`) with a single connection named `freshping`:

```hcl
connection "freshping" {
    plugin = "francois2metz/freshping"

    # Freshping API Key
    # To get it:
    # 1. Login to your account
    # 2. Go to settings > Account settings
    # 3. Copy API Key
    # See https://support.freshping.io/en/support/solutions/articles/50000003709-freshping-api-documentation#Authentication
    key = "NoalNogitnazud1GranfairkyefAymvu"

    # Your freshping subdomain
    subdomain = ""
}
```

You can also use environment variables:

- `FRESHPING_KEY` for the API KEY
- `FRESHPING_SUBDOMAIN` for the subdomain

## Get Involved

* Open source: https://github.com/francois2metz/steampipe-plugin-freshping
