package cache

import (
	"container/list"
	"sync"
	"time"

	model "product/internal/model"
)

// Lock() (Write Lock): Exactly one writer, zero readers.
// RLock() (Read Lock): Multiple readers, zero writers.

// In stage 2 we are adding
// singleflight → prevent cache stampede
// TTL + jitter → avoid mass expiry
// LRU eviction → bounded memory

type entry struct {
	key    string
	value  *model.Product
	expiry time.Time
}

type Cache struct {
	rwMu      sync.RWMutex
	items     map[string]*list.Element
	evictList *list.List
	capacity  int
	ttl       time.Duration
}

func NewCache(cap int, ttl time.Duration) *Cache {
	return &Cache{
		items:     make(map[string]*list.Element),
		evictList: list.New(),
		capacity:  cap,
		ttl:       ttl,
	}
}

// READ FAST PATH
func (c *Cache) Get(key string) (*model.Product, bool) {
	c.rwMu.RLock()
	elem, ok := c.items[key]
	c.rwMu.RUnlock()

	if !ok {
		return nil, false
	}

	ent := elem.Value.(*entry)

	// TTL check
	if time.Now().After(ent.expiry) {
		c.Delete(key)
		return nil, false
	}

	// move to front (LRU)
	c.rwMu.Lock()
	c.evictList.MoveToFront(elem)
	c.rwMu.Unlock()

	return ent.value, true
}

// WRITE PATH
func (c *Cache) Set(key string, val *model.Product) {
	c.rwMu.Lock()
	defer c.rwMu.Unlock()

	// update existing
	if elem, ok := c.items[key]; ok {
		c.evictList.MoveToFront(elem)
		elem.Value.(*entry).value = val
		elem.Value.(*entry).expiry = time.Now().Add(c.jitterTTL())
		return
	}

	// add new
	ent := &entry{
		key:    key,
		value:  val,
		expiry: time.Now().Add(c.jitterTTL()),
	}

	elem := c.evictList.PushFront(ent)
	c.items[key] = elem

	// evict if over capacity
	if c.evictList.Len() > c.capacity {
		c.removeOldest()
	}
}

func (c *Cache) removeOldest() {
	elem := c.evictList.Back()
	if elem != nil {
		c.evictList.Remove(elem)
		ent := elem.Value.(*entry)
		delete(c.items, ent.key)
	}
}

func (c *Cache) Delete(key string) {
	c.rwMu.Lock()
	defer c.rwMu.Unlock()

	if elem, ok := c.items[key]; ok {
		c.evictList.Remove(elem)
		delete(c.items, key)
	}
}

func (c *Cache) jitterTTL() time.Duration {
	// add random 0–20% jitter
	jitter := time.Duration(int64(c.ttl) / 5)
	return c.ttl + time.Duration(time.Now().UnixNano()%int64(jitter))
}
