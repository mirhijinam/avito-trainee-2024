package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mirhijinam/avito-trainee-2024/internal/service"
)

// TODO: implement version control as {v1{...}, v2{...}, v3{...}}

type createBannerRequest struct {
	FeatureId int             `json:"feature_id"`
	TagIds    []int           `json:"tag_ids"`
	Content   json.RawMessage `json:"content"`
	IsActive  bool            `json:"is_active"`
}

func (h *Handler) CreateBanner() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input := createBannerRequest{}
		inpToken := r.Header.Get("token")
		if checkToken(inpToken) == usertokenret {
			h.forbiddenAccessResponse(w, r)
			return
		} else if checkToken(inpToken) == unauthorizedret {
			h.userUnauthorizedResponse(w, r)
			return
		}

		err := readJSONBody(w, r, &input)
		if err != nil {
			h.badRequestResponse(w, r, err)
			return
		}

		versions := service.Versions{
			ContentV1: input.Content,
			ContentV2: json.RawMessage("{}"),
			ContentV3: json.RawMessage("{}"),
		}
		for _, tagId := range input.TagIds {
			b := service.Banner{
				FeatureId: input.FeatureId,
				TagId:     tagId,
				Versions:  versions,
				IsActive:  input.IsActive,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}

			err := h.BannerService.CreateBanner(&b)
			if err != nil {
				h.badRequestResponse(w, r, err)
				return
			}
		}

		h.successBannerCreationResponse(w, r)
	}
}
