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
	Content json.RawMessage `json:"content"`
}

func (h *Handler) GetBanner() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		inpToken := r.Header.Get("token")
		fmt.Println("debug! recieved token:", inpToken)
		if validateJWT(inpToken) == -1 {
			h.userUnauthorizedResponse(w, r)
			return
		}

		queryMap, err := readJSONGetBannerQuery(r)
		fmt.Println("debug! recieved query map:", queryMap)
		if err != nil {
			fmt.Errorf("error! failed to read json request")
			return
		}

		res := make(map[string]interface{})
		if queryMap["useLastRevision"] == true { // DB
			fmt.Println("debug! checking the use last revision")
			isActive, bodyContain, err := h.BannerService.GetBannerFromDB(queryMap)

			if err != nil {
				h.badRequestResponse(w, r, err)
			} else if validateJWT(inpToken) == 0 && !isActive {
				h.forbiddenAccessResponse(w, r)
			} else {
				res["content"] = bodyContain
				h.successGetBannerResponse(w, r, res)
			}

		} else { // LRU Cache
			fmt.Println("debug! oops...")
		}

		return
	}
}
