package cache

import (
	"time"
)

type Cache struct {
	storage map[string]interface{}
}

func New() *Cache {
	return &Cache{storage: make(map[string]interface{})}
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	to := time.After(ttl)

	c.storage[key] = value

	done := make(chan bool, 1)

	go func() {
		for {
			<-to
			done <- true
			delete(c.storage, key)
			return

		}
	}()

	<-done
}

func (c *Cache) Get(key string) interface{} {
	value, ok := c.storage[key]
	if ok {
		return value
	}
	return ok
}

func (c *Cache) Delete(key string) {
	delete(c.storage, key)
}
