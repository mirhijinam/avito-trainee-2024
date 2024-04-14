package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type getBannerListRequest struct {
	FeatureId *int `json:"feature_id"`
	TagId     *int `json:"tag_id"`
	Limit     *int `json:"limit"`
	Offset    *int `json:"offset"`
}

type BannerResponse struct {
	BannerId  int             `json:"banner_id"`
	TagIds    []int           `json:"tag_ids"`
	FeatureId int             `json:"feature_id"`
	Versions  json.RawMessage `json:"content"`
	IsActive  bool            `json:"is_active"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

func (h *Handler) GetBannerList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		inpToken := r.Header.Get("token")
		if validateJWT(inpToken) == 0 {
			h.forbiddenAccessResponse(w, r)
			return
		} else if validateJWT(inpToken) == -1 {
			h.userUnauthorizedResponse(w, r)
			return
		}

		queryMap, err := readJSONGetBannerListQuery(r)
		if err != nil {
			fmt.Errorf("error! failed to read json request")
			return
		}

		bodyContain, err := h.BannerService.GetBannerList(queryMap)
		if err != nil {
			fmt.Errorf("error! failed to GetBannerList")
			return
		}
		res := make(map[string]interface{})
		for _, item := range bodyContain {
			bResponse, ok := item.(BannerResponse)
			if !ok {
				fmt.Errorf("error! interface{} -/-> BannerResponse")
			}
			res["banner_id"] = bResponse.BannerId
			res["feature_id"] = bResponse.FeatureId
			res["tag_ids"] = bResponse.TagIds
			res["content"] = bResponse.Versions
			res["is_active"] = bResponse.IsActive
			res["created_at"] = bResponse.CreatedAt
			res["updated_at"] = bResponse.UpdatedAt
		}

		if err != nil {
			h.badRequestResponse(w, r, err)
		} else {
			h.successGetBannerListResponse(w, r, res)
		}
	}
}
