package cache

import "github.com/gomodule/redigo/redis"

type Cache interface {
	Has(string) (bool, error)
	Get(string) (interface{}, error)
	Set(string, interface{}, ...int)
	Remove(string) error
	ClearByPattern(string) error
	Empty() error
}

// RedisCache
// Conn   : Connection Pool to redis
// Prefix : Prevents cache key collision (between different applications)
type RedisCache struct {
	Conn   *redis.Pool
	Prefix string
}

// Entry is used to store the cache serialization information
type Entry map[string]interface{}
