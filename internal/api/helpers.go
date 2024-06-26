package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	_ "github.com/golang-jwt/jwt/v5"
)

type envelope map[string]interface{}

func checkToken(token string) int {
	adminToken := os.Getenv("ADMIN_TOKEN")
	userToken := os.Getenv("USER_TOKEN")

	admintokenret, err := strconv.Atoi(os.Getenv("ADMIN_TOKEN_RET"))
	if err != nil {
		fmt.Println("debug! failed to convert admin token return value")
		return 0
	}

	usertokenret, err := strconv.Atoi(os.Getenv("USER_TOKEN_RET"))
	if err != nil {
		fmt.Println("debug! failed to convert user token return value")
		return 0
	}

	unauthorizedret, err := strconv.Atoi(os.Getenv("UNAUTHORIZED_RET"))
	if err != nil {
		fmt.Println("debug! failed to convert unauthorized return value")
		return 0
	}

	switch token {
	case userToken:
		return usertokenret
	case adminToken:
		return admintokenret
	default:
		return unauthorizedret
	}
}

func readJSONGetBannerQuery(r *http.Request) (map[string]interface{}, error) {
	params := &getBannerRequest{}
	q := r.URL.Query()

	parseInt := func(key string) (*int, error) {
		if value, present := q[key]; present && len(value) > 0 {
			i, err := strconv.Atoi(value[0])
			if err != nil {
				return nil, err
			}
			return &i, nil
		}
		return nil, nil
	}

	parseBool := func(key string, defaultVal bool) (bool, error) {
		value, present := q[key]
		if !present || len(value) == 0 {
			return defaultVal, nil
		}
		b, err := strconv.ParseBool(value[0])
		if err != nil {
			return defaultVal, err
		}
		return b, nil
	}
	var err error
	if params.FeatureId, err = parseInt("feature_id"); err != nil {
		return nil, err
	}
	if params.TagId, err = parseInt("tag_id"); err != nil {
		return nil, err
	}
	x, err := parseBool("use_last_revision", false)
	params.UseLastRevision = &x
	if err != nil {
		return nil, err
	}

	queryMap := map[string]interface{}{
		"featureId":       nil,
		"tagId":           nil,
		"useLastRevision": false,
	}

	queryMap["featureId"] = *params.FeatureId
	queryMap["tagId"] = *params.TagId

	fmt.Println("debug! useLastRevision", *params.UseLastRevision)
	if params.UseLastRevision != nil {
		queryMap["useLastRevision"] = *params.UseLastRevision
	}

	return queryMap, nil
}

func readJSONGetBannerListQuery(r *http.Request) (map[string]interface{}, error) {
	params := &getBannerListRequest{}
	q := r.URL.Query()

	parseInt := func(key string) (*int, error) {
		if value, present := q[key]; present && len(value) > 0 {
			i, err := strconv.Atoi(value[0])
			if err != nil {
				return nil, err
			}
			return &i, nil
		}
		return nil, nil
	}

	var err error
	if params.FeatureId, err = parseInt("feature_id"); err != nil {
		return nil, err
	}
	if params.TagId, err = parseInt("tag_id"); err != nil {
		return nil, err
	}
	if params.Limit, err = parseInt("limit"); err != nil {
		return nil, err
	}
	if params.Offset, err = parseInt("offset"); err != nil {
		return nil, err
	}

	queryMap := map[string]interface{}{
		"featureId": nil,
		"tagId":     nil,
		"limit":     nil,
		"offset":    nil,
	}

	if params.FeatureId != nil {
		queryMap["featureId"] = *params.FeatureId
	}
	if params.TagId != nil {
		queryMap["tagId"] = *params.TagId
	}
	if params.Limit != nil {
		queryMap["limit"] = *params.Limit
	}
	if params.Offset != nil {
		queryMap["offset"] = *params.Offset
	}

	return queryMap, nil
}

func readJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	dec := json.NewDecoder(r.Body)

	dec.DisallowUnknownFields()
	err := dec.Decode(dst)

	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		// case errors.Is(err, io.EOF):
		//	return errors.New("body must not be empty")

		case err.Error() == "http: request body too large":
			return fmt.Errorf("body must not be larger than %d bytes", maxBytes)

		case errors.As(err, &invalidUnmarshalError):
			panic(err)

		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value")
	}

	return nil
}

func writeJSONBody(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	js = append(js, '\n')

	for key, val := range headers {
		w.Header()[key] = val
	}

	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(status)

	w.Write(js)

	return nil
}
