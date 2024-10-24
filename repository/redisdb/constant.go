package redis

import (
	"time"

	redis "github.com/redis/go-redis/v9"
	logger "github.com/mixi-gaminh/core-framework/logs"
	rejonson "github.com/mixi-gaminh/core-framework/repository/redisdb/lib_rejonson"
)

// RedisDB - exported as symbol named "RedisDB"
// var RedisDB Cache

type responseAll struct {
	Data []interface{} `json:"data"`
}

// Cache - Redis Cache Struct
type Cache struct{}

var redisClientRead0, redisClientRead1 *redis.Client
var redisClientWrite0, redisClientWrite1 *redis.Client

var redisJSONRead0 *rejonson.Client
var redisJSONWrite0 *rejonson.Client

// RedisHost - Redis DB URL
var RedisHost string

// RedisConstructor - Create Redis Connection
func (c *Cache) RedisConstructor(_url string, _maxClients, _minIdle int, _password string) {
	RedisHost := _url
	redisClientRead1 = redis.NewClient(&redis.Options{
		Addr:         RedisHost,
		Password:     _password,
		DB:           1,
		PoolSize:     _maxClients,
		PoolTimeout:  30 * time.Second,
		MinIdleConns: _minIdle,
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
		DialTimeout:  1 * time.Minute,
		MaxRetries:   3,
	})
	redisClientWrite1 = redis.NewClient(&redis.Options{
		Addr:         RedisHost,
		Password:     _password,
		DB:           1,
		PoolSize:     _maxClients,
		PoolTimeout:  30 * time.Second,
		MinIdleConns: _minIdle,
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
		DialTimeout:  1 * time.Minute,
		MaxRetries:   3,
	})

	redisClientRead0 = redis.NewClient(&redis.Options{
		Addr:         RedisHost,
		Password:     _password,
		DB:           0,
		PoolSize:     _maxClients,
		PoolTimeout:  30 * time.Second,
		MinIdleConns: _minIdle,
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
		DialTimeout:  1 * time.Minute,
		MaxRetries:   3,
	})
	redisClientWrite0 = redis.NewClient(&redis.Options{
		Addr:         RedisHost,
		Password:     _password,
		DB:           0,
		PoolSize:     _maxClients,
		PoolTimeout:  30 * time.Second,
		MinIdleConns: _minIdle,
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
		DialTimeout:  1 * time.Minute,
		MaxRetries:   3,
	})
	redisJSONRead0 = rejonson.ExtendClient(redisClientRead0)
	redisJSONWrite0 = rejonson.ExtendClient(redisClientWrite0)
	//logger.Constructor(logger.IsDevelopment)
	logger.NewLogger()
	logger.INFO("RedisDB Constructor Successfull")
}

// Close - Close
func (c *Cache) Close() {
	redisClientRead0.Close()
	redisClientWrite0.Close()
	redisClientRead1.Close()
	redisClientWrite1.Close()
	redisJSONRead0.Close()
	redisJSONWrite0.Close()
}
