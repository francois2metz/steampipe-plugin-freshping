package freshping

import (
	"context"
	"errors"
	"os"

	freshping "github.com/francois2metz/steampipe-plugin-freshping/freshping/client"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*freshping.Client, error) {
	// get freshping client from cache
	cacheKey := "freshping"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*freshping.Client), nil
	}

	freshpingConfig := GetConfig(d.Connection)

	api_key := os.Getenv("FRESHPING_API_KEY")
	subdomain := os.Getenv("FRESHPING_SUBDOMAIN")

	if freshpingConfig.APIKey != nil {
		api_key = *freshpingConfig.APIKey
	}
	if freshpingConfig.Subdomain != nil {
		subdomain = *freshpingConfig.Subdomain
	}

	if api_key == "" {
		return nil, errors.New("'api_key' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}
	if subdomain == "" {
		return nil, errors.New("'subdomain' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	c := freshping.New(
		freshping.WithAuth(api_key, subdomain),
	)

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, c)

	return c, nil
}
