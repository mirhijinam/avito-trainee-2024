package service

import (
	"encoding/json"
)

type Banner struct {
	Id             int
	FeatureId      int
	TagId          int
	AdditionalInfo json.RawMessage
	IsActive       bool
}

type BannerRepository interface {
	InsertBanner(*Banner) error
	ExistsTagId(*Banner) error
	ExistsFeatureId(*Banner) error
}

type BannerService struct {
	repo BannerRepository
}

func New(repo BannerRepository) *BannerService {
	return &BannerService{
		repo: repo,
	}
}

func (bs *BannerService) CreateBanner(b Banner) error {
	if err := bs.repo.ExistsTagId(&b); err != nil {
		return err
	}

	// TODO: need some banner validation
	if err := bs.repo.InsertBanner(&b); err != nil {
		return err
	}

	return nil
}
