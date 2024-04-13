package service

import (
	"database/sql"
	"encoding/json"
	"time"
)

// TODO: implement getBanner{List}FromDB
// TODO: implement getBanner{List}FromLRUCache:
//   (https://github.com/karlseguin/ccache,
//    https://github.com/hashicorp/golang-lru)
// TODO: decide how to dermine the size of the cache

type BannerService struct {
	repo BannerRepository
}

type Banner struct {
	Id        int
	FeatureId int
	TagId     int
	Content   json.RawMessage
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BannerRepository interface {
	InsertBanner(*Banner) error
	SelectBannerListFromDB([]interface{}) ([]interface{}, error)
	ExistsTagId(*Banner) error
	ExistsFeatureId(*Banner) error
}

func New(repo BannerRepository) *BannerService {
	return &BannerService{
		repo: repo,
	}
}

func (bs *BannerService) CreateBanner(b Banner) error {
	if err := bs.repo.ExistsTagId(&b); err != nil {
		return err
	} else if err := bs.repo.ExistsFeatureId(&b); err != nil {
		return err
	}

	// TODO: need some banner validation
	if err := bs.repo.InsertBanner(&b); err != nil {
		return err
	}

	return nil
}

func (bs *BannerService) GetBannerListFromDB(m map[string]interface{}) ([]interface{}, error) {
	args := make([]interface{}, 4)
	for i, key := range []string{"featureId", "tagId", "limit", "offset"} {
		if val, ok := m[key]; ok && val != nil {
			args[i] = val
		} else {
			args[i] = sql.NullInt64{Valid: false}
		}
	}
	ans, err := bs.repo.SelectBannerListFromDB(args)
	if err != nil {
		return nil, err
	}

	return ans, nil
}

func (bs *BannerService) GetBannerListFromLRUCache(b Banner) error {
	return nil
}
