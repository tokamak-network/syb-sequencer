package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/tokamak-network/syb-sequencer/sequencer/config"
	"github.com/tokamak-network/syb-sequencer/sequencer/db/historydb"
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

	// Create synchronizer
	sync, err := synchronizer.NewSynchronizer(cfg.EthereumRPC, cfg.ContractAddress, database, logger)
	if err != nil {
		logger.Fatalf("Failed to create synchronizer: %v", err)
	}

	// Create context that can be cancelled
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start synchronizer
	sync.Start(ctx)

	// Wait for interrupt signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	logger.Println("Shutting down...")
}
