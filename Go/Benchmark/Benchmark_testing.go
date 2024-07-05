func BenchmarkRedisCache_Get(b *testing.B) {
    cache, err := NewRedisCache("localhost:6379", "", 0)
    if err != nil {
        b.Fatal("Error connecting to Redis:", err)
    }
    // Prepopulate the cache
    ctx := context.Background()
    for i := 0; i < 1000; i++ {
        key := fmt.Sprintf("key%d", i)
        value := fmt.Sprintf("value%d", i)
        cache.Set(ctx, key, value, time.Minute)
    }
    benchmarkGet(b, cache)
}

func BenchmarkRedisCache_Delete(b *testing.B) {
    cache, err := NewRedisCache("localhost:6379", "", 0)
    if err != nil {
        b.Fatal("Error connecting to Redis:", err)
    }
    // Prepopulate the cache
    ctx := context.Background()
    for i := 0; i < 1000; i++ {
        key := fmt.Sprintf("key%d", i)
        value := fmt.Sprintf("value%d", i)
        cache.Set(ctx, key, value, time.Minute)
    }
    benchmarkDelete(b, cache)
}

func BenchmarkMultiBackendCache_Set(b *testing.B) {
    inMemoryCache := NewLRUCache(1000)
    redisCache, err := NewRedisCache("localhost:6379", "", 0)
    if err != nil {
        b.Fatal("Error connecting to Redis:", err)
    }
    cache := NewMultiBackendCache(inMemoryCache, redisCache)
    benchmarkSet(b, cache)
}

func BenchmarkMultiBackendCache_Get(b *testing.B) {
    inMemoryCache := NewLRUCache(1000)
    redisCache, err := NewRedisCache("localhost:6379", "", 0)
    if err != nil {
        b.Fatal("Error connecting to Redis:", err)
    }
    cache := NewMultiBackendCache(inMemoryCache, redisCache)
    // Prepopulate the cache
    ctx := context.Background()
    for i := 0; i < 1000; i++ {
        key := fmt.Sprintf("key%d", i)
        value := fmt.Sprintf("value%d", i)
        cache.Set(ctx, key, value, time.Minute)
    }
    benchmarkGet(b, cache)
}

func BenchmarkMultiBackendCache_Delete(b *testing.B) {
    inMemoryCache := NewLRUCache(1000)
    redisCache, err := NewRedisCache("localhost:6379", "", 0)
    if err != nil {
        b.Fatal("Error connecting to Redis:", err)
    }
    cache := NewMultiBackendCache(inMemoryCache, redisCache)
    // Prepopulate the cache
    ctx := context.Background()
    for i := 0; i < 1000; i++ {
        key := fmt.Sprintf("key%d", i)
        value := fmt.Sprintf("value%d", i)
        cache.Set(ctx, key, value, time.Minute)
    }
    benchmarkDelete(b, cache)
}
