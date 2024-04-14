package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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
		fmt.Println("debug! i'am in the getbanner")
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
		path := r.URL.Path
		parts := strings.Split(path, "/")
		fmt.Println("debug!", len(parts))

		version := 1
		if len(parts) > 2 {
			if v, err := strconv.Atoi(parts[2][1:]); err == nil {
				version = v
			}
		}

		fmt.Println("debug! version =", version)
		if validateJWT(inpToken) < 1 && version > 1 {
			h.forbiddenAccessResponse(w, r)
			return
		}

		res := make(map[string]interface{})
		if queryMap["useLastRevision"] == true { // DB
			fmt.Println("debug! checking the use last revision true")
			isActive, bodyContain, err := h.BannerService.GetBannerFromDB(queryMap, version)

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

			isActive, bodyContain, err := h.BannerService.GetBannerFromCache(queryMap, version)
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
