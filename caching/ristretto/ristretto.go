package ristretto

import (
	"github.com/dgraph-io/ristretto/v2"
)

// NewRistrettoClient Returns new ristretto in memory cache client
func NewRistrettoClient() (cache *ristretto.Cache[string, string], err error) {
	cache, err = ristretto.NewCache(&ristretto.Config[string, string]{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})

	return
}
