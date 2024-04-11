package api

import (
	"net/http"
)

type updateBannerRequest struct {
	Id int `json:"id"`
}

func updateBanner(bs BannerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
