package test

import (
	"database/sql"
	"lib"
	"os"
	"testing"
)

var db *sql.DB

func TestMain(m *testing.M) {
	lib.LoadEnv()

	db = lib.DbOpen()
	defer db.Close()

	code := m.Run()

	truncateTable()

	defer os.Exit(code)
}

func truncateTable() {
	db.Query("DELETE FROM spa;")
}

func assertEqual(t *testing.T, actual string, expect string) {
	if actual != expect {
		t.Fatalf("%v != %v", string(actual), expect)
	}
}
