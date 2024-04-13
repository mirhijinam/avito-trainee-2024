package service

import (
	"log"

	lru "github.com/hashicorp/golang-lru/v2"
)

func createLRU(sz int) *lru.Cache[int, Banner] {
	cache, err := lru.New[int, Banner](sz)
	if err != nil {
		log.Fatalf("Failed to create LRU cache: %v", err)
	}

	return cache
}
