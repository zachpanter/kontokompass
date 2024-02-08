package storage

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/zachpanter/kontokompass/internal/config"
	"log"
	"runtime"
)

func OpenDBPool(ctx context.Context) *Queries {
	conf := config.NewConfig()
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=verify-full", conf.DBUser, conf.DBPass, conf.DBHost, conf.DBPort, conf.DBSchema)
	dbPool, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to connect DB", err)
	}

	pingErr := dbPool.PingContext(ctx)
	if pingErr != nil {
		log.Fatal("failed to ping DB", pingErr)
	}
	dbPool.SetMaxOpenConns(runtime.NumCPU())
	dbPool.SetMaxIdleConns(runtime.NumCPU())

	queries := New(dbPool)
	return queries
}
