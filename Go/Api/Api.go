import (
	"net/http"
	"strconv"

	"github.com/TejaTuhinprf/GoLangnew/Go/Multicache"
	"github.com/gin-gonic/gin"
)

var cache = multi_cache.NewMultiCache()

// multi cache
// Endpoint to retrieve a value by key
func GetCacheValue(ctx *gin.Context) {
	k := ctx.Param("key")
	ctx.JSON(http.StatusOK, cache.Get(k))

}

// Endpoint to delete a value by key
func DeleteCacheValue(ctx *gin.Context) {
	//storing the key value
	k := ctx.Param("key")
	//calling the delete method

	cache.Del(k)
}

// Endpoint to set a key-value pair
func SetCacheValue(ctx *gin.Context, length int) {
	//storing the key and value pair
	k := ctx.Param("key")
	v := ctx.Param("value")
	tstr := ctx.Param("time")

	// Convert time string to integer
	t, err := strconv.Atoi(tstr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid time parameter"})
		return
	}
	//calling the set methods and sending the key,value and length
	st := cache.Set(k, v, length, t)
	if st == "empty string" {
		ctx.JSON(400, gin.H{"error": "empty key"})
	} else {
		ctx.JSON(400, gin.H{"success": "key value pair added succesfully"})
	}

}

// Endpoint to print the redis cache contents
func PrintCache(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, cache.Print())
}

// Endpoint to delete entire data
func DeleteAll(ctx *gin.Context) {
	cache.Del_ALL()

}

// redis
// Endpoint to delete entire data
func RedisDeleteAll(ctx *gin.Context) {
	cache.Redis_Del_ALL()
}

func RedisPrint(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, cache.Redis_Print())
}

func RedisGet(ctx *gin.Context) {
	k := ctx.Param("key")
	ctx.JSON(http.StatusOK, cache.Redis_Get(k))
}

func RedisDeleteValue(ctx *gin.Context) {
	//storing the key value
	k := ctx.Param("key")
	//calling the delete method
	cache.Redis_Del(k)
}

func RedisSetValue(ctx *gin.Context, length int) {
	//storing the key and value pair
	k := ctx.Param("key")
	v := ctx.Param("value")
	tstr := ctx.Param("time")
	// Convert time string to integer
	t, err := strconv.Atoi(tstr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid time parameter"})
		return
	}
	//calling the set methods and sending the key,value and length
	st := cache.Redis_Set(k, v, length, t)
	if st == "empty string" {
		ctx.JSON(400, gin.H{"error": "empty key"})
	} else {
		ctx.JSON(400, gin.H{"success": "key value pair added succesfully"})
	}

}

// in memory
// Endpoint to delete entire data
func InmemoryDeleteAll(ctx *gin.Context) {
	cache.Inmemory_Del_ALL()
}

func InmemoryPrint(ctx *gin.Context) {
	s := cache.Inmemory_Print()

	ctx.JSON(http.StatusOK, s)
}

func InmemoryGet(ctx *gin.Context) {
	k := ctx.Param("key")
	ctx.JSON(http.StatusOK, cache.Inmemory_Get(k))
}

func InmemoryDeleteValue(ctx *gin.Context) {
	//storing the key value
	k := ctx.Param("key")
	//calling the delete method
	cache.Inmemory_Del(k)
}

func InmemorySetValue(ctx *gin.Context, length int) {
	//storing the key and value pair
	k := ctx.Param("key")
	v := ctx.Param("value")
	tstr := ctx.Param("time")
	// Convert time string to integer
	t, err := strconv.Atoi(tstr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid time parameter"})
		return
	}
	//calling the set methods and sending the key,value and length
	st := cache.Inmemory_Set(k, v, length, t)
	if st == "empty string" {
		ctx.JSON(400, gin.H{"error": "empty key"})
	} else {
		ctx.JSON(400, gin.H{"success": "key value pair added succesfully"})
	}
}
