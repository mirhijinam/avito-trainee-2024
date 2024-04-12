package api

import "github.com/mirhijinam/avito-trainee-2024/internal/service"

type BannerService interface {
	CreateBanner(b service.Banner) error
}
