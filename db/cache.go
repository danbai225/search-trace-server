package db

import (
	"github.com/muesli/cache2go"
)

var cache *cache2go.CacheTable

func initCache() {
	cache = cache2go.Cache("cache")
}
func GetCache() *cache2go.CacheTable {
	return cache
}
