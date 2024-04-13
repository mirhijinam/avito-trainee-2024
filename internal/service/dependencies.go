package service

import "encoding/json"

type BannerRepository interface {
	InsertBanner(*Banner) (int, error)
	SelectBannerListFromDB([]interface{}) ([]interface{}, error)
	SelectBannerFromDB([]interface{}) (json.RawMessage, error)
	ExistsTagId(*Banner) error
	ExistsFeatureId(*Banner) error
}
