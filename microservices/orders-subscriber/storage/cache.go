package storage

import (
	"sync"
	"time"
)

type CacheInterface interface {
	Set(string, interface{}, time.Duration)
	Get(string) (interface{}, bool)
	ForEach(func(key string, value interface{}))
	Delete(string)
	Clear()
}

type CacheItem struct {
	Data    interface{}
	Expires time.Time
}

type Cache struct {
	items map[string]CacheItem
	mutex sync.RWMutex
	ttl   time.Duration
}

var _ CacheInterface = &Cache{}

func ConstructorCache() *Cache {
	cache := &Cache{
		items: make(map[string]CacheItem),
	}

	return cache
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	expires := time.Now().Add(ttl)

	item := CacheItem{
		Data:    value,
		Expires: expires,
	}

	c.items[key] = item
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

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

func (c *Cache) ForEach(callback func(key string, value interface{})) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for key, item := range c.items {
		if time.Now().After(item.Expires) {
			delete(c.items, key)

			continue
		}

		callback(key, item.Data)
	}
}

func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.items, key)
}

func (c *Cache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.items = make(map[string]CacheItem)
}
