package api

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/tokamak-network/syb-sequencer/sequencer/common"
)

type TxResponse struct {
	ItemID      int64             `json:"item_id"`
	BatchNum    uint32            `json:"batch_num"`
	Position    string            `json:"position"`
	Type        string            `json:"type"`
	FromIdx     common.AccountIdx `json:"from_idx,omitempty"`
	FromEthAddr string            `json:"from_eth_addr,omitempty"`
	ToIdx       common.AccountIdx `json:"to_idx"`
	ToEthAddr   string            `json:"to_eth_addr,omitempty"`
	Amount      string            `json:"amount"`
	BlockNumber uint64            `json:"block_number"`
	Timestamp   uint64            `json:"timestamp"`
	GasFee      string            `json:"gas_fee"`
	TxHash      string            `json:"tx_hash"`
	IsTxForged  bool              `json:"is_tx_forged"`
}
type PaginatedTxResponse struct {
	Message      string       `json:"message"`
	Transactions []TxResponse `json:"transactions"`
	Pagination   Pagination   `json:"pagination"`
}

type Pagination struct {
	CurrentPage  int   `json:"currentPage"`
	ItemsPerPage int   `json:"itemsPerPage"`
	TotalItems   int64 `json:"totalItems"`
	TotalPages   int   `json:"totalPages"`
}
type AccountResponse struct {
	Idx      string              `json:"idx"`
	EthAddr  string              `json:"eth_addr"`
	Balance  string              `json:"balance"`
	Score    string              `json:"score"`
	Vouchers []common.AccountIdx `json:"vouchers"`
	ScoreInt string              `json:"score_int,omitempty"`
}

type PaginatedAccountResponse struct {
	Message       string            `json:"message"`
	Accounts      []AccountResponse `json:"accounts"`
	TotalAccounts int64             `json:"total_accounts"`
}

type ScoreMerkleProofResponse struct {
	Idx       string   `json:"idx"`
	ScoreRoot string   `json:"score_root"`
	Score     string   `json:"score"`
	Siblings  [][]byte `json:"siblings"`
}

const (
	GetAllTransactionsResponseMessage       = "Retrieved all transactions"
	GetTransactionsByAccountResponseMessage = "Retrieved transactions for account %s"
	GetTransactionsPaginatedResponseMessage = "Retrieved transactions in paginated format"
	NoTransactionsFoundResponseMessage      = "No transactions found for account %s"
	GetTransactionByHashResponseMessage     = "Retrieved transaction by hash %s"

	// Add these to your existing constants
	GetAllAccountsResponseMessage  = "Retrieved all accounts"
	GetAccountByIdxResponseMessage = "Retrieved account with index %s"
	AccountNotFoundResponseMessage = "Account with index %s not found"
)

func (a *API) convertTxToResponse(tx *common.Tx) TxResponse {
	resp := TxResponse{
		ItemID:      tx.ItemID,
		BatchNum:    tx.BatchNum,
		Position:    tx.Position.String(),
		Type:        tx.Type,
		FromIdx:     tx.FromIdx,
		ToIdx:       tx.ToIdx,
		BlockNumber: tx.BlockNumber,
		Timestamp:   tx.Timestamp,
	}

	ctx := context.Background()
	callOpts := &bind.CallOpts{
		Context: ctx,
	}

	if tx.Amount != nil {
		resp.Amount = tx.Amount.String()
	} else {
		resp.Amount = "0"
	}

	if tx.GasFee != nil {
		resp.GasFee = tx.GasFee.String()
	} else {
		resp.GasFee = "0"
	}

	if len(tx.FromEthAddr) > 0 {
		resp.FromEthAddr = ethCommon.Bytes2Hex(tx.FromEthAddr)
	}

	if len(tx.ToEthAddr) > 0 {
		resp.ToEthAddr = ethCommon.Bytes2Hex(tx.ToEthAddr)
	}

	if len(tx.TxHash) > 0 {
		resp.TxHash = ethCommon.Bytes2Hex(tx.TxHash)
	}

	resp.IsTxForged = false
	lastForgedTx, err := a.sybilContract.LastForgedTxn(callOpts)
	if err != nil {
		fmt.Printf("Error fetching last forged transaction: %v\n", err)
	} else if lastForgedTx != nil {
		isPositionForged := tx.Position.Cmp(lastForgedTx) <= 0
		lastTxExists := lastForgedTx.Cmp(big.NewInt(0)) > 0

		if lastTxExists && isPositionForged {
			resp.IsTxForged = true
		}
	}
	return resp
}

func (a *API) GetAllTransactions(c *gin.Context) {
	txs, err := a.db.GetAllTxs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transactions: " + err.Error()})
		return
	}

	txResponses := make([]TxResponse, len(txs))
	for i, tx := range txs {
		txResponses[i] = a.convertTxToResponse(tx)
	}

	c.JSON(http.StatusOK, gin.H{"transactions": txResponses, "message": GetAllTransactionsResponseMessage})
}

func (a *API) GetTransactionsByAccount(c *gin.Context) {
	accountAddress := c.Param("accountAddress")

	if !ethCommon.IsHexAddress(accountAddress) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account address format"})
		return
	}

	txs, err := a.db.GetTxsByAccountAddress(accountAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transactions for account: " + err.Error()})
		return
	}

	if len(txs) == 0 {
		c.JSON(http.StatusOK, gin.H{"transactions": []TxResponse{}, "message": fmt.Sprintf(NoTransactionsFoundResponseMessage, accountAddress)})
		return
	}

	// Convert to response format
	accTxResponses := make([]TxResponse, len(txs))
	for i, tx := range txs {
		accTxResponses[i] = a.convertTxToResponse(tx)
	}

	c.JSON(http.StatusOK, gin.H{"transactions": accTxResponses, "message": fmt.Sprintf(GetTransactionsByAccountResponseMessage, accountAddress)})
}

func (a *API) GetTransactionsPaginated(c *gin.Context) {
	defaultPage := 1
	defaultLimit := 10
	defaultSortBy := "tx_timestamp"
	defaultSortOrder := "DESC"

	pageStr := c.DefaultQuery("page", strconv.Itoa(defaultPage))
	limitStr := c.DefaultQuery("limit", strconv.Itoa(defaultLimit))

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = defaultPage
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = defaultLimit
	}

	sortBy := defaultSortBy
	sortOrder := defaultSortOrder

	offset := (page - 1) * limit

	txs, totalItems, err := a.db.GetTxsPaginated(limit, offset, sortBy, sortOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transactions: " + err.Error()})
		return
	}

	txResponses := make([]TxResponse, len(txs))
	for i, tx := range txs {
		txResponses[i] = a.convertTxToResponse(tx)
	}

	totalPages := 0
	if totalItems > 0 {
		totalPages = (int(totalItems) + limit - 1) / limit
	}

	c.JSON(http.StatusOK, PaginatedTxResponse{
		Message:      GetTransactionsPaginatedResponseMessage,
		Transactions: txResponses,
		Pagination: Pagination{
			CurrentPage:  page,
			ItemsPerPage: limit,
			TotalItems:   totalItems,
			TotalPages:   totalPages,
		},
	})
}

func (a *API) GetTransactionByHash(c *gin.Context) {
	txHash := c.Param("txHash")
	txHashBytes := ethCommon.HexToHash(txHash).Bytes()

	tx, err := a.db.GetTxByHash(txHashBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transaction: " + err.Error()})
		return
	}

	if tx == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	txResponse := a.convertTxToResponse(tx)
	c.JSON(http.StatusOK, gin.H{
		"transaction": txResponse,
		"message":     fmt.Sprintf(GetTransactionByHashResponseMessage, txHash),
	})
}

func (a *API) convertAccountToResponse(account *common.Account, vouchers []common.AccountIdx) AccountResponse {
	resp := AccountResponse{
		Idx:     account.Idx.String(),
		EthAddr: account.EthAddr.Hex(),
	}

	if account.Balance != nil {
		resp.Balance = account.Balance.String()
	} else {
		resp.Balance = "0"
	}

	accountScore, err := a.statedb.GetScore(common.ScoreIdx(account.Idx))
	if err != nil {
		account.Score = big.NewInt(0)
	} else {
		account.Score = accountScore.Score
	}

	if account.Score != nil {
		resp.Score = account.Score.String()
		if account.Score.IsInt64() {
			scoreInt := account.Score.Int64()
			resp.ScoreInt = strconv.FormatInt(scoreInt, 10)
		}
	} else {
		resp.Score = "0"
	}

	resp.Vouchers = vouchers

	return resp
}

func (a *API) GetAllAccounts(c *gin.Context) {
	accounts, totalItems, err := a.db.GetAllAccounts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve accounts: " + err.Error()})
		return
	}

	accountResponses := make([]AccountResponse, len(accounts))
	for i, account := range accounts {
		vouchers, err := a.db.GetVouchersByIdx(account.Idx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve vouchers: " + err.Error()})
			return
		}
		accountResponses[i] = a.convertAccountToResponse(account, vouchers)
	}

	c.JSON(http.StatusOK, PaginatedAccountResponse{
		Message:       GetAllAccountsResponseMessage,
		Accounts:      accountResponses,
		TotalAccounts: totalItems,
	})
}

func (a *API) GetAccountByIdx(c *gin.Context) {
	idxStr := c.Param("idx")

	idx, err := strconv.ParseUint(idxStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account index format"})
		return
	}

	accountIdx := common.AccountIdx(idx)

	account, err := a.db.GetAccountByIdx(uint32(accountIdx))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve account: " + err.Error()})
		return
	}

	if account == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf(AccountNotFoundResponseMessage, idxStr),
		})
		return
	}

	vouchers, err := a.db.GetVouchersByIdx(account.Idx)

	accountResponse := a.convertAccountToResponse(account, vouchers)

	c.JSON(http.StatusOK, gin.H{
		"account": accountResponse,
		"message": fmt.Sprintf(GetAccountByIdxResponseMessage, idxStr),
	})
}

func (a *API) GetScoreMerkleProof(c *gin.Context) {
	idxStr := c.Param("idx")

	idx, err := strconv.ParseUint(idxStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account index format"})
		return
	}

	scoreIdx := common.ScoreIdx(idx)
	score, err := a.statedb.GetScore(scoreIdx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve score: " + err.Error()})
		return
	}

	// Get ScoreRoot
	scoreRoot := a.statedb.GetSTRoot()

	siblings, err := a.statedb.GetSiblings(scoreIdx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve sibling: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, ScoreMerkleProofResponse{
		Idx:       scoreIdx.String(),
		ScoreRoot: scoreRoot.String(),
		Score:     score.Score.String(),
		Siblings:  siblings,
	})
}
