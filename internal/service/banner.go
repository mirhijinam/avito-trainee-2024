package service

type BannerRepository interface {
}

type BannerService struct {
	repo BannerRepository
}

func New(repo BannerRepository) *BannerService {
	return &BannerService{
		repo: repo,
	}
}
