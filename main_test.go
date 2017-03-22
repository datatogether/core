package archive

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/gchaincl/dotsql"
	"os"
	"testing"
)

var appDB *sql.DB

func TestMain(m *testing.M) {
	flag.Parse()

	teardown := setupTestDatabase()

	retCode := m.Run()
	teardown()
	os.Exit(retCode)
}

func setupTestDatabase() func() {
	var err error
	appDB, err = SetupConnection(os.Getenv("POSTGRES_URL"))
	if err != nil {
		appDB.Close()
		panic(err.Error())
	}

	teardown, err := initializeAppSchema(appDB)
	if err != nil {
		panic(err.Error())
	}

	if err := resetTestData(appDB, "primers", "subprimers", "urls", "links", "metadata", "snapshots", "archive_requests"); err != nil {
		panic(err.Error())
	}

	return teardown
}

// WARNING - THIS ZAPS WHATEVER DB IT'S GIVEN. DO NOT CALL THIS SHIT.
// used for testing only, returns a teardown func
func initializeAppSchema(db *sql.DB) (func(), error) {
	schema, err := dotsql.LoadFromFile("sql/schema.sql")
	if err != nil {
		return nil, err
	}

	for _, cmd := range []string{
		"drop-all",
		"create-primers",
		"create-subprimers",
		"create-urls",
		"create-links",
		"create-metadata",
		"create-snapshots",
		"create-archive_requests",
	} {
		if _, err := schema.Exec(db, cmd); err != nil {
			fmt.Println(cmd, "error:", err)
			return nil, err
		}
	}

	teardown := func() {
		if _, err := schema.Exec(db, "drop-all"); err != nil {
			panic(err.Error())
		}
	}

	return teardown, nil
}

// drops test data tables & re-inserts base data from sql/test_data.sql, based on
// passed in table names
func resetTestData(db *sql.DB, tables ...string) error {
	schema, err := dotsql.LoadFromFile("sql/test_data.sql")
	if err != nil {
		return err
	}
	for _, t := range tables {
		if _, err := schema.Exec(db, fmt.Sprintf("delete-%s", t)); err != nil {
			return err
		}
		if _, err := schema.Exec(db, fmt.Sprintf("insert-%s", t)); err != nil {
			return err
		}
	}
	return nil
}
