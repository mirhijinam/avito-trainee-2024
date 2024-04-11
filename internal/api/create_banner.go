package api

import (
	"net/http"

	"github.com/mirhijinam/avito-trainee-2024/internal/service"
)

type createBannerRequest struct {
	FeatureId      int    `json:"feature_id"`
	TagId          []int  `json:"tag_ids"`
	AdditionalInfo string `json:"additional_info"`
	IsActive       bool   `json:"is_active"`
}

func createBanner(bs *service.BannerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
