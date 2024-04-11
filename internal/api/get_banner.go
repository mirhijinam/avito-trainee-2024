package api

import (
	"net/http"
)

type getBannerRequest struct {
	FeatureId       int   `json:"feature_id"`
	TagId           []int `json:"tag_ids"`
	UseLastRevision bool  `json:"use_last_revision"`
}

func getBanner(bs BannerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
