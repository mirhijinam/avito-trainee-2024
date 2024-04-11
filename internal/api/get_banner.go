package api

import (
	"net/http"

	"github.com/mirhijinam/avito-trainee-2024/internal/service"
)

type getBannerRequest struct {
	FeatureId       int   `json:"feature_id"`
	TagId           []int `json:"tag_ids"`
	UseLastRevision bool  `json:"use_last_revision"`
}

func getBanner(bs *service.BannerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
