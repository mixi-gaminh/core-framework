package redis

import (
	"context"
	"time"

	logger "github.com/mixi-gaminh/core-framework/logs"

	"github.com/go-redis/redis/v8"
)

//SaveRecordPipeline - SaveRecordPipeline
func (c *Cache) SaveRecordPipeline(ctx context.Context, userID, deviceID, bucketID, recordID string, bodyBytes []byte) error {
	redisKey := userID + "$" + deviceID + "$" + bucketID + "$" + recordID
	// var args interface{} = ""
	timestamp := []float64{float64(time.Now().Local().Unix())}
	members := []string{redisKey}
	var redisZ []*redis.Z
	for i, m := range members {
		redisz := &redis.Z{
			Score:  timestamp[i],
			Member: m,
		}
		redisZ = append(redisZ, redisz)
	}

	c.ReJSONSet(ctx, redisKey, ".", string(bodyBytes), "")

	pipe1 := redisClientWrite1.TxPipeline()
	pipe1.ZAdd(ctx, "BM"+"$"+userID+"$"+deviceID+"$"+bucketID, redisZ...)
	pipe1.ZAdd(ctx, "all$"+userID, redisZ...)
	_, err1 := pipe1.Exec(ctx)
	if err1 != nil {
		redisJSONWrite0.Del(ctx, redisKey)
		redisJSONWrite0.ZRem(ctx, "BM"+"$"+userID+"$"+deviceID+"$"+bucketID, redisKey)
		redisJSONWrite0.ZRem(ctx, "all$"+userID, redisKey)
	}

	pipe0 := redisJSONWrite0.TxPipeline()
	pipe0.HSet(ctx, "all$"+userID, redisKey, string(bodyBytes))
	pipe0.HSet(ctx, "List$BM$"+userID+"$"+deviceID+"$"+bucketID, recordID, "")
	_, err0 := pipe0.Exec(ctx)
	if err0 != nil {
		// Rollback when Pipeline Failed
		logger.ERROR(err0)
		redisJSONWrite0.ZRem(ctx, "BM"+"$"+userID+"$"+deviceID+"$"+bucketID, redisKey)
		redisJSONWrite0.ZRem(ctx, "all$"+userID, redisKey)
		redisJSONWrite0.Del(ctx, redisKey)
		redisJSONWrite0.HDel(ctx, "all$"+userID, redisKey)
		redisJSONWrite0.HDel(ctx, "List$BM$"+userID+"$"+deviceID+"$"+bucketID, recordID)
		return err0
	}
	return nil
}

// DeleteRecordsPipeline - DeleteRecordsPipeline
func (c *Cache) DeleteRecordsPipeline(ctx context.Context, userID, deviceID, bucketID string, recordID ...string) {
	var listKey []string
	for _, r := range recordID {
		listKey = append(listKey, userID+"$"+deviceID+"$"+bucketID+"$"+r)
	}
	pipe1 := redisClientWrite1.TxPipeline()
	pipe1.ZRem(ctx, "all$"+userID, listKey)
	pipe1.ZRem(ctx, "BM"+"$"+userID+"$"+deviceID+"$"+bucketID, listKey)
	pipe1.Exec(ctx)

	pipe0 := redisJSONWrite0.TXPipeline()
	pipe0.Del(ctx, listKey...)
	pipe0.HDel(ctx, "all$"+userID, listKey...)
	pipe0.HDel(ctx, "List$"+"BM"+"$"+userID+"$"+deviceID+"$"+bucketID, listKey...)
	pipe0.Exec(ctx)
}

// UpdateRecordPipeline - UpdateRecordPipeline
// func (c *Cache) UpdateRecordPipeline(ctx context.Context, userID, deviceID, bucketID, recordID string, newBodyBytes []byte, oldData interface{}, oldDataInHash string) error {
// 	redisKey := userID + "$" + deviceID + "$" + bucketID + "$" + recordID
// 	var args interface{} = ""
// 	timestamp := []float64{float64(time.Now().Local().Unix())}
// 	members := []string{redisKey}
// 	var redisZ []*redis.Z
// 	for i, m := range members {
// 		redisz := &redis.Z{
// 			Score:  timestamp[i],
// 			Member: m,
// 		}
// 		redisZ = append(redisZ, redisz)
// 	}

// 	pipe0 := redisJSONWrite0.TXPipeline()
// 	pipe0.JsonSet(ctx, redisKey, ".", string(newBodyBytes), args)
// 	pipe0.HSet(ctx, "all$"+userID, redisKey, string(newBodyBytes))
// 	_, err0 := pipe0.Exec(ctx)
// 	if err0 != nil {
// 		// Rollback when Pipeline Failed
// 		logger.ERROR(err0)
// 		oldBodyBytes, err := json.Marshal(oldData)
// 		if err != nil {
// 			oldBodyBytes = nil
// 		}
// 		pipe0.JsonSet(ctx, redisKey, ".", string(oldBodyBytes), args)
// 		pipe0.HSet(ctx, "all$"+userID, redisKey, oldDataInHash)
// 		return err0
// 	}
// 	pipe1 := redisClientWrite1.TxPipeline()
// 	pipe1.ZAdd(ctx, "BM"+"$"+userID+"$"+deviceID+"$"+bucketID, redisZ...)
// 	pipe1.ZAdd(ctx, "all$"+userID, redisZ...)
// 	pipe1.Exec(ctx)
// 	return nil
// }
