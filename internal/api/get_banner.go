package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type getBannerRequest struct {
	FeatureId       *int  `json:"feature_id"`
	TagId           *int  `json:"tag_ids"`
	UseLastRevision *bool `json:"use_last_revision"`
}

type BannerContentResponse struct {
	Content   json.RawMessage `json:"content"`
	ContentV2 json.RawMessage `json:"content_v2,omitempty"`
	ContentV3 json.RawMessage `json:"content_v3,omitempty"`
}

func (h *Handler) GetBanner() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		inpToken := r.Header.Get("token")
		fmt.Println("debug! received token:", inpToken)
		if validateJWT(inpToken) == -1 {
			h.userUnauthorizedResponse(w, r)
			return
		}

		queryMap, err := readJSONGetBannerQuery(r)
		fmt.Println("debug! received query map:", queryMap)
		if err != nil {
			fmt.Errorf("error! failed to read json request")
			return
		}

		res := make(map[string]interface{})
		if queryMap["useLastRevision"] == true { // DB
			fmt.Println("debug! checking the use last revision true")
			isActive, bodyContain, err := h.BannerService.GetBannerFromDB(queryMap)

			if err != nil {
				h.badRequestResponse(w, r, err)
			} else if validateJWT(inpToken) == 0 && !isActive {
				h.forbiddenAccessResponse(w, r)
			} else {
				res["content"] = bodyContain
				h.successGetBannerResponse(w, r, res)
			}

		} else { // Cache
			fmt.Println("debug! checking the use last revision false")
			isActive, bodyContain, err := h.BannerService.GetBannerFromCache(queryMap)
			if err != nil {
				h.badRequestResponse(w, r, err)
			} else if validateJWT(inpToken) == 0 && !isActive {
				h.forbiddenAccessResponse(w, r)
			} else {
				res["content"] = bodyContain
				h.successGetBannerResponse(w, r, res)
			}
		}

		return
	}
}
