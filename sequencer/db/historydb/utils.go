package historydb

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/gobuffalo/packr/v2"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/tokamak-network/syb-sequencer/sequencer/common"
	"golang.org/x/sync/semaphore"
)

var migrations *migrate.PackrMigrationSource

var (
	dbClient *pg.DB
	once     sync.Once
)

func init() {
	migrations = &migrate.PackrMigrationSource{
		Box: packr.New("syb-db-migrations", "./migrations"),
	}
	ms, err := migrations.FindMigrations()
	if err != nil {
		panic(err)
	}
	if len(ms) == 0 {
		panic(fmt.Errorf("no SQL migrations found"))
	}
}

// MigrationsUp runs the SQL migration up
func MigrationsUp(db *sql.DB) error {
	nMigrations, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return common.Wrap(err)
	}
	fmt.Printf("successfully ran migration up: %d \n", nMigrations)
	return nil
}

// MigrationsDown runs the SQL migration down
func MigrationsDown(db *sql.DB) error {
	_, err := migrate.Exec(db, "postgres", migrations, migrate.Down)
	if err != nil {
		return common.Wrap(err)
	}
	fmt.Println("successfully ran migration down")
	return nil
}

// ConnectSQLDB connects to the SQL DB
func InitSQLDB(port int, host, user, password, name string) (*sql.DB, error) {
	// Establish Connection
	psqlConn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		name,
	)
	db, err := sql.Open("postgres", psqlConn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect SQL DB, %v", err)
	}
	// Run DB migrations
	if err := MigrationsUp(db); err != nil {
		return nil, fmt.Errorf("failed to run migration up, %v", err)
	}
	return db, nil
}

// GetDB returns the singleton database client instance.
// It assumes InitialiseDB has been called successfully at least once.
func GetDB() *pg.DB {
	if dbClient == nil {
		panic("database client is not initialized. Call InitialiseDB first.")
	}
	return dbClient
}

// APIConnectionController is used to limit the SQL open connections used by the API
type APIConnectionController struct {
	smphr   *semaphore.Weighted
	timeout time.Duration
}

// NewAPIConnectionController initializes the APIConnectionController
func NewAPIConnectionController(maxConnections int, timeout time.Duration) *APIConnectionController {
	return &APIConnectionController{
		smphr:   semaphore.NewWeighted(int64(maxConnections)),
		timeout: timeout,
	}
}
