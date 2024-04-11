package api

import (
	"net/http"

	"github.com/mirhijinam/avito-trainee-2024/internal/service"
)

type getBannerListRequest struct {
	FeatureId int   `json:"feature_id"`
	TagId     []int `json:"tag_ids"`
	Limit     int   `json:"limit"`
	Offset    int   `json:"offset"`
}

func getBannerList(bs *service.BannerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
