package repository

import (
	"context"
	"database/sql"
	"time"
	// "github.com/lib/pq"
)

type BannerRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *BannerRepository {
	return &BannerRepository{
		db: db,
	}
}

type Banner struct {
	// what need we to send as a response
}

func (br BannerRepository) Insert(b *Banner) error {
	query := `
	`
	args := []interface{}{} // add smth
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return br.db.QueryRowContext(ctx, query, args...).Scan() // add that we return from the query
}
