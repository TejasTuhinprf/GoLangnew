package main

import (
	"container/list"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

type CacheEntry struct {
	key        string
	value      interface{}
	expiration time.Time
}

type LRUCache struct {
	mu        sync.Mutex
	capacity  int
	items     map[string]*list.Element
	evictList *list.List
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity:  capacity,
		items:     make(map[string]*list.Element),
		evictList: list.New(),
	}
}

// Set with TTL
func (c *LRUCache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if el, ok := c.items[key]; ok {
		c.evictList.MoveToFront(el)
		entry := el.Value.(*CacheEntry)
		entry.value = value
		entry.expiration = time.Now().Add(ttl)
	} else {
		entry := &CacheEntry{
			key:        key,
			value:      value,
			expiration: time.Now().Add(ttl),
		}
		el := c.evictList.PushFront(entry)
		c.items[key] = el

		if c.evictList.Len() > c.capacity {
			c.removeOldest()
		}
	}
}

// Get
func (c *LRUCache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if el, ok := c.items[key]; ok {
		entry := el.Value.(*CacheEntry)
		if time.Now().Before(entry.expiration) {
			c.evictList.MoveToFront(el)
			return entry.value, true
		}
		c.removeElement(el)
	}
	return nil, false
}

// Delete
func (c *LRUCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if el, ok := c.items[key]; ok {
		c.removeElement(el)
	}
}

func (c *LRUCache) removeOldest() {
	el := c.evictList.Back()
	if el != nil {
		c.removeElement(el)
	}
}

func (c *LRUCache) removeElement(el *list.Element) {
	c.evictList.Remove(el)
	entry := el.Value.(*CacheEntry)
	delete(c.items, entry.key)
}

// Redis
type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(addr, password string, db int) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisCache{client: rdb}
}

func (c *RedisCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	return c.client.Set(ctx, key, value, ttl).Err()
}

func (c *RedisCache) Get(ctx context.Context, key string) (interface{}, error) {
	return c.client.Get(ctx, key).Result()
}

func (c *RedisCache) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}

type Cache interface {
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string) (interface{}, error)
	Delete(ctx context.Context, key string) error
}

type MultiBackendCache struct {
	inMemoryCache Cache
	redisCache    Cache
}

func NewMultiBackendCache(inMemoryCache Cache, redisCache Cache) *MultiBackendCache {
	return &MultiBackendCache{
		inMemoryCache: inMemoryCache,
		redisCache:    redisCache,
	}
}

func (c *MultiBackendCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	if err := c.inMemoryCache.Set(ctx, key, value, ttl); err != nil {
		return err
	}
	return c.redisCache.Set(ctx, key, value, ttl)
}

func (c *MultiBackendCache) Get(ctx context.Context, key string) (interface{}, error) {
	if value, err := c.inMemoryCache.Get(ctx, key); err == nil {
		return value, nil
	}
	return c.redisCache.Get(ctx, key)
}

func (c *MultiBackendCache) Delete(ctx context.Context, key string) error {
	if err := c.inMemoryCache.Delete(ctx, key); err != nil {
		return err
	}
	return c.redisCache.Delete(ctx, key)
}
