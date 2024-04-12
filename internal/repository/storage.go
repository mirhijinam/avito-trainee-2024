package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/mirhijinam/avito-trainee-2024/internal/service"
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

func (br BannerRepository) ExistsFeatureId(b *service.Banner) error {
	var exists bool
	tagQuery := `SELECT EXISTS(SELECT 1 FROM features WHERE id = $1)`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := br.db.QueryRowContext(ctx, tagQuery, b.FeatureId).Scan(&exists)
	if err != nil {
		log.Println("failed to get featureId")
		return err
	}
	if !exists {
		return fmt.Errorf("tag with ID %d does not exist", b.FeatureId)
	}

	return nil
}

func (br BannerRepository) ExistsTagId(b *service.Banner) error {
	var exists bool
	tagQuery := `SELECT EXISTS(SELECT 1 FROM tags WHERE id = $1)`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := br.db.QueryRowContext(ctx, tagQuery, b.TagId).Scan(&exists)
	if err != nil {
		log.Println("failed to get tagId")
		return err
	}
	if !exists {
		return fmt.Errorf("tag with ID %d does not exist", b.TagId)
	}

	return nil
}

func (br BannerRepository) InsertBanner(b *service.Banner) error {
	bannerQuery := `
		INSERT INTO banners (feature_id, additional_info, is_active)
		VALUES ($1, $2, $3)
		RETURNING id`
	args := []interface{}{b.FeatureId, b.AdditionalInfo, b.IsActive}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := br.db.QueryRowContext(ctx, bannerQuery, args...).Scan(&b.Id)
	if err != nil {
		return err
	}

	relBannersTagsQuery := `
		INSERT INTO banners_tags (banner_id, tag_id)
		VALUES ($1, $2)`

	stmt, err := br.db.Prepare(relBannersTagsQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(b.Id, b.TagId)
	if err != nil {
		return err
	}

	return nil
}
