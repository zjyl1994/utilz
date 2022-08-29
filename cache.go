package utilz

import (
	"context"
	"time"

	"github.com/gomodule/redigo/redis"
)

type fetchFn func(context.Context) ([]byte, error)

func CacheGet(ctx context.Context, conn redis.Conn, key string, expire time.Duration, callback fetchFn) ([]byte, error) {
	cachedData, err := redis.Bytes(conn.Do("GET", key))
	if err == nil {// cache hit
		return cachedData, nil
	}
	if err == redis.ErrNil {// cache not found
		expireSecond := int(expire.Seconds())
		if data, err := callback(ctx); err == nil { // fetch source
			if expireSecond > 0 { // expire time set
				_, err = conn.Do("SETEX", key, expireSecond, data)
			} else {
				_, err = conn.Do("SET", key, data)
			}
			return data, err
		} else {
			return nil, err
		}
	}
	return nil, err
}

func CacheHash(data any) string {
	return MD5String(ToJSONStringNoError(data))
}
