package storage

import (
	"sync"
	"time"
)

type CacheInterface interface {
	Set(key string, value interface{}, ttl time.Duration)
	Get(key string) (interface{}, bool)
	Delete(key string)
	Clear()
}

type cacheItem struct {
	Data    interface{}
	Expires time.Time
}

type cache struct {
	items map[string]cacheItem
	mutex sync.RWMutex
	ttl   time.Duration
}

var _ CacheInterface = &cache{}

func NewCache() *cache {
	cache := &cache{
		items: make(map[string]cacheItem),
	}

	return cache
}

func (c *cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	expires := time.Now().Add(ttl)

	item := cacheItem{
		Data:    value,
		Expires: expires,
	}

	c.items[key] = item
}

func (c *cache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	item, found := c.items[key]

	if !found {
		return nil, false
	}

	if time.Now().After(item.Expires) {
		delete(c.items, key)

		return nil, false
	}

	return item.Data, true
}

func (c *cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.items, key)
}

func (c *cache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.items = make(map[string]cacheItem)
}
