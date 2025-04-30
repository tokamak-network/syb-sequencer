package historydb

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// HistoryDB persist the historic of the rollup
type HistoryDB struct {
	dbRead     *sql.DB
	dbWrite    *sql.DB
	apiConnCon *APIConnectionController
}

// NewHistoryDB initializes the DB
func NewHistoryDB(dbRead, dbWrite *sql.DB, apiConnCon *APIConnectionController) *HistoryDB {
	return &HistoryDB{
		dbRead:     dbRead,
		dbWrite:    dbWrite,
		apiConnCon: apiConnCon,
	}
}
