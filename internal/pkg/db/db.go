package db

import (
	"context"
	"fmt"
	"log"

	"github.com/mirhijinam/avito-trainee-2024/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func MustOpenDB(ctx context.Context, cfg config.DBConfig) *pgxpool.Pool {
	// create default pool config
	config, err := pgxpool.ParseConfig("")
	if err != nil {
		log.Fatalf("failed to parse config (OpenDB): %w", err)
	}

	// fill pool config
	config.ConnConfig.User = cfg.PgUser
	config.ConnConfig.Password = cfg.PgPass
	config.ConnConfig.Host = cfg.PgHost
	config.ConnConfig.Port = cfg.PgPort
	config.ConnConfig.Database = cfg.PgDb

	// try to create new connection pool
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("failed to connect (OpenDB): %w", err)
	}

	// check is it alive
	if err = pool.Ping(ctx); err != nil {
		log.Fatalf("failed to ping (OpenDB): %w", err)
	}

	fmt.Println("success")
	return pool
}
