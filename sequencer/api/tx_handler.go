package api

import (
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tokamak-network/syb-sequencer/sequencer/db/historydb"
)

// TxResponse represents the response format for transaction data
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

// convertTxToResponse converts a database Tx to a response format
func convertTxToResponse(tx *historydb.Tx) TxResponse {
	resp := TxResponse{
		ItemID:   tx.ItemID,
		BatchNum: tx.BatchNum,
		Position: tx.Position,
		Type:     tx.Type,
		FromIdx:  tx.FromIdx,
		ToIdx:    tx.ToIdx,
	}

	// Convert amount to string
	if tx.Amount != nil {
		resp.Amount = tx.Amount.String()
	} else {
		resp.Amount = "0"
	}

	// Convert Ethereum addresses to hex strings
	if tx.FromEthAddr != nil && len(tx.FromEthAddr) > 0 {
		resp.FromEthAddr = "0x" + hex.EncodeToString(tx.FromEthAddr)
	}

	if tx.ToEthAddr != nil && len(tx.ToEthAddr) > 0 {
		resp.ToEthAddr = "0x" + hex.EncodeToString(tx.ToEthAddr)
	}

	return resp
}

// GetAllTransactions handles the API request to get all transactions
func (a *API) GetAllTransactions(c *gin.Context) {
	// Get all transactions from database
	txs, err := a.db.GetAllTxs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transactions: " + err.Error()})
		return
	}

	// Convert to response format
	txResponses := make([]TxResponse, len(txs))
	for i, tx := range txs {
		txResponses[i] = convertTxToResponse(tx)
	}

	// Return response
	c.JSON(http.StatusOK, gin.H{"transactions": txResponses})
}
