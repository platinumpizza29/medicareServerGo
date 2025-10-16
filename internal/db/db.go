package db

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	pool    *pgxpool.Pool
	once    sync.Once
	initErr error
)

// ConnectDB creates the global pgxpool.Pool once and verifies the connection.
func ConnectDB(ctx context.Context, url string) error {
	once.Do(func() {
		var err error
		pool, err = pgxpool.New(ctx, url)
		if err != nil {
			initErr = fmt.Errorf("failed to create pool: %w", err)
			return
		}
		if err = pool.Ping(ctx); err != nil {
			initErr = fmt.Errorf("failed to ping database: %w", err)
			return
		}
		fmt.Println("✅ Database connected successfully")
	})
	return initErr
}

// Pool returns the global connection pool.
// Panics if ConnectDB was never called.
func Pool() *pgxpool.Pool {
	if pool == nil {
		panic("database not initialized: call ConnectDB first")
	}
	return pool
}

// Close releases the pool for graceful shutdown.
func Close() {
	if pool != nil {
		pool.Close()
		fmt.Println("🛑 Database connection closed")
	}
}
