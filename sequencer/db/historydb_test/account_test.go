package historydb_test

import (
	"database/sql"
	"math/big"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tokamak-network/syb-sequencer/sequencer/common"
	"github.com/tokamak-network/syb-sequencer/sequencer/db/historydb"
)

// ---- Tests ----

func TestAddAccount(t *testing.T) {
	tests := []struct {
		name      string
		account   *common.Account
		setupMock func(mock sqlmock.Sqlmock)
		wantErr   bool
		errMsg    string
	}{
		{
			name: "successful account addition",
			account: &common.Account{
				Idx:           3,
				EthAddr:       ethCommon.HexToAddress("0x3333333333333333333333333333333333333333"),
				Balance:       big.NewInt(3000),
				Score:         big.NewInt(30),
				ScoreSiblings: []*big.Int{big.NewInt(7), big.NewInt(8)},
			},
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO account").
					WithArgs(
						common.AccountIdx(3),
						ethCommon.HexToAddress("0x3333333333333333333333333333333333333333").Bytes(),
						"3000",
						"30",
						pq.Array([]string{"7", "8"}),
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name: "account with nil score siblings",
			account: &common.Account{
				Idx:           4,
				EthAddr:       ethCommon.HexToAddress("0x4444444444444444444444444444444444444444"),
				Balance:       big.NewInt(4000),
				Score:         big.NewInt(40),
				ScoreSiblings: []*big.Int{nil, nil},
			},
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO account").
					WithArgs(
						common.AccountIdx(4),
						ethCommon.HexToAddress("0x4444444444444444444444444444444444444444").Bytes(),
						"4000",
						"40",
						pq.Array([]string{"0", "0"}),
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name: "database error",
			account: &common.Account{
				Idx:           5,
				EthAddr:       ethCommon.HexToAddress("0x5555555555555555555555555555555555555555"),
				Balance:       big.NewInt(5000),
				Score:         big.NewInt(50),
				ScoreSiblings: []*big.Int{},
			},
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("INSERT INTO account").
					WithArgs(
						common.AccountIdx(5),
						ethCommon.HexToAddress("0x5555555555555555555555555555555555555555").Bytes(),
						"5000",
						"50",
						pq.Array([]string{}),
					).
					WillReturnError(sql.ErrConnDone)
			},
			wantErr: true,
			errMsg:  "failed to add account",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			hdb := historydb.NewHistoryDB(db, db, nil)
			tt.setupMock(mock)

			// Execute
			err = hdb.AddAccount(tt.account)

			// Assert
			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
			} else {
				assert.NoError(t, err)
			}

			// Verify all expectations were met
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetAccountByIdx(t *testing.T) {
	tests := []struct {
		name      string
		idx       uint32
		setupMock func(mock sqlmock.Sqlmock)
		want      *common.Account
		wantErr   bool
	}{
		{
			name: "successful retrieval",
			idx:  1,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT eth_addr, balance, score, score_siblings FROM account WHERE idx = \\$1").
					WithArgs(uint32(1)).
					WillReturnRows(
						sqlmock.NewRows([]string{"eth_addr", "balance", "score", "score_siblings"}).
							AddRow(
								ethCommon.HexToAddress("0x1111111111111111111111111111111111111111").Bytes(),
								"1000",
								"10",
								pq.Array([]string{"1", "2"}),
							),
					)
			},
			want: &common.Account{
				Idx:           1,
				EthAddr:       ethCommon.HexToAddress("0x1111111111111111111111111111111111111111"),
				Balance:       big.NewInt(1000),
				Score:         big.NewInt(10),
				ScoreSiblings: []*big.Int{big.NewInt(1), big.NewInt(2)},
			},
			wantErr: false,
		},
		{
			name: "account not found",
			idx:  999,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT eth_addr, balance, score, score_siblings FROM account WHERE idx = \\$1").
					WithArgs(uint32(999)).
					WillReturnError(sql.ErrNoRows)
			},
			want:    nil,
			wantErr: false, // GetAccountByIdx returns nil, nil for not found
		},
		{
			name: "invalid balance format",
			idx:  2,
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT eth_addr, balance, score, score_siblings FROM account WHERE idx = \\$1").
					WithArgs(uint32(2)).
					WillReturnRows(
						sqlmock.NewRows([]string{"eth_addr", "balance", "score", "score_siblings"}).
							AddRow(
								ethCommon.HexToAddress("0x2222222222222222222222222222222222222222").Bytes(),
								"invalid_balance",
								"20",
								pq.Array([]string{}),
							),
					)
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			hdb := historydb.NewHistoryDB(db, db, nil)
			tt.setupMock(mock)

			// Execute
			got, err := hdb.GetAccountByIdx(tt.idx)

			// Assert
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}

			// Verify all expectations were met
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetAllAccounts(t *testing.T) {
	tests := []struct {
		name      string
		setupMock func(mock sqlmock.Sqlmock)
		want      []*common.Account
		wantTotal int64
		wantErr   bool
	}{
		{
			name: "successful retrieval of multiple accounts",
			setupMock: func(mock sqlmock.Sqlmock) {
				// First: expect the SELECT query (this runs first in your implementation)
				mock.ExpectQuery("SELECT idx, eth_addr, balance, score, score_siblings FROM account ORDER BY idx").
					WillReturnRows(
						sqlmock.NewRows([]string{"idx", "eth_addr", "balance", "score", "score_siblings"}).
							AddRow(
								common.AccountIdx(1),
								ethCommon.HexToAddress("0x1111111111111111111111111111111111111111").Bytes(),
								"1000",
								"10",
								pq.Array([]string{"1", "2"}),
							).
							AddRow(
								common.AccountIdx(2),
								ethCommon.HexToAddress("0x2222222222222222222222222222222222222222").Bytes(),
								"2000",
								"20",
								pq.Array([]string{"4", "5"}),
							),
					)

				// Second: expect the COUNT query (this runs after in your implementation)
				mock.ExpectQuery("SELECT COUNT\\(\\*\\) FROM account").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))
			},
			want: []*common.Account{
				{
					Idx:           1,
					EthAddr:       ethCommon.HexToAddress("0x1111111111111111111111111111111111111111"),
					Balance:       big.NewInt(1000),
					Score:         big.NewInt(10),
					ScoreSiblings: []*big.Int{big.NewInt(1), big.NewInt(2)},
				},
				{
					Idx:           2,
					EthAddr:       ethCommon.HexToAddress("0x2222222222222222222222222222222222222222"),
					Balance:       big.NewInt(2000),
					Score:         big.NewInt(20),
					ScoreSiblings: []*big.Int{big.NewInt(4), big.NewInt(5)},
				},
			},
			wantTotal: 2,
			wantErr:   false,
		},
		{
			name: "select query error",
			setupMock: func(mock sqlmock.Sqlmock) {
				// SELECT query fails
				mock.ExpectQuery("SELECT idx, eth_addr, balance, score, score_siblings FROM account ORDER BY idx").
					WillReturnError(sql.ErrConnDone)
			},
			want:      nil,
			wantTotal: 0,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			hdb := historydb.NewHistoryDB(db, db, nil)
			tt.setupMock(mock)

			// Execute
			got, gotTotal, err := hdb.GetAllAccounts()

			// Assert
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wantTotal, gotTotal)
				assert.Equal(t, tt.want, got)
			}

			// Verify all expectations were met
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

// TestUpdateAccountBalance tests the UpdateAccountBalance functionality
func TestUpdateAccountBalance(t *testing.T) {
	tests := []struct {
		name       string
		idx        common.AccountIdx
		newBalance *big.Int
		setupMock  func(mock sqlmock.Sqlmock)
		wantErr    bool
		errMsg     string
	}{
		{
			name:       "successful balance update",
			idx:        1,
			newBalance: big.NewInt(5000),
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("UPDATE account SET balance = \\$1 WHERE idx = \\$2").
					WithArgs("5000", common.AccountIdx(1)).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			wantErr: false,
		},
		{
			name:       "account not found",
			idx:        999,
			newBalance: big.NewInt(5000),
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("UPDATE account SET balance = \\$1 WHERE idx = \\$2").
					WithArgs("5000", common.AccountIdx(999)).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			wantErr: true,
			errMsg:  "no account found",
		},
		{
			name:       "database error",
			idx:        1,
			newBalance: big.NewInt(5000),
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("UPDATE account SET balance = \\$1 WHERE idx = \\$2").
					WithArgs("5000", common.AccountIdx(1)).
					WillReturnError(sql.ErrConnDone)
			},
			wantErr: true,
			errMsg:  "failed to update balance",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			hdb := historydb.NewHistoryDB(db, db, nil)
			tt.setupMock(mock)

			// Execute
			err = hdb.UpdateAccountBalance(tt.idx, tt.newBalance)

			// Assert
			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
			} else {
				assert.NoError(t, err)
			}

			// Verify all expectations were met
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestGetAccountByEthAddress(t *testing.T) {
	tests := []struct {
		name      string
		ethAddr   ethCommon.Address
		setupMock func(mock sqlmock.Sqlmock)
		want      *common.Account
		wantErr   bool
		errMsg    string
	}{
		{
			name:    "successful retrieval by eth address",
			ethAddr: ethCommon.HexToAddress("0x1111111111111111111111111111111111111111"),
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT idx, balance, score, score_siblings FROM account WHERE eth_addr = \\$1").
					WithArgs(ethCommon.HexToAddress("0x1111111111111111111111111111111111111111").Bytes()).
					WillReturnRows(
						sqlmock.NewRows([]string{"idx", "balance", "score", "score_siblings"}).
							AddRow(
								common.AccountIdx(1),
								"1000",
								"10",
								pq.Array([]string{"1", "2"}),
							),
					)
			},
			want: &common.Account{
				Idx:           1,
				EthAddr:       ethCommon.HexToAddress("0x1111111111111111111111111111111111111111"),
				Balance:       big.NewInt(1000),
				Score:         big.NewInt(10),
				ScoreSiblings: []*big.Int{big.NewInt(1), big.NewInt(2)},
			},
			wantErr: false,
		},
		{
			name:    "account not found by eth address",
			ethAddr: ethCommon.HexToAddress("0x9999999999999999999999999999999999999999"),
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT idx, balance, score, score_siblings FROM account WHERE eth_addr = \\$1").
					WithArgs(ethCommon.HexToAddress("0x9999999999999999999999999999999999999999").Bytes()).
					WillReturnError(sql.ErrNoRows)
			},
			want:    nil,
			wantErr: false, // GetAccountByEthAddress returns nil, nil for not found
		},
		{
			name:    "invalid balance format",
			ethAddr: ethCommon.HexToAddress("0x2222222222222222222222222222222222222222"),
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT idx, balance, score, score_siblings FROM account WHERE eth_addr = \\$1").
					WithArgs(ethCommon.HexToAddress("0x2222222222222222222222222222222222222222").Bytes()).
					WillReturnRows(
						sqlmock.NewRows([]string{"idx", "balance", "score", "score_siblings"}).
							AddRow(
								common.AccountIdx(2),
								"invalid_balance",
								"20",
								pq.Array([]string{"4", "5"}),
							),
					)
			},
			want:    nil,
			wantErr: true,
			errMsg:  "failed to parse balance string",
		},
		{
			name:    "invalid score format",
			ethAddr: ethCommon.HexToAddress("0x3333333333333333333333333333333333333333"),
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT idx, balance, score, score_siblings FROM account WHERE eth_addr = \\$1").
					WithArgs(ethCommon.HexToAddress("0x3333333333333333333333333333333333333333").Bytes()).
					WillReturnRows(
						sqlmock.NewRows([]string{"idx", "balance", "score", "score_siblings"}).
							AddRow(
								common.AccountIdx(3),
								"3000",
								"invalid_score",
								pq.Array([]string{"6", "7"}),
							),
					)
			},
			want:    nil,
			wantErr: true,
			errMsg:  "failed to parse score string",
		},
		{
			name:    "account with nil score siblings",
			ethAddr: ethCommon.HexToAddress("0x4444444444444444444444444444444444444444"),
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT idx, balance, score, score_siblings FROM account WHERE eth_addr = \\$1").
					WithArgs(ethCommon.HexToAddress("0x4444444444444444444444444444444444444444").Bytes()).
					WillReturnRows(
						sqlmock.NewRows([]string{"idx", "balance", "score", "score_siblings"}).
							AddRow(
								common.AccountIdx(4),
								"4000",
								"40",
								pq.Array([]string{"invalid_sibling", "0"}),
							),
					)
			},
			want: &common.Account{
				Idx:           4,
				EthAddr:       ethCommon.HexToAddress("0x4444444444444444444444444444444444444444"),
				Balance:       big.NewInt(4000),
				Score:         big.NewInt(40),
				ScoreSiblings: []*big.Int{big.NewInt(0), big.NewInt(0)}, // invalid_sibling becomes 0
			},
			wantErr: false,
		},
		{
			name:    "database connection error",
			ethAddr: ethCommon.HexToAddress("0x5555555555555555555555555555555555555555"),
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT idx, balance, score, score_siblings FROM account WHERE eth_addr = \\$1").
					WithArgs(ethCommon.HexToAddress("0x5555555555555555555555555555555555555555").Bytes()).
					WillReturnError(sql.ErrConnDone)
			},
			want:    nil,
			wantErr: true,
			errMsg:  "failed to get account with eth_addr",
		},
		{
			name:    "empty score siblings array",
			ethAddr: ethCommon.HexToAddress("0x6666666666666666666666666666666666666666"),
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT idx, balance, score, score_siblings FROM account WHERE eth_addr = \\$1").
					WithArgs(ethCommon.HexToAddress("0x6666666666666666666666666666666666666666").Bytes()).
					WillReturnRows(
						sqlmock.NewRows([]string{"idx", "balance", "score", "score_siblings"}).
							AddRow(
								common.AccountIdx(6),
								"6000",
								"60",
								pq.Array([]string{}), // empty array
							),
					)
			},
			want: &common.Account{
				Idx:           6,
				EthAddr:       ethCommon.HexToAddress("0x6666666666666666666666666666666666666666"),
				Balance:       big.NewInt(6000),
				Score:         big.NewInt(60),
				ScoreSiblings: []*big.Int{}, // empty slice
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			hdb := historydb.NewHistoryDB(db, db, nil)
			tt.setupMock(mock)

			// Execute
			got, err := hdb.GetAccountByEthAddress(tt.ethAddr)

			// Assert
			if tt.wantErr {
				assert.Error(t, err)
				if tt.errMsg != "" {
					assert.Contains(t, err.Error(), tt.errMsg)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}

			// Verify all expectations were met
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
