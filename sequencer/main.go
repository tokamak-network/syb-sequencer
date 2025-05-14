package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/tokamak-network/syb-sequencer/sequencer/api"
	"github.com/tokamak-network/syb-sequencer/sequencer/config"
	"github.com/tokamak-network/syb-sequencer/sequencer/db/historydb"
	"github.com/tokamak-network/syb-sequencer/sequencer/forger"
	"github.com/tokamak-network/syb-sequencer/sequencer/synchronizer"
)

func main() {
	logger := log.New(os.Stdout, "SEQUENCER: ", log.LstdFlags|log.Lshortfile)
	logger.Println("Starting Sybil Sequencer...")

	// Load configuration
	cfg := config.LoadConfig()

	db, err := historydb.InitSQLDB(cfg.DBPort, cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	if err != nil {
		log.Fatalf("Error initializing sql db: %w", err)
	}

	// Connect to database
	database := historydb.NewHistoryDB(db, db, nil)
	logger.Println("Connected to database successfully")

	// Create Forger
	forger := forger.NewForger(database, logger)

	// Create synchronizer
	sync, err := synchronizer.NewSynchronizer(cfg.EthereumRPC, cfg.ContractAddress, database, logger, forger)
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
	apiServer := api.NewAPI(database)
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
