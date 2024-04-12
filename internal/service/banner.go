package service

import "encoding/json"

type Banner struct {
	Id             int
	FeatureId      int
	TagId          int
	AdditionalInfo json.RawMessage
	IsActive       bool
}

type BannerRepository interface {
	Insert(banner *Banner) error
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
	// need some validation
	err := bs.repo.Insert(&b)
	if err != nil {
		return err
	}

	return nil
}

// TODO: somehow implement CreateBannerTag (for relation table)
