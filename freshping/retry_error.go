package freshping

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func shouldRetryError() plugin.ErrorPredicateWithContext {
	return func(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData, err error) bool {
		if strings.Contains(err.Error(), "429") {
			plugin.Logger(ctx).Debug("freshping.shouldRetryError", "rate_limit_error", err)
			return true
		}
		return false
	}
}
