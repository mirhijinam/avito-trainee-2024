package api

import (
	"encoding/json"
	"net/http"

	"github.com/mirhijinam/avito-trainee-2024/internal/service"
)

type createBannerRequest struct {
	FeatureId      int             `json:"feature_id"`
	TagIds         []int           `json:"tag_ids"`
	AdditionalInfo json.RawMessage `json:"content"`
	IsActive       bool            `json:"is_active"`
}

func (h *Handler) CreateBanner() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input := createBannerRequest{}
		inpToken := r.Header.Get("token")
		if validateJWT(inpToken) == 0 {
			h.forbiddenAccessResponse(w, r)
			return
		} else if validateJWT(inpToken) == -1 {
			h.userUnauthorizedResponse(w, r)
			return
		}

		err := readJSON(w, r, &input)
		if err != nil {
			h.badRequestResponse(w, r, err)
		}

		for _, tagId := range input.TagIds {
			b := service.Banner{
				FeatureId:      input.FeatureId,
				TagId:          tagId,
				AdditionalInfo: input.AdditionalInfo,
				IsActive:       input.IsActive,
			}
			err := h.BannerService.CreateBanner(b)
			if err != nil {
				h.badRequestResponse(w, r, err)
			} else {
				h.successResponse(w, r)
			}
		}

	}
}
