package api

import (
	"fmt"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/tokamak-network/syb-sequencer/sequencer/abis/bindings"
	"github.com/tokamak-network/syb-sequencer/sequencer/db/historydb"
	"github.com/tokamak-network/syb-sequencer/sequencer/db/statedb"
)

// API represents the API server
type API struct {
	router        *gin.Engine
	db            *historydb.HistoryDB
	sybilContract *bindings.Sybil
	statedb       *statedb.LocalStateDB
}

// NewAPI creates a new API server
func NewAPI(db *historydb.HistoryDB, ethRPC, contractAddressHex string, statedb *statedb.LocalStateDB) (*API, error) {
	client, err := ethclient.Dial(ethRPC)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %v", err)
	}

	contractAddress := ethCommon.HexToAddress(contractAddressHex)
	sybilContract, err := bindings.NewSybil(contractAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate Sybil contract: %v", err)
	}

	router := gin.Default()
	api := &API{
		router:        router,
		db:            db,
		sybilContract: sybilContract,
		statedb:       statedb,
	}

	// Set up routes
	api.setupRoutes()

	return api, nil
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

		// Account routes
		v1.GET("/accounts", a.GetAllAccounts)
		v1.GET("/account/:idx", a.GetAccountByIdx)

		// Add other endpoints as needed
	}
}

// Run starts the API server
func (a *API) Run(addr string) error {
	return a.router.Run(addr)
}
