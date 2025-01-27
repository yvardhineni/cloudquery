package client

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func getTestConnection() string {
	testConn := os.Getenv("CQ_DEST_PG_TEST_CONN")
	if testConn == "" {
		return "postgresql://postgres:pass@localhost:5432/postgres?sslmode=disable"
	}
	return testConn
}

var safeMigrations = plugin.SafeMigrations{
	AddColumn:           true,
	AddColumnNotNull:    false,
	RemoveColumn:        true,
	RemoveColumnNotNull: false,
	ChangeColumn:        false,
}

func TestPgPlugin(t *testing.T) {
	ctx := context.Background()
	p := plugin.NewPlugin("postgresql", "development", New)
	s := &Spec{
		ConnectionString: getTestConnection(),
		PgxLogLevel:      LogLevelTrace,
	}
	b, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	err = p.Init(ctx, b, plugin.NewClientOptions{})
	if err != nil {
		t.Fatal(err)
	}
	testOpts := schema.TestSourceOptions{
		SkipMaps:      true,
		TimePrecision: time.Microsecond, // only us precision supported by time cols
	}
	plugin.TestWriterSuiteRunner(t,
		p,
		plugin.WriterTestSuiteTests{
			SkipDeleteRecord: true,
			SafeMigrations:   safeMigrations,
		},
		plugin.WithTestDataOptions(testOpts),
	)
}
