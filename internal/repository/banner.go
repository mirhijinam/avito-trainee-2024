package repository

import (
	"context"
	"database/sql"
	"fmt"
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

func (br BannerRepository) Insert(b *service.Banner) error {
	bannerQuery := `
		INSERT INTO banners (feature_id, additional_info, is_active)
		VALUES ($1, $2, $3)
		RETURNING id`
	args := []interface{}{b.FeatureId, b.AdditionalInfo, b.IsActive}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var exists bool
	er := br.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM tags WHERE id = $1)", b.TagId).Scan(&exists)
	if er != nil {
		return er // handle error properly
	}
	if !exists {
		// Insert the tag or handle the error/condition
		// For example, return an error stating the tag does not exist
		return fmt.Errorf("tag with ID %d does not exist", b.TagId)
	}
	err := br.db.QueryRowContext(ctx, bannerQuery, args...).Scan(&b.Id)
	if err != nil {
		return err
	}

	relBannersTagsQuery := `
		INSERT INTO banners_tags (banner_id, tag_id)
		VALUES ($1, $2)
		RETURNING 0`

	check := 0
	err = br.db.QueryRowContext(ctx, relBannersTagsQuery, b.Id, b.TagId).Scan(&check)
	if err != nil {
		return err
	}

	return nil
}
