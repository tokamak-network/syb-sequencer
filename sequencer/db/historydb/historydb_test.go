package historydb

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var historyDB *HistoryDB

func TestMain(m *testing.M) {
	// init DB
	db, err := InitSQLDB(5432, "localhost", "postgres", "123", "syb")
	if err != nil {
		panic(err)
	}
	apiConnCon := NewAPIConnectionController(1, time.Second)
	historyDB = NewHistoryDB(db, db, apiConnCon)
	// Run test
	result := m.Run()
	if err := db.Close(); err != nil {
		fmt.Println("Error closing the history DB", err)
	}
	os.Exit(result)
}

func TestConn(t *testing.T) {
	err := historyDB.dbRead.Ping()
	if err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}
	t.Logf("Successfully connected to the db")
}
