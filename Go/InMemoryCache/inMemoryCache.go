package main

import (
	containerlist
	sync
	time
)

type CacheEntry struct {
	key        string
	value      interface{}
	expiration time.Time
}

type LRUCache struct {
	mu       sync.Mutex
	capacity int
	items    map[string]list.Element
	evictList list.List
}

func NewLRUCache(capacity int) LRUCache {
	return &LRUCache{
		capacity capacity,
		items    make(map[string]list.Element),
		evictList list.New(),
	}
}

func (c LRUCache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if el, ok = c.items[key]; ok {
		c.evictList.MoveToFront(el)
		entry = el.Value.(CacheEntry)
		entry.value = value
		entry.expiration = time.Now().Add(ttl)
	} else {
		entry = &CacheEntry{
			key        key,
			value      value,
			expiration time.Now().Add(ttl),
		}
		el = c.evictList.PushFront(entry)
		c.items[key] = el

		if c.evictList.Len()  c.capacity {
			c.removeOldest()
		}
	}
}

func (c LRUCache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if el, ok = c.items[key]; ok {
		entry = el.Value.(CacheEntry)
		if time.Now().Before(entry.expiration) {
			c.evictList.MoveToFront(el)
			return entry.value, true
		}
		c.removeElement(el)
	}
	return nil, false
}

func (c LRUCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if el, ok = c.items[key]; ok {
		c.removeElement(el)
	}
}

func (c LRUCache) removeOldest() {
	el = c.evictList.Back()
	if el != nil {
		c.removeElement(el)
	}
}

func (c LRUCache) removeElement(el list.Element) {
	c.evictList.Remove(el)
	entry = el.Value.(CacheEntry)
	delete(c.items, entry.key)
}
