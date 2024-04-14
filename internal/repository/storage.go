package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"
	"github.com/mirhijinam/avito-trainee-2024/internal/api"
	"github.com/mirhijinam/avito-trainee-2024/internal/service"
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
		return fmt.Errorf("feature with ID %d does not exist", b.FeatureId)
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

func (br BannerRepository) InsertBanner(b *service.Banner) (int, error) {
	fmt.Println("debug! content_v1:", string(b.Versions.ContentV1))
	fmt.Println("debug! content_v2:", string(b.Versions.ContentV2))
	fmt.Println("debug! content_v3:", string(b.Versions.ContentV3))

	bannerQuery := `
		INSERT INTO banners (feature_id, content, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	contentData, err := json.Marshal(map[string]json.RawMessage{
		"v1": b.Versions.ContentV1,
		"v2": b.Versions.ContentV2,
		"v3": b.Versions.ContentV3,
	})
	if err != nil {
		return -1, fmt.Errorf("failed to serialize content: %w", err)
	}

	args := []interface{}{
		b.FeatureId,
		contentData,
		b.IsActive,
		b.CreatedAt,
		b.UpdatedAt}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int
	err = br.db.QueryRowContext(ctx, bannerQuery, args...).Scan(&id)
	b.Id = id
	fmt.Println("debug! repo: b.Id", b.Id)
	if err != nil {
		return -1, err
	}

	relBannersTagsQuery := `
		INSERT INTO banners_tags (banner_id, tag_id)
		VALUES ($1, $2)`

	stmt, err := br.db.Prepare(relBannersTagsQuery)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(b.Id, b.TagId)
	if err != nil {
		return -1, err
	}

	return id, err
}

func (br BannerRepository) SelectBannerList(args []interface{}) ([]interface{}, error) {
	bannerListQuery := `
		SELECT
			banners.id AS banner_id,
			banners.feature_id,
			banners.content AS content,
			banners.is_active,
			banners.created_at,
			banners.updated_at,
			array_agg(banners_tags.tag_id) FILTER (WHERE banners_tags.tag_id IS NOT NULL) AS tag_ids
		FROM
			banners
		LEFT JOIN banners_tags ON banners.id = banners_tags.banner_id
		WHERE
			1 = 1
			AND ($1::integer IS NULL OR banners.feature_id = $1::integer)
			AND ($2::integer IS NULL OR banners_tags.tag_id = $2::integer)
		GROUP BY
			banners.id, banners.feature_id, banners.content, banners.is_active, banners.created_at, banners.updated_at
		ORDER BY
			banners.id
		LIMIT $3::integer OFFSET $4::integer;
		 `
	stmt, err := br.db.Prepare(bannerListQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bannerResponseList []interface{}
	for rows.Next() {
		var bResponse api.BannerResponse
		var tagIds []sql.NullInt32
		err := rows.Scan(
			&bResponse.BannerId,
			&bResponse.FeatureId,
			&bResponse.Versions,
			&bResponse.IsActive,
			&bResponse.CreatedAt,
			&bResponse.UpdatedAt,
			pq.Array(&tagIds),
		)
		if err != nil {
			fmt.Printf("Error scanning rows: %v\n", err)
			return nil, err
		}
		fmt.Printf("debug! banner id: %d, content: %s, tag_ids: %v\n", bResponse.BannerId, bResponse.Versions, tagIds)

		for _, tag := range tagIds {
			if tag.Valid {
				bResponse.TagIds = append(bResponse.TagIds, int(tag.Int32))
			}
		}

		bannerResponseList = append(bannerResponseList, bResponse)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bannerResponseList, nil
}

func (br BannerRepository) SelectBannerFromDB(args []interface{}) (bool, json.RawMessage, error) {
	bannerQuery := `
	SELECT b.is_active, b.content FROM banners b
	INNER JOIN 
		banners_tags bt ON b.id = bt.banner_id
	WHERE b.feature_id = $1 AND bt.tag_id = $2;       
    `
	stmt, err := br.db.Prepare(bannerQuery)
	if err != nil {
		return false, nil, err
	}
	defer stmt.Close()

	var isActive bool
	var content json.RawMessage

	err = stmt.QueryRow(args...).Scan(&isActive, &content)
	if err != nil {
		return false, nil, err
	}

	return isActive, content, nil
}
