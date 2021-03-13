package redis

import "context"

//LPush - Push a Element to Redis List
func (c *Cache) LPush(ctx context.Context, key string, value interface{}) (int64, error) {
	ret, err := redisClientWrite0.LPush(ctx, key, value).Result()
	return ret, err
}

//RPush - Push many Element to Redis List
func (c *Cache) RPush(ctx context.Context, key string, value ...interface{}) (int64, error) {
	ret, err := redisClientWrite0.RPush(ctx, key, value...).Result()
	return ret, err
}

//LPop - Pop a Element from Redis List
func (c *Cache) LPop(ctx context.Context, key string) (string, error) {
	ret, err := redisClientWrite0.LPop(ctx, key).Result()
	return ret, err
}

//LRemove - Remove a Element from Redis List
func (c *Cache) LRemove(ctx context.Context, key string, value interface{}) (int64, error) {
	ret, err := redisClientWrite0.LRem(ctx, key, 1, value).Result()
	return ret, err
}

//LRange - Get all Element from Redis List
func (c *Cache) LRange(ctx context.Context, key string) ([]string, error) {
	ret, err := redisClientRead0.LRange(ctx, key, 0, -1).Result()
	return ret, err
}

//LIndex - Get Element's Value from Index of Redis List
func (c *Cache) LIndex(ctx context.Context, key string, index int64) (string, error) {
	ret, err := redisClientRead0.LIndex(ctx, key, index).Result()
	return ret, err
}

//LDel - Delete a Redis List
func (c *Cache) LDel(ctx context.Context, key ...string) (int64, error) {
	ret, err := redisClientWrite0.Del(ctx, key...).Result()
	return ret, err
}

//LDelDB1 - Delete a Redis List DB1
func (c *Cache) LDelDB1(ctx context.Context, key ...string) (int64, error) {
	ret, err := redisClientWrite1.Del(ctx, key...).Result()
	return ret, err
}

//LKeys - Get all List Keys with pattern
func (c *Cache) LKeys(ctx context.Context, key string) []string {
	ret := redisClientRead0.Keys(ctx, key+"*").Val()
	return ret
}
