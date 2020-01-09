package net

import (
	"sync"
)

const (
	prefixCachePreAlloc = 1500000
)

var (
	pfxc *pfxCache
)

func init() {
	pfxc = newPfxCache()
}

type pfxCache struct {
	cache   map[Prefix]*Prefix
	cacheMu sync.Mutex
}

func newPfxCache() *pfxCache {
	return &pfxCache{
		cache: make(map[Prefix]*Prefix, prefixCachePreAlloc),
	}
}

func (pfxc *pfxCache) get(pfx Prefix) *Prefix {
	pfx.addr = pfx.addr.Dedup()
	pfxc.cacheMu.Lock()

	if p, exists := pfxc.cache[pfx]; exists {
		pfxc.cacheMu.Unlock()
		return p
	}

	pfxc._set(pfx)
	pfxc.cacheMu.Unlock()
	return pfxc.cache[pfx]
}

func (pfxc *pfxCache) _set(pfx Prefix) {
	pfxc.cache[pfx] = &pfx
}