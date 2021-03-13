package redis

import (
	"context"

	redis "github.com/go-redis/redis/v8"
)

/*---------------------------------------------------------------------------------------------*/

//ZAdd - Adds one or more members to a sorted set, or updates its score, if it already exists
func (c *Cache) ZAdd(ctx context.Context, db int, key string, scores []float64, members []string) (int64, error) {
	var redisZ []*redis.Z
	for i, m := range members {
		redisz := &redis.Z{
			Score:  scores[i],
			Member: m,
		}
		redisZ = append(redisZ, redisz)
	}
	val, err := redisClientWrite1.ZAdd(ctx, key, redisZ...).Result()
	return val, err
}

//ZKeys - List All keys of Sorted Set
func (c *Cache) ZKeys(ctx context.Context, db int, key string) []string {
	keys := redisClientRead1.Keys(ctx, key+"*").Val()
	return keys
}

//ZRange - Returns a range of members in a sorted set, by index
func (c *Cache) ZRange(ctx context.Context, db int, key string, start, stop int64) ([]string, error) {
	ret, err := redisClientRead1.ZRange(ctx, key, start, stop).Result()
	return ret, err
}

//ZRangeWithScore - Returns a range of members in a sorted set, by index
func (c *Cache) ZRangeWithScore(ctx context.Context, db int, key string, start, stop int64) []redis.Z {
	redisZ, _ := redisClientRead1.ZRangeWithScores(ctx, key, start, stop).Result()
	return redisZ
}

//ZRank - Determines the index of a member in a sorted set
func (c *Cache) ZRank(ctx context.Context, db int, key, member string) (int64, error) {
	ret, err := redisClientRead1.ZRank(ctx, key, member).Result()
	return ret, err
}

//ZCount - Counts the members in a sorted set with scores within the given values
func (c *Cache) ZCount(ctx context.Context, db int, key, min, max string) (int64, error) {
	ret, err := redisClientRead1.ZCount(ctx, key, min, max).Result()
	return ret, err
}

//ZIncrBy - Increments the score of a member in a sorted set
func (c *Cache) ZIncrBy(ctx context.Context, db int, key, member string, incr float64) (float64, error) {
	ret, err := redisClientWrite1.ZIncrBy(ctx, key, incr, member).Result()
	return ret, err
}

//ZIncr - Increments the score of a member in a sorted set
func (c *Cache) ZIncr(ctx context.Context, db int, key, member string) (float64, error) {
	redisZ := &redis.Z{Member: member}
	ret, err := redisClientWrite1.ZIncr(ctx, key, redisZ).Result()
	return ret, err
}

//ZScore - Gets the score associated with the given member in a sorted set
func (c *Cache) ZScore(ctx context.Context, db int, key, member string) (float64, error) {
	ret, err := redisClientRead1.ZScore(ctx, key, member).Result()
	return ret, err
}

//ZRem - Removes one or more members from a sorted set
func (c *Cache) ZRem(ctx context.Context, db int, key string, member ...string) (int64, error) {
	ret, err := redisClientWrite1.ZRem(ctx, key, member).Result()
	return ret, err
}

//ZDel - Delete a Sorted Set Key
func (c *Cache) ZDel(ctx context.Context, key string) error {
	err := redisClientWrite1.Del(ctx, key).Err()
	return err
}

/*---------------------------------------------------------------------------------------------*/

//ZCard - Gets the number of members in a sorted set
func (c *Cache) ZCard(ctx context.Context, db int, key string) (int64, error) {
	ret, err := redisClientRead1.ZCard(ctx, key).Result()
	return ret, err
}

//ZRemRangeByRank - Removes all members in a sorted set within the given indexes
func (c *Cache) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) (int64, error) {
	ret, err := redisClientWrite1.ZRemRangeByRank(ctx, key, start, stop).Result()
	return ret, err
}

//ZRevRange - Returns a range of members in a sorted set, by index, with scores ordered from high to low
func (c *Cache) ZRevRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	ret, err := redisClientRead1.ZRevRange(ctx, key, start, stop).Result()
	return ret, err
}
