package cache

import (
	"sync"

	model "product/internal/model"
)

// Lock() (Write Lock): Exactly one writer, zero readers.
// RLock() (Read Lock): Multiple readers, zero writers.

type Cache struct {
	rwMu  sync.RWMutex
	items map[string]*model.Product
}

func NewCache() *Cache {
	return &Cache{
		items: make(map[string]*model.Product),
	}
}

// READ FAST PATH
func (c *Cache) Get(id string) (*model.Product, bool) {
	// multiple readers are allowed and no writer is allowed 
	// So it is like : I am reading this right now. Writers, do not touch this until I call RUnlock()."
	c.rwMu.RLock()
	defer c.rwMu.RUnlock()

	p, ok := c.items[id]
	return p, ok
}

// WRITE PATH
func (c *Cache) Set(id string, p *model.Product) {
	// when writing we are blocking all writers and readers,  only one writer  is allowed
	c.rwMu.Lock()
	defer c.rwMu.Unlock()

	c.items[id] = p
}