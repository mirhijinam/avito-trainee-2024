package api

import (
	"net/http"
)

type getBannerListRequest struct {
	FeatureId int   `json:"feature_id"`
	TagId     []int `json:"tag_ids"`
	Limit     int   `json:"limit"`
	Offset    int   `json:"offset"`
}

func (h *Handler) GetBannerList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
