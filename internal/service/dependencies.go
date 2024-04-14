package service

import "encoding/json"

type BannerRepository interface {
	InsertBanner(*Banner) (int, error)
	SelectBannerList([]interface{}) ([]interface{}, error)
	SelectBannerFromDB([]interface{}, int) (bool, json.RawMessage, error)
	ExistsTagId(*Banner) error
	ExistsFeatureId(*Banner) error
}
