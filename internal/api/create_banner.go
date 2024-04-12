package api

import (
	"net/http"
)

type createBannerRequest struct {
	FeatureId      int    `json:"feature_id"`
	TagId          []int  `json:"tag_ids"`
	AdditionalInfo string `json:"additional_info"`
	IsActive       bool   `json:"is_active"`
}

func (h *Handler) CreateBanner() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input := createBannerRequest{}
		err := readJSON(w, r, &input)
		if err != nil {
			h.badRequestResponse(w, r, err)
		}
	}
}
