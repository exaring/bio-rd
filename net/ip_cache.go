package net

import (
	"sync"
)

const (
	ipCachePreAlloc = 1500000
)

var (
	ipc *ipCache
)

func init() {
	ipc = newIPCache()
}

type ipCache struct {
	cache   map[IP]*IP
	cacheMu sync.Mutex
}

func newIPCache() *ipCache {
	return &ipCache{
		cache: make(map[IP]*IP, ipCachePreAlloc),
	}
}

func (ipc *ipCache) get(addr IP) *IP {
	ipc.cacheMu.Lock()

	if a, exists := ipc.cache[addr]; exists {
		ipc.cacheMu.Unlock()
		return a
	}

	ipc._set(addr)
	ipc.cacheMu.Unlock()
	return ipc.cache[addr]
}

func (ipc *ipCache) _set(addr IP) {
	ipc.cache[addr] = &addr
}