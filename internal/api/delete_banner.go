package api

import (
	"net/http"
)

type deleteBannerRequest struct {
	FeatureId int   `json:"feature_id"`
	TagId     []int `json:"tag_ids"`
	Limit     int   `json:"limit"`
	Offset    int   `json:"offset"`
}

func (h *Handler) DeleteBanner() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
