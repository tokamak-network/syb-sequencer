package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tokamak-network/syb-sequencer/sequencer/db/historydb"
)

// API represents the API server
type API struct {
	router *gin.Engine
	db     *historydb.HistoryDB
}

// NewAPI creates a new API server
func NewAPI(db *historydb.HistoryDB) *API {
	router := gin.Default()
	api := &API{
		router: router,
		db:     db,
	}

	// Set up routes
	api.setupRoutes()

	return api
}

// setupRoutes configures the API routes
func (a *API) setupRoutes() {
	// API version group
	v1 := a.router.Group("/api/v1")
	{
		// Transaction endpoint - simple version that returns all transactions
		v1.GET("/transactions", a.GetAllTransactions)
		// Transaction endpoint for a specific account
		v1.GET("/transactions/:accountAddress", a.GetTransactionsByAccount)
		// Transaction endpoint for paginated and sorted list
		v1.GET("/transactions/list", a.GetTransactionsPaginated)
		// Transaction endpoint to get transaction by hash
		v1.GET("/transactions/hash/:txHash", a.GetTransactionByHash)

		// Add other endpoints as needed
	}
}

// Run starts the API server
func (a *API) Run(addr string) error {
	return a.router.Run(addr)
}
