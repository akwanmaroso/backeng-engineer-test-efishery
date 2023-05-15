package cache

import (
	"sync"
	"time"
)

type Cache struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, ok := c.data[key]

	return value, ok
}

func (c *Cache) Set(key string, value interface{}, expiration time.Duration) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	c.data[key] = value
	if expiration > 0 {
		go func() {
			<-time.After(expiration)
			c.mu.Lock()
			defer c.mu.Unlock()

			delete(c.data, key)
		}()
	}
}
