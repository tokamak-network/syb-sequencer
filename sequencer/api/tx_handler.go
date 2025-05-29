package api

import (
	"fmt"
	"net/http"
	"strconv"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/tokamak-network/syb-sequencer/sequencer/common"
)

type TxResponse struct {
	ItemID      int64             `json:"item_id"`
	BatchNum    int64             `json:"batch_num"`
	Position    int               `json:"position"`
	Type        string            `json:"type"`
	FromIdx     common.AccountIdx `json:"from_idx,omitempty"`
	FromEthAddr string            `json:"from_eth_addr,omitempty"`
	ToIdx       common.AccountIdx `json:"to_idx"`
	ToEthAddr   string            `json:"to_eth_addr,omitempty"`
	Amount      string            `json:"amount"`
	BlockNumber uint64            `json:"block_number"`
	Timestamp   uint64            `json:"timestamp"`
	GasFee      string            `json:"gas_fee"`
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

const (
	GetAllTransactionsResponseMessage       = "Retrieved all transactions"
	GetTransactionsByAccountResponseMessage = "Retrieved transactions for account %s"
	GetTransactionsPaginatedResponseMessage = "Retrieved transactions in paginated format"
	NoTransactionsFoundResponseMessage      = "No transactions found for account %s"
)

func convertTxToResponse(tx *common.Tx) TxResponse {
	resp := TxResponse{
		ItemID:      tx.ItemID,
		BatchNum:    tx.BatchNum,
		Position:    tx.Position,
		Type:        tx.Type,
		FromIdx:     tx.FromIdx,
		ToIdx:       tx.ToIdx,
		BlockNumber: tx.BlockNumber,
		Timestamp:   tx.Timestamp,
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
		txResponses[i] = convertTxToResponse(tx)
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
		accTxResponses[i] = convertTxToResponse(tx)
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
		txResponses[i] = convertTxToResponse(tx)
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
