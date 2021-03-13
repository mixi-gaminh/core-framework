package redis

import (
	"context"
	"time"
)

/* Đánh index cho record để phân trang dữ liệu, tham khảo: https://christophermcdowell.dev/post/pagination-with-redis */

// GetKeysByIndex - GetKeysByIndex
func (c *Cache) GetKeysByIndex(ctx context.Context, key string, start, stop int64) ([]string, error) {
	return c.ZRevRange(ctx, key, start, stop)
}

//SetIndex - SetIndex
func (c *Cache) SetIndex(ctx context.Context, key, member string) error {
	timestamp := []float64{float64(time.Now().Local().Unix())}
	members := []string{member}
	_, err := c.ZAdd(ctx, 1, key, timestamp, members)
	return err
}

// RemoveIndex - RemoveIndex
func (c *Cache) RemoveIndex(ctx context.Context, key string, member ...string) error {
	_, err := c.ZRem(ctx, 1, key, member...)
	return err
}

/***********************************/
