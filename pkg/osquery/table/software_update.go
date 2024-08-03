package table

import (
	"context"

	"github.com/osquery/osquery-go"
	"github.com/osquery/osquery-go/plugin/table"
)

func CreateSoftwareUpdateTable(server *osquery.ExtensionManagerServer) {
	server.RegisterPlugin(table.NewPlugin("software_update", columns(), generate))
}

func columns() []table.ColumnDefinition {
	return []table.ColumnDefinition{
		table.TextColumn("version"),
		table.TextColumn("last_updated_at"),
	}
}

func generate(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	return []map[string]string{
		{
			"version":         "1.0.0",
			"last_updated_at": "2022-08-09",
		},
	}, nil
}
