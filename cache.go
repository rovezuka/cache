package cache

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type value struct {
	expire int64
	value  any
}

type Cache struct {
	storage map[string]value
	mu      sync.Mutex
}

func New() *Cache {
	return &Cache{storage: make(map[string]value)}
}

func (c *Cache) Set(key string, val any, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.storage[key] = value{
		expire: time.Now().Unix() + int64(ttl.Seconds()),
		value:  val,
	}
}

func (c *Cache) Get(key string) (any, error) {

	val, ok := c.storage[key]
	if ok {
		if time.Now().Unix() > c.storage[key].expire {
			delete(c.storage, key)
			return nil, fmt.Errorf("время жизни значения %s уже завершилось", key)
		}

		return val.value, nil

	}
	return nil, errors.New("данного ключа нет в хранилище")
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.storage, key)
}
