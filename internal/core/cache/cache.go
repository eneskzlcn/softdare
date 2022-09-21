package cache

import "time"

type Cache interface {
	SetWithExpire(key interface{}, value interface{}, expiration time.Duration) error
	Get(key interface{}) (interface{}, error)
}
