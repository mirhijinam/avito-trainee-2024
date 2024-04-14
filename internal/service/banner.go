package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	lru "github.com/hashicorp/golang-lru/v2"
)

// TODO: decide how to dermine the size of the cache

type BannerService struct {
	repo  BannerRepository
	cache lru.Cache[feature_tag, Banner]
}

type Versions struct {
	ContentV1 json.RawMessage `json:"v1"`
	ContentV2 json.RawMessage `json:"v2"`
	ContentV3 json.RawMessage `json:"v3"`
}

type Banner struct {
	Id        int
	FeatureId int
	TagId     int
	Versions  Versions
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(repo BannerRepository, sz int) *BannerService {
	return &BannerService{
		repo:  repo,
		cache: *createLRU(sz),
	}
}

func (bs *BannerService) CreateBanner(b *Banner) error {
	if err := bs.repo.ExistsTagId(b); err != nil {
		return err
	} else if err := bs.repo.ExistsFeatureId(b); err != nil {
		return err
	}

	// TODO: banner validation

	id, err := bs.repo.InsertBanner(b)
	b.Id = id
	fmt.Println("debug! service: b.Id =", b.Id)

	if err != nil {
		return err
	}

	ft := feature_tag{
		b.FeatureId,
		b.TagId,
	}
	bs.cache.Add(ft, *b)
	createdBanner, _ := bs.cache.Peek(ft)
	contentOfBanner, _ := json.Marshal(&createdBanner.Versions.ContentV1)
	fmt.Println("debug! cache: created b.Id =", b.Id, string(contentOfBanner))

	return nil
}

func (bs *BannerService) GetBannerList(qm map[string]interface{}) ([]interface{}, error) {
	args := make([]interface{}, 4)
	for i, key := range []string{"featureId", "tagId", "limit", "offset"} {
		if val, ok := qm[key]; ok && val != nil {
			args[i] = val
		} else {
			args[i] = sql.NullInt64{Valid: false}
		}
	}

	ans, err := bs.repo.SelectBannerList(args)
	if err != nil {
		return nil, err
	}

	return ans, nil
}

func (bs *BannerService) GetBannerFromDB(qm map[string]interface{}, v int) (bool, json.RawMessage, error) {
	args := make([]interface{}, 2)
	args[0] = qm["featureId"]
	args[1] = qm["tagId"]

	isActive, content, err := bs.repo.SelectBannerFromDB(args, v)
	if err != nil {
		return false, nil, err
	}
	return isActive, content, nil
}

func (bs *BannerService) GetBannerFromCache(qm map[string]interface{}, v int) (bool, json.RawMessage, error) {
	args := feature_tag{
		qm["featureId"].(int),
		qm["tagId"].(int),
	}

	banner, ok := bs.cache.Get(args)
	if !ok {
		return false, nil, fmt.Errorf("failed to get from cache")
	}

	versionsJSON, err := json.Marshal(banner.Versions)
	if err != nil {
		fmt.Println("failed marshaling versions:", err)
	} else {
		fmt.Println("debug! raw content:", string(versionsJSON))
	}

	var versions Versions
	err = json.Unmarshal(versionsJSON, &versions)
	if err != nil {
		return false, nil, fmt.Errorf("failed to unmarshal content: %v", err)
	}

	var retContent json.RawMessage
	switch v {
	case 1:
		retContent = versions.ContentV1
	case 2:
		retContent = versions.ContentV2
	case 3:
		retContent = versions.ContentV3
	}

	return banner.IsActive, retContent, nil
}
