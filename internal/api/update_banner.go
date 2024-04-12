package api

import (
	"net/http"
)

type updateBannerRequest struct {
	Id int `json:"id"`
}

func (h *Handler) UpdateBanner() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
