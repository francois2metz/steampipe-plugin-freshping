package freshping

import (
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/schema"
)

type freshpingConfig struct {
	Key       *string `cty:"key"`
	Subdomain *string `cty:"subdomain"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"key": {
		Type: schema.TypeString,
	},
	"subdomain": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &freshpingConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) freshpingConfig {
	if connection == nil || connection.Config == nil {
		return freshpingConfig{}
	}
	config, _ := connection.Config.(freshpingConfig)
	return config
}
