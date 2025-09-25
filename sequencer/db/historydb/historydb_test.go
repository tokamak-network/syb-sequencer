package historydb

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/semaphore"
)

func TestNewHistoryDB_TableDriven(t *testing.T) {
	tests := []struct {
		name       string
		setupDB    func() (*sql.DB, *sql.DB)
		apiConnCon *APIConnectionController
		testCase   int
	}{
		{
			name: "Successfully intialized db",
			setupDB: func() (*sql.DB, *sql.DB) {
				dbRead, _, err := sqlmock.New()
				require.NoError(t, err)

				dbWrite, _, err := sqlmock.New()
				require.NoError(t, err)

				return dbRead, dbWrite
			},
			apiConnCon: &APIConnectionController{
				smphr:   semaphore.NewWeighted(1),
				timeout: 1 * time.Second,
			},
			testCase: 1,
		},
		{
			name: "Error intializing db - nil values passed",
			setupDB: func() (*sql.DB, *sql.DB) {
				return nil, nil
			},
			apiConnCon: nil,
			testCase:   2,
		},
		{
			name: "Error intializing db - error in api connection",
			setupDB: func() (*sql.DB, *sql.DB) {
				dbRead, _, err := sqlmock.New()
				require.NoError(t, err)
				dbRead.Close()

				dbWrite, _, err := sqlmock.New()
				require.NoError(t, err)
				dbWrite.Close()

				return dbRead, dbWrite
			},
			apiConnCon: nil,
			testCase:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbRead, dbWrite := tt.setupDB()

			defer func() {
				if dbRead != nil {
					_ = dbRead.Close()
				}
				if dbWrite != nil {
					_ = dbWrite.Close()
				}
			}()

			historyDB := NewHistoryDB(dbRead, dbWrite, tt.apiConnCon)

			switch tt.testCase {
			case 1:
				assert.NotNil(t, historyDB)
				assert.IsType(t, &HistoryDB{}, historyDB)
				assert.Equal(t, dbRead, historyDB.dbRead)
				assert.Equal(t, dbWrite, historyDB.dbWrite)
				assert.Equal(t, tt.apiConnCon, historyDB.apiConnCon)
			case 2:
				assert.Nil(t, historyDB.dbRead)
			case 3:
				assert.Nil(t, historyDB.apiConnCon)
				assert.NotNil(t, historyDB.dbRead)

			default:
				return
			}
		})
	}
}
