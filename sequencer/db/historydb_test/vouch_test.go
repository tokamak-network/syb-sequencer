package historydb_test

import (
	"database/sql"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tokamak-network/syb-sequencer/sequencer/common"
	"github.com/tokamak-network/syb-sequencer/sequencer/db/historydb"
)

func calculateVouchIdx(fromIdx, toIdx common.AccountIdx) common.VouchIdx {
	vouchTableKeyStr := fmt.Sprintf("%d%d", fromIdx, toIdx)
	vouchTableKeyBigInt, ok := new(big.Int).SetString(vouchTableKeyStr, 10)
	if !ok {
		panic(fmt.Sprintf("failed to create vouch table key from string '%s'", vouchTableKeyStr))
	}
	return common.VouchIdx(vouchTableKeyBigInt.Uint64())
}

func TestCalculateVouchIdx(t *testing.T) {
	tests := []struct {
		name     string
		fromIdx  common.AccountIdx
		toIdx    common.AccountIdx
		expected common.VouchIdx
	}{
		{
			name:     "basic case",
			fromIdx:  common.AccountIdx(10),
			toIdx:    common.AccountIdx(20),
			expected: common.VouchIdx(1020),
		},
		{
			name:     "single digits",
			fromIdx:  common.AccountIdx(1),
			toIdx:    common.AccountIdx(2),
			expected: common.VouchIdx(12),
		},
		{
			name:     "larger numbers",
			fromIdx:  common.AccountIdx(123),
			toIdx:    common.AccountIdx(456),
			expected: common.VouchIdx(123456),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculateVouchIdx(tt.fromIdx, tt.toIdx)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func createMockVouch() *common.Vouch {
	fromIdx := common.AccountIdx(10)
	toIdx := common.AccountIdx(20)
	return &common.Vouch{
		Idx:         calculateVouchIdx(fromIdx, toIdx), // FromIdx(10) + ToIdx(20) = VouchIdx(1020)
		FromIdx:     fromIdx,
		FromEthAddr: ethCommon.HexToAddress("0x1234567890123456789012345678901234567890"),
		ToIdx:       toIdx,
		ToEthAddr:   ethCommon.HexToAddress("0x0987654321098765432109876543210987654321"),
	}
}

func TestHistoryDB_AddVouch(t *testing.T) {
	tests := []struct {
		name        string
		vouch       *common.Vouch
		mockSetup   func(sqlmock.Sqlmock)
		expectError bool
		errorMsg    string
	}{
		{
			name:  "successful save",
			vouch: createMockVouch(),
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`INSERT INTO vouch`).
					WithArgs(
						common.VouchIdx(1020),
						common.AccountIdx(10),
						ethCommon.HexToAddress("0x1234567890123456789012345678901234567890").Bytes(),
						common.AccountIdx(20),
						ethCommon.HexToAddress("0x0987654321098765432109876543210987654321").Bytes(),
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectError: false,
		},
		{
			name:  "database connection error",
			vouch: createMockVouch(),
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`INSERT INTO vouch`).
					WithArgs(
						common.VouchIdx(1020),
						common.AccountIdx(10),
						ethCommon.HexToAddress("0x1234567890123456789012345678901234567890").Bytes(),
						common.AccountIdx(20),
						ethCommon.HexToAddress("0x0987654321098765432109876543210987654321").Bytes(),
					).
					WillReturnError(sql.ErrConnDone)
			},
			expectError: true,
			errorMsg:    "AddVouch: failed to add vouch",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			apiConnCon := historydb.NewAPIConnectionController(1, 1*time.Second)

			historyDB := historydb.NewHistoryDB(db, db, apiConnCon)
			tt.mockSetup(mock)

			err = historyDB.AddVouch(tt.vouch)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorMsg != "" {
					assert.Contains(t, err.Error(), tt.errorMsg)
				}
			} else {
				assert.NoError(t, err)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestHistoryDB_GetVouchByIdx(t *testing.T) {
	mockAddr1 := ethCommon.HexToAddress("0x1234567890123456789012345678901234567890")
	mockAddr2 := ethCommon.HexToAddress("0x0987654321098765432109876543210987654321")

	tests := []struct {
		name        string
		vouchIdx    common.VouchIdx
		mockSetup   func(sqlmock.Sqlmock)
		expected    *common.Vouch
		expectError bool
		errorMsg    string
	}{
		{
			name:     "successful retrieval",
			vouchIdx: calculateVouchIdx(common.AccountIdx(10), common.AccountIdx(20)), // VouchIdx(1020)
			mockSetup: func(mock sqlmock.Sqlmock) {
				vouchIdx := calculateVouchIdx(common.AccountIdx(10), common.AccountIdx(20))
				rows := sqlmock.NewRows([]string{
					"from_idx", "from_eth_addr", "to_idx", "to_eth_addr",
				}).
					AddRow(common.AccountIdx(10), mockAddr1.Bytes(), common.AccountIdx(20), mockAddr2.Bytes())

				mock.ExpectQuery(`SELECT from_idx, from_eth_addr, to_idx, to_eth_addr FROM vouch WHERE idx = \$1`).
					WithArgs(vouchIdx).
					WillReturnRows(rows)
			},
			expected: &common.Vouch{
				Idx:         calculateVouchIdx(common.AccountIdx(10), common.AccountIdx(20)), // VouchIdx(1020) - set by implementation
				FromIdx:     common.AccountIdx(10),
				FromEthAddr: mockAddr1,
				ToIdx:       common.AccountIdx(20),
				ToEthAddr:   mockAddr2,
			},
			expectError: false,
		},
		{
			name:     "vouch not found",
			vouchIdx: common.VouchIdx(999),
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT from_idx, from_eth_addr, to_idx, to_eth_addr FROM vouch WHERE idx = \$1`).
					WithArgs(common.VouchIdx(999)).
					WillReturnError(sql.ErrNoRows)
			},
			expected:    nil,
			expectError: false, // Function returns nil, nil for not found
		},
		{
			name:     "database query error",
			vouchIdx: common.VouchIdx(1),
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT from_idx, from_eth_addr, to_idx, to_eth_addr FROM vouch WHERE idx = \$1`).
					WithArgs(common.VouchIdx(1)).
					WillReturnError(sql.ErrConnDone)
			},
			expectError: true,
			errorMsg:    "GetVouchByIdx: failed to get vouch",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			apiConnCon := historydb.NewAPIConnectionController(1, 1*time.Second)
			historyDB := historydb.NewHistoryDB(db, db, apiConnCon)
			tt.mockSetup(mock)

			result, err := historyDB.GetVouchByIdx(tt.vouchIdx)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorMsg != "" {
					assert.Contains(t, err.Error(), tt.errorMsg)
				}
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				if tt.expected == nil {
					assert.Nil(t, result)
				} else {
					assert.NotNil(t, result)
					assert.Equal(t, tt.expected.Idx, result.Idx)
					assert.Equal(t, tt.expected.FromIdx, result.FromIdx)
					assert.Equal(t, tt.expected.ToIdx, result.ToIdx)
					assert.Equal(t, tt.expected.FromEthAddr, result.FromEthAddr)
					assert.Equal(t, tt.expected.ToEthAddr, result.ToEthAddr)
				}
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestHistoryDB_GetVouchesByEthAddress(t *testing.T) {
	mockAddr1 := ethCommon.HexToAddress("0x1234567890123456789012345678901234567890")
	mockAddr2 := ethCommon.HexToAddress("0x0987654321098765432109876543210987654321")

	tests := []struct {
		name        string
		address     ethCommon.Address
		mockSetup   func(sqlmock.Sqlmock)
		expectedLen int
		expectError bool
		errorMsg    string
	}{
		{
			name:    "successful retrieval with results",
			address: mockAddr1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{
					"idx", "from_idx", "from_eth_addr", "to_idx", "to_eth_addr",
				}).
					AddRow(calculateVouchIdx(common.AccountIdx(10), common.AccountIdx(20)), common.AccountIdx(10), mockAddr1.Bytes(), common.AccountIdx(20), mockAddr2.Bytes()).
					AddRow(calculateVouchIdx(common.AccountIdx(10), common.AccountIdx(25)), common.AccountIdx(10), mockAddr1.Bytes(), common.AccountIdx(25), mockAddr2.Bytes())
				mock.ExpectQuery(`SELECT idx, from_idx, from_eth_addr, to_idx, to_eth_addr FROM vouch WHERE from_eth_addr = \$1`).
					WithArgs(mockAddr1.Bytes()).
					WillReturnRows(rows)
			},
			expectedLen: 2,
			expectError: false,
		},
		{
			name:    "no vouches found",
			address: mockAddr1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{
					"idx", "from_idx", "from_eth_addr", "to_idx", "to_eth_addr",
				})
				mock.ExpectQuery(`SELECT idx, from_idx, from_eth_addr, to_idx, to_eth_addr FROM vouch WHERE from_eth_addr = \$1`).
					WithArgs(mockAddr1.Bytes()).
					WillReturnRows(rows)
			},
			expectedLen: 0,
			expectError: false,
		},
		{
			name:    "database query error",
			address: mockAddr1,
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT idx, from_idx, from_eth_addr, to_idx, to_eth_addr FROM vouch WHERE from_eth_addr = \$1`).
					WithArgs(mockAddr1.Bytes()).
					WillReturnError(sql.ErrConnDone)
			},
			expectError: true,
			errorMsg:    "GetVouchesByEthAddress: failed to query vouches",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			apiConnCon := historydb.NewAPIConnectionController(1, 1*time.Second)

			historyDB := historydb.NewHistoryDB(db, db, apiConnCon)
			tt.mockSetup(mock)

			result, err := historyDB.GetVouchesByEthAddress(tt.address)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorMsg != "" {
					assert.Contains(t, err.Error(), tt.errorMsg)
				}
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)

				if tt.expectedLen > 0 {
					assert.NotNil(t, result)
					assert.Len(t, result, tt.expectedLen)
				} else {
					assert.Empty(t, result)
				}

				// Verify first result if any
				if len(result) > 0 {
					assert.NotNil(t, result)
					expectedIdx := calculateVouchIdx(common.AccountIdx(10), common.AccountIdx(20)) // VouchIdx(1020)
					assert.Equal(t, expectedIdx, result[0].Idx)
					assert.Equal(t, common.AccountIdx(10), result[0].FromIdx)
					assert.Equal(t, common.AccountIdx(20), result[0].ToIdx)
					assert.Equal(t, mockAddr1, result[0].FromEthAddr)
					assert.Equal(t, mockAddr2, result[0].ToEthAddr)
				}
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestHistoryDB_GetVouchersByIdx(t *testing.T) {
	tests := []struct {
		name        string
		toIdx       common.AccountIdx
		mockSetup   func(sqlmock.Sqlmock)
		expected    []common.AccountIdx
		expectError bool
		errorMsg    string
	}{
		{
			name:  "successful retrieval with multiple vouchers",
			toIdx: common.AccountIdx(20),
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"from_idx"}).
					AddRow(common.AccountIdx(10)).
					AddRow(common.AccountIdx(15))
				mock.ExpectQuery(`SELECT DISTINCT from_idx FROM vouch WHERE to_idx = \$1 ORDER BY from_idx`).
					WithArgs(common.AccountIdx(20)).
					WillReturnRows(rows)
			},
			expected:    []common.AccountIdx{10, 15},
			expectError: false,
		},
		{
			name:  "no vouchers found",
			toIdx: common.AccountIdx(99),
			mockSetup: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"from_idx"})
				mock.ExpectQuery(`SELECT DISTINCT from_idx FROM vouch WHERE to_idx = \$1 ORDER BY from_idx`).
					WithArgs(common.AccountIdx(99)).
					WillReturnRows(rows)
			},
			expected:    nil,
			expectError: false,
		},
		{
			name:  "database query error",
			toIdx: common.AccountIdx(20),
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT DISTINCT from_idx FROM vouch WHERE to_idx = \$1 ORDER BY from_idx`).
					WithArgs(common.AccountIdx(20)).
					WillReturnError(sql.ErrConnDone)
			},
			expectError: true,
			errorMsg:    "GetVouchersByToIdx: failed to query vouchers",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			apiConnCon := historydb.NewAPIConnectionController(1, 1*time.Second)

			historyDB := historydb.NewHistoryDB(db, db, apiConnCon)
			tt.mockSetup(mock)

			result, err := historyDB.GetVouchersByIdx(tt.toIdx)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorMsg != "" {
					assert.Contains(t, err.Error(), tt.errorMsg)
				}
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestHistoryDB_DeleteVouchByIdx(t *testing.T) {
	tests := []struct {
		name        string
		vouchIdx    common.VouchIdx
		mockSetup   func(sqlmock.Sqlmock)
		expectError bool
		errorMsg    string
	}{
		{
			name:     "successful delete",
			vouchIdx: common.VouchIdx(1),
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM vouch WHERE idx = \$1`).
					WithArgs(common.VouchIdx(1)).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			expectError: false,
		},
		{
			name:     "no rows deleted - vouch not found",
			vouchIdx: common.VouchIdx(999),
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM vouch WHERE idx = \$1`).
					WithArgs(common.VouchIdx(999)).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			expectError: true,
			errorMsg:    "no vouch found with idx 999 to delete",
		},
		{
			name:     "database execution error",
			vouchIdx: common.VouchIdx(1),
			mockSetup: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(`DELETE FROM vouch WHERE idx = \$1`).
					WithArgs(common.VouchIdx(1)).
					WillReturnError(sql.ErrConnDone)
			},
			expectError: true,
			errorMsg:    "DeleteVouchByIdx: failed to delete vouch",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			apiConnCon := historydb.NewAPIConnectionController(1, 1*time.Second)

			historyDB := historydb.NewHistoryDB(db, db, apiConnCon)
			tt.mockSetup(mock)

			err = historyDB.DeleteVouchByIdx(tt.vouchIdx)

			if tt.expectError {
				assert.Error(t, err)
				if tt.errorMsg != "" {
					assert.Contains(t, err.Error(), tt.errorMsg)
				}
			} else {
				assert.NoError(t, err)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
