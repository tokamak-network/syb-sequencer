package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/tokamak-network/syb-sequencer/sequencer/api"
	"github.com/tokamak-network/syb-sequencer/sequencer/config"
	"github.com/tokamak-network/syb-sequencer/sequencer/db/historydb"
	"github.com/tokamak-network/syb-sequencer/sequencer/db/statedb"
	"github.com/tokamak-network/syb-sequencer/sequencer/forger"
	"github.com/tokamak-network/syb-sequencer/sequencer/synchronizer"
)

func main() {
	logger := log.New(os.Stdout, "SEQUENCER: ", log.LstdFlags|log.Lshortfile)
	logger.Println("Starting Sybil Sequencer...")

	// Load configuration
	cfg := config.LoadConfig()

	if err := os.MkdirAll(cfg.Path, 0755); err != nil {
		log.Fatalf("Error creating base statedb directory %s: %v", cfg.Path, err)
	}

	syncDbPath := filepath.Join(cfg.Path, "synchronizer")
	forgerDbPath := filepath.Join(cfg.Path, "forger")

	db, err := historydb.InitSQLDB(cfg.DBPort, cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		log.Fatalf("Error initializing sql db: %v", err)
	}

	// Connect to historyDB
	historyDB := historydb.NewHistoryDB(db, db, nil)
	logger.Println("Connected to database successfully")

	// Create StateDB for synchronizer
	synchronizerStateDB, err := statedb.NewStateDB(statedb.Config{
		Path:    syncDbPath,
		Keep:    cfg.Keep,
		Type:    statedb.TypeSynchronizer,
		NLevels: statedb.MaxNLevels,
	})
	if err != nil {
		log.Fatalf("Error initializing synchronizer state db: %v", err)
	}

	// Create StateDB for forger
	forgerStateDB, err := statedb.NewLocalStateDB(statedb.Config{
		Path:    forgerDbPath,
		Keep:    cfg.Keep,
		Type:    statedb.TypeBatchBuilder,
		NLevels: statedb.MaxNLevels,
	}, synchronizerStateDB)

	if err != nil {
		log.Fatalf("Error initializing forger state db: %v", err)
	}

	log.Println("StateDBs initialized successfully.")

	// Create Forger
	forger := forger.NewForger(historyDB, forgerStateDB, logger)

	// Create synchronizer
	sync, err := synchronizer.NewSynchronizer(cfg.EthereumRPC, cfg.ContractAddress, historyDB, synchronizerStateDB, logger, forger)
	if err != nil {
		logger.Fatalf("Failed to create synchronizer: %v", err)
	}

	// Create context that can be cancelled
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start synchronizer
	go func() {
		sync.Start(ctx)
	}()

	// Initialize and start API server
	apiServer := api.NewAPI(historyDB)
	go func() {
		// Use environment variable for API port or default to 8080
		apiPort := os.Getenv("API_PORT")
		if apiPort == "" {
			apiPort = "8080"
		}

		logger.Printf("Starting API server on port %s...", apiPort)
		if err := apiServer.Run(":" + apiPort); err != nil {
			logger.Fatalf("API server failed: %v", err)
		}
	}()

	// Wait for interrupt signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	logger.Println("Shutting down...")
}
