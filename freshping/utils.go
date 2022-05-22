package freshping

import (
	"context"
	"errors"
	"os"

	freshping "github.com/francois2metz/steampipe-plugin-freshping/freshping/client"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func connect(ctx context.Context, d *plugin.QueryData) (*freshping.Client, error) {
	// get freshping client from cache
	cacheKey := "freshping"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*freshping.Client), nil
	}

	freshpingConfig := GetConfig(d.Connection)

	if &freshpingConfig == nil {
		return nil, errors.New("You must have a freshping config file")
	}

	key := os.Getenv("FRESHPING_KEY")
	subdomain := os.Getenv("FRESHPING_SUBDOMAIN")

	if freshpingConfig.Key != nil {
		key = *freshpingConfig.Key
	}
	if freshpingConfig.Subdomain != nil {
		subdomain = *freshpingConfig.Subdomain
	}

	if key == "" {
		return nil, errors.New("'key' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}
	if subdomain == "" {
		return nil, errors.New("'subdomain' must be set in the connection configuration. Edit your connection configuration file and then restart Steampipe")
	}

	c := freshping.New(
		freshping.WithAuth(key, subdomain),
	)

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, c)

	return c, nil
}
