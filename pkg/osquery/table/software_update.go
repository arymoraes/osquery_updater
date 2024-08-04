package table

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"runtime"

	_ "github.com/mattn/go-sqlite3"

	"github.com/osquery/osquery-go"
	"github.com/osquery/osquery-go/plugin/table"
)

var softwareUpdateDataFiles = map[string]string{
	"windows": "AppData/Local/osquerything/data/SoftwareUpdate.sqlite",
	"darwin":  "Library/Application Support/osquerything/Data/SoftwareUpdate.sqlite",
	"linux":   "/opt/osquerything/SoftwareUpdate.sqlite",
}

func SoftwareUpdate(server *osquery.ExtensionManagerServer) {
	server.RegisterPlugin(table.NewPlugin("software_update", columns(), generate))
}

func columns() []table.ColumnDefinition {
	return []table.ColumnDefinition{
		table.TextColumn("version"),
		table.TextColumn("last_updated_at"),
	}
}

func generate(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	filePath := fromOSPath(runtime.GOOS)

	db, err := sql.Open("sqlite3", filePath)

	if err != nil {
		return nil, fmt.Errorf("connecting to sqlite db: %w", err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT version, last_updated_at FROM software_updates")
	if err != nil {
		return nil, fmt.Errorf("error querying rows from software_update: %w", err)
	}
	defer rows.Close()

	var results []map[string]string

	for rows.Next() {
		var version string
		var last_updated_at string
		err = rows.Scan(&version, &last_updated_at)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, map[string]string{
			"version":         version,
			"last_updated_at": last_updated_at,
		})
	}

	return results, nil
}

func fromOSPath(osName string) string {
	return softwareUpdateDataFiles[osName]
}
