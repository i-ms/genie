package cache

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

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

// Has return true if value exist in redis else false along with error if any
func (c *RedisCache) Has(str string) (bool, error) {
	key := fmt.Sprintf("%s:%s", c.Prefix, str)
	conn := c.Conn.Get()
	defer conn.Close()

	ok, err := redis.Bool(conn.Do("EXIST", key))
	if err != nil {
		return false, err
	}
	return ok, nil
}
