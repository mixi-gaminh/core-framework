package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// IRedisDB - IRedisDB
type IRedisDB interface {
	RedisConstructor()
	Close()

	// Keys - Get all Keys with pattern
	Keys(context.Context, string) []string

	// Set - Set a with to Redis DB
	Set(context.Context, string, interface{}) (string, error)

	// SetExpire - Set a with to Redis with Expire Time
	SetExpire(context.Context, string, interface{}, time.Duration) (string, error)

	// SetNX - Set a with to Redis (when is not exists)
	SetNX(context.Context, string, interface{}, time.Duration) (bool, error)

	// HSet - Put dữ liệu Hash ứng với 1 Key
	HSet(context.Context, string, string, interface{}) error

	// HGet - Get dữ liệu Hash ứng với 1 Key
	HGet(context.Context, string, string) (string, error)

	// HGetAll - Get All dữ liệu Hash ứng với 1 Key
	HGetAll(context.Context, string) (map[string]string, error)

	// HExists - Check Field cua Hash co ton tai khong
	HExists(context.Context, string, string) bool

	// HExistsDB1 - Check Field cua Hash trong DB1 co ton tai khong
	HExistsDB1(context.Context, string, string) bool

	// HLen - Get Number Field of Hash
	HLen(context.Context, string) int

	// HDel - Delete một hoặc nhiều dữ liệu trong Hash với tương ứng
	HDel(context.Context, string, ...string) (int64, error)

	// HKeys - Get danh sach Key
	HKeys(context.Context, string) ([]string, error)

	// HMSet - Put nhiều dữ liệu Hash ứng với 1 Key
	HMSet(context.Context, string, map[string]interface{}) error

	// HMGet - Get nhiều dữ liệu Hash ứng với 1 Key
	HMGet(context.Context, string, ...string) ([]interface{}, error)

	// HScan - Scan a Hash Key
	HScan(context.Context, string, string) ([]string, error)

	// Get - Get dữ liệu ứng với 1 Key
	Get(context.Context, string) (string, error)

	// Exists - Kiểm tra nếu 1 tồn tại trong DB
	Exists(context.Context, string) bool

	// SetNXExpire - Set NX a with to Redis with ExpireTime (Second)
	SetNXExpire(context.Context, string, interface{}, int) (bool, error)

	// Incr - Increase a in Redis
	Incr(context.Context, string) (int64, error)

	// Ping - Ping
	Ping(context.Context) (string, error)

	// Save - Save
	Save(context.Context) (string, error)

	// ReJSONUpdate - JSON Redis Driver Update Data
	ReJSONUpdate(context.Context, string, string, string) (string, error)

	// ReJSONSet - JSON Redis Driver Set Data
	ReJSONSet(context.Context, string, string, string, interface{}) (string, error)

	// ReJSONGetString - JSON Redis Driver Get a Data
	ReJSONGetString(context.Context, string, string) (string, error)

	// ReJSONGet - JSON Redis Driver Get a Data
	ReJSONGet(context.Context, string, string) (interface{}, error)

	// HashGetAll - JSON Redis Driver Get All Data
	HashGetAll(context.Context, string, string, []string, ...string) (interface{}, error)

	// ReJSONGetAll - JSON Redis Driver Get All Data in Bucket
	ReJSONGetAll(context.Context, string, string, []string, ...string) (interface{}, error)

	// ReJSONGetAnyAll - ReJSONGetAnyAll
	ReJSONGetAnyAll(context.Context, []string, string, ...string) (interface{}, error)

	// ReJSONArrLen - JSON Redis Driver Get Array Length
	ReJSONArrLen(context.Context, string, string) (interface{}, error)

	// Delete - Delete a Record
	Delete(context.Context, ...string) error

	// DeleteAll - DeleteAll
	DeleteAll(context.Context, string)

	// LPush - Push a Element to Redis List
	LPush(context.Context, string, interface{}) (int64, error)

	// RPush - Push many Element to Redis List
	RPush(context.Context, string, ...interface{}) (int64, error)

	//LPop - Pop a Element from Redis List
	LPop(context.Context, string) (string, error)

	// LRemove - Remove a Element from Redis List
	LRemove(context.Context, string, interface{}) (int64, error)

	// LRange - Get all Element from Redis List
	LRange(context.Context, string) ([]string, error)

	// LIndex - Get Element's from Index of Redis List
	LIndex(context.Context, string, int64) (string, error)

	// LDel - Delete a Redis List
	LDel(context.Context, ...string) (int64, error)

	// LDelDB1 - Delete a Redis List DB1
	LDelDB1(context.Context, ...string) (int64, error)

	// LKeys - Get all List Keys with pattern
	LKeys(context.Context, string) []string

	// SaveRecordPipeline - SaveRecordPipeline
	SaveRecordPipeline(context.Context, string, string, string, string, []byte) error

	// DeleteRecordsPipeline - DeleteRecordsPipeline
	DeleteRecordsPipeline(context.Context, string, string, string, ...string)

	// ZAdd - Adds one or more members to a sorted set, or updates its score, if it already exists
	ZAdd(context.Context, int, string, []float64, []string) (int64, error)

	// ZKeys - List All keys of Sorted Set
	ZKeys(context.Context, int, string) []string

	// ZRange - Returns a range of members in a sorted set, by index
	ZRange(context.Context, int, string, int64, int64) ([]string, error)

	// ZRangeWithScore - Returns a range of members in a sorted set, by index
	ZRangeWithScore(context.Context, int, string, int64, int64) []redis.Z

	// ZRank - Determines the index of a member in a sorted set
	ZRank(context.Context, int, string, string) (int64, error)

	// ZCount - Counts the members in a sorted set with scores within the given values
	ZCount(context.Context, int, string, string, string) (int64, error)

	// ZIncrBy - Increments the score of a member in a sorted set
	ZIncrBy(context.Context, int, string, string, float64) (float64, error)

	// ZIncr - Increments the score of a member in a sorted set
	ZIncr(context.Context, int, string, string) (float64, error)

	// ZScore - Gets the score associated with the given member in a sorted set
	ZScore(context.Context, int, string, string) (float64, error)

	// ZRem - Removes one or more members from a sorted set
	ZRem(context.Context, int, string, ...string) (int64, error)

	// ZDel - Delete a Sorted Set Key
	ZDel(context.Context, string) error

	// ZCard - Gets the number of members in a sorted set
	ZCard(context.Context, int, string) (int64, error)

	// ZRemRangeByRank - Removes all members in a sorted set within the given indexes
	ZRemRangeByRank(context.Context, string, int64, int64) (int64, error)

	// ZRevRange - Returns a range of members in a sorted set, by index, with scores ordered from high to low
	ZRevRange(context.Context, string, int64, int64) ([]string, error)

	// GetKeysByIndex - GetKeysByIndex
	GetKeysByIndex(context.Context, string, int64, int64) ([]string, error)

	// SetIndex - SetIndex
	SetIndex(context.Context, string, string) error

	// RemoveIndex - RemoveIndex
	RemoveIndex(context.Context, string, ...string) error
}
