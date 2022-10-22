package utilz

import (
	"context"
	"time"

	"github.com/jellydator/ttlcache/v3"
)

type FetchFn func(context.Context) ([]byte, error)

type CacheProvidor interface {
	Get(key string) ([]byte, error)
	Set(key string, data []byte, expire time.Duration) error
}

func CacheGet(ctx context.Context, providor CacheProvidor, key string, expire time.Duration, callback FetchFn) ([]byte, error) {
	cachedData, err := providor.Get(key)
	if err == nil { // cache hit
		return cachedData, nil
	}
	if cachedData == nil { // cache not found
		if data, err := callback(ctx); err == nil { // fetch source
			return data, providor.Set(key, data, expire)
		} else {
			return nil, err
		}
	}
	return nil, err
}

func CacheHash(data any) string {
	return MD5String(ToJSONStringNoError(data))
}

var MemCacheProvidor *memCacheProvidor

type memCacheProvidor struct {
	cache *ttlcache.Cache[string, []byte]
}

func init() {
	cache := ttlcache.New[string, []byte]()
	go cache.Start()

	MemCacheProvidor = &memCacheProvidor{cache: cache}
}

func (p *memCacheProvidor) Get(key string) ([]byte, error) {
	item := p.cache.Get(key)
	if item == nil {
		return nil, nil
	} else {
		return item.Value(), nil
	}
}

func (p *memCacheProvidor) Set(key string, data []byte, expire time.Duration) error {
	p.cache.Set(key, data, expire)
	return nil
}
