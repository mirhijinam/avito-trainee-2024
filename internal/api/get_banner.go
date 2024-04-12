package api

import (
	"fmt"
	"net/http"
)

type getBannerRequest struct {
	FeatureId       int   `json:"feature_id"`
	TagId           []int `json:"tag_ids"`
	UseLastRevision bool  `json:"use_last_revision"`
}

func (h *Handler) GetBanner() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("get banner handler")
		return
	}
}
