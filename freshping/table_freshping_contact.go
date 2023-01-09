package freshping

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableFreshpingContact() *plugin.Table {
	return &plugin.Table{
		Name:        "freshping_contact",
		Description: "Contacts who can receive alerts.",
		List: &plugin.ListConfig{
			Hydrate: listContact,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_INT,
				Description: "Unique ID of the contact.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "Name of the contact.",
			},
			{
				Name:        "email",
				Type:        proto.ColumnType_STRING,
				Description: "Email of the contact.",
			},
		},
	}
}

func listContact(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("freshping_contact.listContact", "connection_error", err)
		return nil, err
	}
	res, err := client.GetUsers()
	if err != nil {
		plugin.Logger(ctx).Error("freshping_user.listContact", err)
		return nil, err
	}
	for _, contact := range res.Contacts {
		d.StreamListItem(ctx, contact)
	}
	return nil, nil
}
