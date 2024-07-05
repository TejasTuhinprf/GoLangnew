# GoLangnew
**The Multi-Backend Caching Library** in Go provides a unified caching solution that supports multiple backends, including an in-memory cache with LRU (Least Recently Used) eviction policy and external caching solutions such as Redis. The library offers a simple and consistent API for performing cache operations, ensuring high performance and scalability.

**Features**
-In-memory cache with LRU eviction policy.
-Integration with Redis for external caching.
-Unified API for setting, getting, and deleting cache entries.
-Support for cache invalidation and expiration policies.

**Installation** 
To install the library, you can use go get: go get github.com/TejasTuhinprf/golang/multicache

**Usage**
Initializing the Cache: You can initialize the cache with different backends (i.e. in-memory, Redis) as needed by you. For eg:

package main
import ( "context" "log" "time"
"github.com/Tejasprf/golang/multicache"
)
func main() {
inMemoryCache := multicache.NewLRUCache(10)
redisCache, err := multicache.NewRedisCache("localhost:6379", "", 0)
if err != nil {
    log.Fatal("Error connecting to Redis:", err)
}

multiCache := multicache.NewMultiBackendCache(inMemoryCache, redisCache)
ctx := context.Background()

err = multiCache.Set(ctx, "key1", "value1", time.Minute)
if err != nil {
    log.Fatal("Error setting value:", err)
}

value, err := multiCache.Get(ctx, "key1")
if err != nil {
    log.Fatal("Error getting value:", err)
}
log.Println("Got value:", value)

err = multiCache.Delete(ctx, "key1")
if err != nil {
    log.Fatal("Error deleting value:", err)
}
}

**Methods**
The methods utilized in this cache library include : Set, Get, Delete

Set Stores a value in the cache with a specified time-to-live (TTL). 
usage: 
func (c *LRUCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
func (c *RedisCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error 
func (c *MultiBackendCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error

Parameters utilized in Set: 
ctx context.Context: Context for controlling cancellation or timeouts 
key string: The unique identifier for the cache entry. value interface{}: The value to be stored. 
ttl time.Duration: The duration for which the cache entry should remain valid and usable. 
Returns: An error if something goes wrong while setting the value.

Get Retrieves a value from the cache based on the provided key. 
usage: 
func (c *LRUCache) Get(ctx context.Context, key string) (interface{}, error) 
func (c *RedisCache) Get(ctx context.Context, key string) (interface{}, error) 
func (c *MultiBackendCache) Get(ctx context.Context, key string) (interface{}, error)

Parameters utlized in Get: [functionality Same as above] 
Returns: The cached value and an error if the value is not found or some error occurs.

Delete Removes a value from the cache based on the provided key. 
usage: 
func (c *LRUCache) Delete(ctx context.Context, key string) error 
func (c *RedisCache) Delete(ctx context.Context, key string) error 
func (c *MultiBackendCache) Delete(ctx context.Context, key string) error

Parameters utilized in Delete: [functionality same as above] Returns: An error if something goes wrong while deleting the value.

**Benchmarking results In-Memory Cache Performance:**

Set Operation: 1,262 ns/op Get Operation: 3,431 ns/op Delete Operation: 5,082 ns/op

**Redis Cache Performance:**

Set Operation: 41,498 ns/op Get Operation: 38,066 ns/op Delete Operation: 36,725 ns/op
