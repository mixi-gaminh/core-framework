package redis

import (
	"context"
	"time"
)

//Keys - Get all Keys with pattern
func (c *Cache) Keys(ctx context.Context, key string) []string {
	keys := redisClientRead0.Keys(ctx, key+"*").Val()
	return keys
}

//Set - Set a value with key to Redis DB
func (c *Cache) Set(ctx context.Context, key string, value interface{}) (string, error) {
	ret, err := redisClientWrite0.Set(ctx, key, value, 0).Result()
	return ret, err
}

//SetExpire - Set a value with key to Redis DB with Expire Time
func (c *Cache) SetExpire(ctx context.Context, key string, value interface{}, t time.Duration) (string, error) {
	ret, err := redisClientWrite0.Set(ctx, key, value, t).Result()
	return ret, err
}

//SetNX - Set a value with key to Redis DB (when key is not exists)
func (c *Cache) SetNX(ctx context.Context, key string, value interface{}, expTime time.Duration) (bool, error) {
	ret, err := redisClientWrite0.SetNX(ctx, key, value, expTime).Result()
	return ret, err
}

//HSet - Put dữ liệu Hash ứng với 1 Key
func (c *Cache) HSet(ctx context.Context, key, field string, value interface{}) error {
	_, err := redisClientWrite0.HSet(ctx, key, field, value).Result()
	return err
}

//HGet - Get dữ liệu Hash ứng với 1 Key
func (c *Cache) HGet(ctx context.Context, key, field string) (string, error) {
	ret, err := redisClientWrite0.HGet(ctx, key, field).Result()
	return ret, err
}

//HGetAll - Get All dữ liệu Hash ứng với 1 Key
func (c *Cache) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	ret, err := redisClientWrite0.HGetAll(ctx, key).Result()
	return ret, err
}

//HExists - Check Field cua Hash co ton tai khong
func (c *Cache) HExists(ctx context.Context, key, field string) bool {
	ret, _ := redisClientWrite0.HExists(ctx, key, field).Result()
	return ret
}

//HExistsDB1 - Check Field cua Hash trong DB1 co ton tai khong
func (c *Cache) HExistsDB1(ctx context.Context, key, field string) bool {
	ret, _ := redisClientWrite1.HExists(ctx, key, field).Result()
	return ret
}

//HLen - Get Number Field of Hash
func (c *Cache) HLen(ctx context.Context, key string) int {
	ret64, _ := redisClientWrite0.HLen(ctx, key).Result()
	ret := int(ret64)
	return ret
}

//HDel - Delete một hoặc nhiều dữ liệu trong Hash với Key tương ứng
func (c *Cache) HDel(ctx context.Context, key string, fields ...string) (int64, error) {
	ret, err := redisClientWrite0.HDel(ctx, key, fields...).Result()
	return ret, err
}

//HKeys - Get danh sach key
func (c *Cache) HKeys(ctx context.Context, key string) ([]string, error) {
	ret, err := redisClientWrite0.HKeys(ctx, key).Result()
	return ret, err
}

//HMSet - Put nhiều dữ liệu Hash ứng với 1 Key
func (c *Cache) HMSet(ctx context.Context, key string, fields map[string]interface{}) error {
	_, err := redisClientWrite0.HMSet(ctx, key, fields).Result()
	return err
}

//HMGet - Get nhiều dữ liệu Hash ứng với 1 Key
func (c *Cache) HMGet(ctx context.Context, key string, fields ...string) ([]interface{}, error) {
	ret, err := redisClientRead0.HMGet(ctx, key, fields...).Result()
	return ret, err
}

//HScan - Scan a Hash Key
func (c *Cache) HScan(ctx context.Context, key string, patternField string) ([]string, error) {
	ret, _, err := redisClientRead0.HScan(ctx, key, 0, patternField+"*", 0).Result()
	return ret, err
}

//Get - Get dữ liệu ứng với 1 Key
func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	ret, err := redisClientRead0.Get(ctx, key).Result()
	return ret, err
}

//Exists - Kiểm tra nếu 1 key tồn tại trong DB
func (c *Cache) Exists(ctx context.Context, key string) bool {
	ret := redisClientRead0.Exists(ctx, key).Val()
	return ret != 0
	// if ret != 0 {
	// 	return true
	// }
	// return false
}

//SetNXExpire - Set NX a value with key to Redis DB with ExpireTime (Second)
func (c *Cache) SetNXExpire(ctx context.Context, key string, value interface{}, expireTime int) (bool, error) {
	ret, err := redisClientWrite0.SetNX(ctx, key, value, time.Duration(expireTime)*time.Second).Result()
	return ret, err
}

// Incr - Increase a value in Redis
func (c *Cache) Incr(ctx context.Context, key string) (int64, error) {
	return redisClientWrite0.Incr(ctx, key).Result()
}

// Ping - Ping
func (c *Cache) Ping(ctx context.Context) (string, error) {
	return redisClientWrite0.Ping(ctx).Result()
}

// Save - Save
func (c *Cache) Save(ctx context.Context) (string, error) {
	return redisClientWrite0.Save(ctx).Result()
}

//HSetExpire - Put dữ liệu Hash ứng với 1 Key và có expire time
func (c *Cache) HSetExpire(ctx context.Context, key string, field string, value interface{}, timeExpire time.Duration) error {
	if _, err := redisClientWrite0.HSet(ctx, key, field, value).Result(); err != nil {
		return err
	}

	if ok, err := redisClientWrite0.Expire(ctx, key, timeExpire).Result(); err != nil || !ok {
		return err
	}
	return nil
}

// Rename - Rename
func (c *Cache) Rename(ctx context.Context, key, newkey string) (string, error) {
	ret, err := redisClientWrite0.Rename(ctx, key, newkey).Result()
	return ret, err
}

// Expire - Cập nhật expire time cho key
func (c *Cache) Expire(ctx context.Context, key string, timeExpire time.Duration) error {
	if ok, err := redisClientWrite0.Expire(ctx, key, timeExpire).Result(); err != nil || !ok {
		return err
	}
	return nil
}
