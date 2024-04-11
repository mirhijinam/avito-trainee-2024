package service

import repo "github.com/mirhijinam/avito-trainee-2024/internal/repository"

type BannerService struct {
	repo *repo.BannerRepository
}

func New(repo *repo.BannerRepository) *BannerService {
	return &BannerService{
		repo: repo,
	}
}
