package api

import (
	"encoding/json"

	"github.com/mirhijinam/avito-trainee-2024/internal/service"
)

type BannerService interface {
	CreateBanner(b *service.Banner) error
	GetBannerList(map[string]interface{}) ([]interface{}, error)
	GetBannerFromDB(map[string]interface{}, int) (bool, json.RawMessage, error)
	GetBannerFromCache(map[string]interface{}, int) (bool, json.RawMessage, error)
}
