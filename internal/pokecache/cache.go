package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	dict map[string]cacheEntry	
	mut  *sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val		  []byte
}


func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		dict: make(map[string]cacheEntry),
		mut: &sync.RWMutex{},
	}

	go cache.reapLoop(interval)
	return cache
}


func (c *Cache) Add(key string, value []byte) {
	c.mut.Lock()
	defer c.mut.Unlock()
	c.dict[key] = cacheEntry{
		createdAt: time.Now(),
		val: value,
	}
}


func (c *Cache) Get(key string) ([]byte, bool){
	c.mut.RLock()
	defer c.mut.RUnlock()
	entry, ok := c.dict[key] 
	if !ok {return nil, false}
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration){
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}


func (c *Cache) reap(interval time.Duration) {
	c.mut.Lock()
	defer c.mut.Unlock()
	timePassed := time.Now().UTC().Add(-interval)
	for key, entry := range c.dict{
		if entry.createdAt.Before(timePassed){
			delete(c.dict, key)
		}
	}

}