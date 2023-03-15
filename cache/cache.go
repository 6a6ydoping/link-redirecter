package cache

import (
	"log"
	"sync"
	"technodom_test/middlewares"
	"technodom_test/models"
	"time"
)

var UserCache *Cache

type Cache struct {
	items    map[string]cacheItem
	capacity int
	mutex    sync.RWMutex
}

type cacheItem struct {
	value      string
	expiration int64
}

type ICache interface {
	Add(key, value string, ttl time.Duration)
	Get(key string) (value string, ok bool)
	Len() int
}

func NewCache(capacity int) *Cache {
	c := &Cache{
		items:    make(map[string]cacheItem),
		capacity: capacity,
	}
	go c.StartCleaner()
	return c
}

func (c *Cache) Get(key string) (string, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	item, ok := c.items[key]
	if !ok {
		return "", false
	}

	if item.expiration > 0 && time.Now().Unix() > item.expiration {
		delete(c.items, key)
		return "", false
	}

	return item.value, true
}

func (c *Cache) Add(key string, value string, ttl time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if len(c.items) >= c.capacity {
		c.removeOldest()
	}

	expiration := int64(0)
	if ttl > 0 {
		expiration = time.Now().Add(ttl).Unix()
	}
	c.items[key] = cacheItem{
		value:      value,
		expiration: expiration,
	}
}

func (c *Cache) StartCleaner() {
	for {
		time.Sleep(time.Minute)
		c.mutex.Lock()
		for key, item := range c.items {
			if item.expiration > 0 && time.Now().Unix() > item.expiration {
				delete(c.items, key)
			}
		}
		c.mutex.Unlock()
	}
}

func (c *Cache) Len() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return len(c.items)
}

func (c *Cache) removeOldest() {
	var oKey string
	var oExpiration int64

	for key, item := range c.items {
		if oExpiration == 0 || item.expiration < oExpiration {
			oKey = key
			oExpiration = item.expiration
		}
	}
	delete(c.items, oKey)
}

func (c *Cache) WarmUpCache() {
	var links []models.Link
	links, err := middlewares.FindLinksByCategory("smartfony")
	if err != nil {
		log.Printf("Error retrieving links: %v\n", err)
		return
	}
	for _, link := range links {
		c.Add(link.ActiveLink, link.HistoryLink, 24*time.Minute)
	}
}
