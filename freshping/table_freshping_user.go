package freshping

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableFreshpingUser() *plugin.Table {
	return &plugin.Table{
		Name:        "freshping_user",
		Description: "Users who have access to the freshping organization.",
		List: &plugin.ListConfig{
			Hydrate: listUser,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Description: "Unique ID of the user.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the user.",
			},
			{
				Name:        "email",
				Type:        proto.ColumnType_STRING,
				Description: "Email of the user.",
			},
			{
				Name:        "role",
				Type:        proto.ColumnType_STRING,
				Description: "Role of the user.",
			},
			{
				Name:        "disable_weekly_report_emails",
				Type:        proto.ColumnType_BOOL,
				Description: "Has the user disabled the weekly report email.",
			},
			{
				Name:        "disable_alert_emails",
				Type:        proto.ColumnType_BOOL,
				Description: "Has the user disabled alert emails.",
			},
		},
	}
}

func listUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("freshping_user.listUser", "connection_error", err)
		return nil, err
	}
	res, err := client.GetUsers()
	if err != nil {
		plugin.Logger(ctx).Error("freshping_user.listUser", err)
		return nil, err
	}
	for _, user := range res.Users {
		d.StreamListItem(ctx, user)
	}
	return nil, nil
}
