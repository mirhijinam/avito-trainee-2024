package api

import (
	"net/http"

	"github.com/mirhijinam/avito-trainee-2024/internal/service"
)

type updateBannerRequest struct {
	Id int `json:"id"`
}

func updateBanner(bs *service.BannerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
