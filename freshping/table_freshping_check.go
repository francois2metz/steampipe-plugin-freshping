package freshping

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
)

func tableFreshpingCheck() *plugin.Table {
	return &plugin.Table{
		Name:        "freshping_check",
		Description: "A check refers to the specifications configured to monitor an end-point.",
		List: &plugin.ListConfig{
			Hydrate: listCheck,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getCheck,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Description: "Unique ID of the check.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the check.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "Status of the check (AV: 'Available', NR: 'Not Responding', RE: 'Reporting Error', PS: 'Monitoring Paused', DME: 'DNS Record Matching Failure', SME: 'String Check Matching Failure', SCE: 'Status Code check Failure').",
			},
			{
				Name:        "performance_status",
				Type:        proto.ColumnType_STRING,
				Description: "Performance status (GP: 'Good Performance', DP: 'Degraded Performance', AV: 'Available', PS: 'Monitoring Paused').",
			},
			{
				Name:        "alert_note",
				Type:        proto.ColumnType_STRING,
				Description: "Troubleshooting instructions.",
			},
			{
				Name:        "location",
				Type:        proto.ColumnType_STRING,
				Description: "The monitoring location.",
			},
			{
				Name:        "monitoring_interval",
				Type:        proto.ColumnType_INT,
				Description: "The check frequency (in seconds).",
			},
			{
				Name:        "url",
				Type:        proto.ColumnType_STRING,
				Description: "The URL to check.",
			},
			{
				Name:        "request_timeout",
				Type:        proto.ColumnType_INT,
				Description: "The wait time for a response from the server (in seconds).",
			},
			{
				Name:        "basic_auth_username",
				Type:        proto.ColumnType_STRING,
				Description: "Basic auth username.",
			},
			{
				Name:        "basic_auth_password",
				Type:        proto.ColumnType_STRING,
				Description: "Basic auth password.",
			},
			{
				Name:        "command_string",
				Type:        proto.ColumnType_STRING,
				Description: "The string to send to the tcp/udp socket.",
			},
			{
				Name:        "success_string",
				Type:        proto.ColumnType_STRING,
				Description: "The string to look for success at the response for tcp/udp check.",
			},
			{
				Name:        "error_string",
				Type:        proto.ColumnType_STRING,
				Description: "The string to look for eror at the response for tcp/udp check.",
			},
		},
	}
}

func listCheck(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("freshping_check.listCheck", "connection_error", err)
		return nil, err
	}
	checks, err := client.GetChecks()
	if err != nil {
		plugin.Logger(ctx).Error("freshping_check.listCheck", err)
		return nil, err
	}
	for _, check := range checks {
		d.StreamListItem(ctx, check)
	}
	return nil, nil
}

func getCheck(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("freshping_check.getCheck", "connection_error", err)
		return nil, err
	}
	id := d.KeyColumnQuals["id"].GetInt64Value()
	check, err := client.GetCheck(id)
	if err != nil {
		plugin.Logger(ctx).Error("freshping_check.getCheck", err)
		return nil, err
	}
	return check, nil
}
