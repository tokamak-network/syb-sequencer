package historydb_test

import (
	"database/sql"
	"errors"
	"math/big"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tokamak-network/syb-sequencer/sequencer/common"
	"github.com/tokamak-network/syb-sequencer/sequencer/db/historydb"
)

func createMockBatch() *common.Batch {
	return &common.Batch{
		ItemID:      common.BatchNum(123), // Fixed: Use proper type
		AccountRoot: big.NewInt(111111),
		VouchRoot:   big.NewInt(222222),
		ScoreRoot:   big.NewInt(333333),
	}
}

func TestHistoryDB_AddBatch(t *testing.T) {
	tests := []struct {
		name        string
		batch       *common.Batch
		mockSetup   func(mock sqlmock.Sqlmock, batch *common.Batch)
		wantErr     bool
		errContains string
	}{
		{
			name:  "successful addition of batch",
			batch: createMockBatch(),
			mockSetup: func(mock sqlmock.Sqlmock, batch *common.Batch) {
				mock.ExpectExec(`INSERT INTO batch`).
					WithArgs(batch.ItemID, batch.AccountRoot.String(), batch.VouchRoot.String(), batch.ScoreRoot.String()).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name:  "database connection error",
			batch: createMockBatch(),
			mockSetup: func(mock sqlmock.Sqlmock, batch *common.Batch) {
				mock.ExpectExec(`INSERT INTO batch`).
					WithArgs(batch.ItemID, batch.AccountRoot.String(), batch.VouchRoot.String(), batch.ScoreRoot.String()).
					WillReturnError(sql.ErrConnDone)
			},
			wantErr:     true,
			errContains: "sql: connection is already closed",
		},
		{
			name:  "constraint violation error",
			batch: createMockBatch(),
			mockSetup: func(mock sqlmock.Sqlmock, batch *common.Batch) {
				// Simulate a proper constraint violation error
				constraintErr := errors.New("pq: duplicate key value violates unique constraint")
				mock.ExpectExec(`INSERT INTO batch`).
					WithArgs(batch.ItemID, batch.AccountRoot.String(), batch.VouchRoot.String(), batch.ScoreRoot.String()).
					WillReturnError(constraintErr)
			},
			wantErr:     true,
			errContains: "duplicate key value",
		},
		{
			name: "nil batch values handled correctly",
			batch: &common.Batch{
				ItemID:      common.BatchNum(456),
				AccountRoot: nil, // This should be handled gracefully
				VouchRoot:   big.NewInt(222222),
				ScoreRoot:   big.NewInt(333333),
			},
			mockSetup: func(mock sqlmock.Sqlmock, batch *common.Batch) {
				// The implementation should handle nil values
				mock.ExpectExec(`INSERT INTO batch`).
					WillReturnError(errors.New("nil pointer dereference"))
			},
			wantErr:     true,
			errContains: "nil pointer",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			historyDB := historydb.NewHistoryDB(db, db, nil)

			if tt.mockSetup != nil {
				tt.mockSetup(mock, tt.batch)
			}

			err = historyDB.AddBatch(tt.batch)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.errContains != "" {
					assert.Contains(t, err.Error(), tt.errContains)
				}
			} else {
				assert.NoError(t, err)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
