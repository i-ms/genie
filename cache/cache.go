package cache

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type Cache interface {
	Has(string) (bool, error)
	Get(string) (interface{}, error)
	Set(string, interface{}, ...int) error
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
	
	ok, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false, err
	}
	return ok, nil
}

// 
func encode(item Entry) ([]byte, error) {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(item)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func decode(str string) (Entry, error) {
	item := Entry{}
	b := bytes.Buffer{}
	b.Write([]byte(str))
	d := gob.NewDecoder(&b)
	err := d.Decode(&item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// Get return value from redis along with error if any
func (c *RedisCache) Get(str string) (interface{}, error) {
	key := fmt.Sprintf("%s:%s", c.Prefix, str)
	conn := c.Conn.Get()
	defer conn.Close()
	
	cacheEntry, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	
	decoded, err := decode(string(cacheEntry))
	if err != nil {
		return nil, err
	}
	item := decoded[key]
	return item, nil
}

func (c *RedisCache) Set(str string, value interface{}, expires ...int) error {
	key := fmt.Sprintf("%s:%s", c.Prefix, str)
	conn := c.Conn.Get()
	defer conn.Close()
	
	entry := Entry{}
	entry[key] = value
	encoded, err := encode(entry)
	if err != nil {
		return err
	}
	
	if len(expires) > 0 {
		// Set with expiration time
		_, err := conn.Do("SETEX", key, expires[0], encoded)
		if err != nil {
			return err
		}
	} else {
		// if no expire is provided
		_, err := conn.Do("SET", key, encoded)
		if err != nil {
			return err
		}
	}
	
	return nil
}

func (c *RedisCache) Remove(str string) error {
	return nil
}

func (c *RedisCache) ClearByPattern(str string) error {
	return nil
}

func (c *RedisCache) Empty() error {
	return nil
}
