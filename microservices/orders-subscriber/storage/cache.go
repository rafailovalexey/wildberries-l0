package storage

import (
	"sync"
	"time"
)

type CacheInterface interface {
	GetCache() map[string]CacheItem
	Get(string) (interface{}, bool)
	Set(string, interface{}, time.Duration)
	Delete(string)
	Clear()
}

type CacheItem struct {
	Data    interface{}
	Expires time.Time
}

type cache struct {
	items map[string]CacheItem
	mutex sync.RWMutex
	ttl   time.Duration
}

var _ CacheInterface = &cache{}

func NewCache() *cache {
	cache := &cache{
		items: make(map[string]CacheItem),
	}

	return cache
}

func (c *cache) GetCache() map[string]CacheItem {
	return c.items
}

func (c *cache) Get(key string) (interface{}, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	item, isExist := c.items[key]

	if !isExist {
		return nil, false
	}

	if time.Now().After(item.Expires) {
		delete(c.items, key)

		return nil, false
	}

	return item.Data, true
}

func (c *cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	expires := time.Now().Add(ttl)

	item := CacheItem{
		Data:    value,
		Expires: expires,
	}

	c.items[key] = item
}

func (c *cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.items, key)
}

func (c *cache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.items = make(map[string]CacheItem)
}
