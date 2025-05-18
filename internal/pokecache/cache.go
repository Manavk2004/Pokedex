package pokecache

import(
	"sync"
	"time"
)

type cacheEntry struct{
	createdAt time.Time
	Val []byte
}


type Cache struct{
	store map[string]cacheEntry
	mux *sync.Mutex
}


func NewCache(interval time.Duration) Cache{
	c := Cache{
		store: make(map[string]cacheEntry),
		mux: &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}


func(c *Cache) Add(key string, value []byte){
	c.mux.Lock()
	defer c.mux.Unlock()
	c.store[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		Val: value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool){
	c.mux.Lock()
	defer c.mux.Unlock()
	val, ok := c.store[key]
	return val.Val, ok
}

func (c *Cache) reapLoop(interval time.Duration){
	ticker := time.NewTicker(interval)
	for range ticker.C{
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration){
	c.mux.Lock()
	defer c.mux.Unlock()
	for k,v := range c.store{
		if v.createdAt.Before(now.Add(-last)){
			delete(c.store, k)
		}
	}
}
