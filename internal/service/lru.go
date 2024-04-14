package service

import (
	"log"

	lru "github.com/hashicorp/golang-lru/v2"
)

type feature_tag struct {
	feature_id int
	tag_id     int
}

func createLRU(sz int) *lru.Cache[feature_tag, Banner] {
	cache, err := lru.New[feature_tag, Banner](sz)
	if err != nil {
		log.Fatalf("error! failed to create lru cache: %v", err)
	}

	return cache
}
