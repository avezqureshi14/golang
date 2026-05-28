package main

import (
	"fmt"
	"sync"
	"time"
)

type item struct {
	value      interface{}
	expiryTime int64
}

type TTLCache struct {
	data map[string]item
	mu   sync.RWMutex
}

func NewTTLCache() *TTLCache {
	return &TTLCache{
		data: make(map[string]item),
	}
}

func (c *TTLCache) Put(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = item{
		value:      value,
		expiryTime: time.Now().Add(ttl).UnixNano(),
	}
}

func (c *TTLCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	it, exists := c.data[key]
	c.mu.RUnlock()

	if !exists {
		return nil, false
	}

	if time.Now().UnixNano() > it.expiryTime {
		c.mu.Lock()
		delete(c.data, key)
		c.mu.Unlock()
		return nil, false
	}
	return it.value, true
}

func main() {
	cache := NewTTLCache()

	cache.Put("a", "hello", 2*time.Second)

	val, ok := cache.Get("a")

	fmt.Println(val, ok)

	time.Sleep(time.Second * 3)
	val, ok = cache.Get("a")
	fmt.Println(val, ok)
}
