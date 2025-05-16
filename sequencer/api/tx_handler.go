package api

import (
	"net/http"
	"strconv"
	"strings"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/tokamak-network/syb-sequencer/sequencer/common"
)

type TxResponse struct {
	ItemID      int64  `json:"item_id"`
	BatchNum    int64  `json:"batch_num"`
	Position    int    `json:"position"`
	Type        string `json:"type"`
	FromIdx     *int64 `json:"from_idx,omitempty"`
	FromEthAddr string `json:"from_eth_addr,omitempty"`
	ToIdx       int64  `json:"to_idx"`
	ToEthAddr   string `json:"to_eth_addr,omitempty"`
	Amount      string `json:"amount"`
}

type PaginatedTxResponse struct {
	Transactions []TxResponse `json:"transactions"`
	Pagination   Pagination   `json:"pagination"`
}

type Pagination struct {
	CurrentPage  int   `json:"currentPage"`
	ItemsPerPage int   `json:"itemsPerPage"`
	TotalItems   int64 `json:"totalItems"`
	TotalPages   int   `json:"totalPages"`
}

func convertTxToResponse(tx *common.Tx) TxResponse {
	resp := TxResponse{
		ItemID:   tx.ItemID,
		BatchNum: tx.BatchNum,
		Position: tx.Position,
		Type:     tx.Type,
		FromIdx:  tx.FromIdx,
		ToIdx:    tx.ToIdx,
	}

	if tx.Amount != nil {
		resp.Amount = tx.Amount.String()
	} else {
		resp.Amount = "0"
	}

	if len(tx.FromEthAddr) > 0 {
		resp.FromEthAddr = ethCommon.BytesToAddress(tx.FromEthAddr).Hex()
	}

	if len(tx.ToEthAddr) > 0 {
		resp.ToEthAddr = ethCommon.BytesToAddress(tx.ToEthAddr).Hex()
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

	c.JSON(http.StatusOK, gin.H{"transactions": txResponses})
}

func (a *API) GetTransactionsByAccount(c *gin.Context) {
	accountAddress := c.Param("accountAddress")

	if !ethCommon.IsHexAddress(accountAddress) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account address format"})
		return
	}

	txs, err := a.db.GetTxsByAccountAddress(accountAddress) // Assumed DB method
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transactions for account: " + err.Error()})
		return
	}

	if len(txs) == 0 {
		c.JSON(http.StatusOK, gin.H{"transactions": []TxResponse{}}) // Return empty list if no txs found
		return
	}

	// Convert to response format
	accTxResponses := make([]TxResponse, len(txs))
	for i, tx := range txs {
		accTxResponses[i] = convertTxToResponse(tx)
	}

	c.JSON(http.StatusOK, gin.H{"transactions": accTxResponses})
}

func (a *API) GetTransactionsPaginated(c *gin.Context) {
	defaultPage := 1
	defaultLimit := 10
	defaultSortBy := "item_id"
	defaultSortOrder := "DESC"

	// Parse query parameters
	pageStr := c.DefaultQuery("page", strconv.Itoa(defaultPage))
	limitStr := c.DefaultQuery("limit", strconv.Itoa(defaultLimit))
	sortParam := c.DefaultQuery("sort", defaultSortBy+"_"+defaultSortOrder) // e.g., "item_id_desc"

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
	sortParts := strings.Split(sortParam, "_")
	if len(sortParts) == 2 {
		potentialSortBy := sortParts[0]
		potentialSortOrder := strings.ToUpper(sortParts[1])

		allowedSortBy := map[string]bool{"item_id": true, "batch_num": true, "type": true} // Add other sortable fields from your common.Tx
		if allowedSortBy[potentialSortBy] {
			sortBy = potentialSortBy
		}

		if potentialSortOrder == "ASC" || potentialSortOrder == "DESC" {
			sortOrder = potentialSortOrder
		}
	}

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
		totalPages = (int(totalItems) + limit - 1) / limit // Ceiling division
	}

	c.JSON(http.StatusOK, PaginatedTxResponse{
		Transactions: txResponses,
		Pagination: Pagination{
			CurrentPage:  page,
			ItemsPerPage: limit,
			TotalItems:   totalItems,
			TotalPages:   totalPages,
		},
	})
}
