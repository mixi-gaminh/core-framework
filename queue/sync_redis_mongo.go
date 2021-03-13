package queue

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	logger "github.com/mixi-gaminh/core-framework/logs"
)

// SynchronizeEvent - SynchronizeEvent
func (q *Queue) SynchronizeEvent(ctx context.Context, memberID string) {
	time.Sleep(20 * time.Second)
	for {
		if isLeader(ctx, memberID) {
			logger.INFO("I'm a LEADER MEMBERSHIP")
			db.Save(ctx)
			synchronizeRedisToMongo(ctx)
		}
		time.Sleep(10 * time.Minute)
	}
}
func synchronizeRedisToMongo(ctx context.Context) {
	logger.INFO("Starting Synchronize Redis to MongoDB...")
	syncDeviceMgntHandler(ctx)
	syncBucketMgntHandler(ctx)
	syncRecordMgntHandler(ctx)
	syncAPNDeviceMgntHandler(ctx)
	logger.INFO("Done Synchronize Redis to MongoDB")
}

func syncRecordMgntHandler(ctx context.Context) {
	keys := db.Keys(ctx, "all$")
	if len(keys) <= 0 {
		return
	}
	for _, k := range keys {
		m, err := db.HGetAll(ctx, k)
		if err != nil {
			logger.INFO("GET ALL RECORD: ", err)
			continue
		}
		for k, v := range m {
			keyRecordInHash := k
			k = strings.ReplaceAll(k, "$", "@")
			k = "BM" + "@" + k
			kArr := strings.Split(k, "@")
			if len(kArr) <= 0 {
				continue
			}
			collection := strings.Join(kArr[:len(kArr)-1], "@")
			record := kArr[len(kArr)-1]
			persistData := make(map[string]interface{})
			err := json.Unmarshal([]byte(v), &persistData)
			if err != nil {
				logger.ERROR(err)
				return
			}
			delete(persistData, "_id")
			mgodb.SaveMongo(dbName, collection, record, persistData)
			mgodb.SaveMongo(dbName, "all@"+kArr[1], keyRecordInHash, persistData)
		}
	}
	logger.INFO("Sync Record Management table Finish")
}

func syncBucketMgntHandler(ctx context.Context) {
	keys := db.Keys(ctx, "BM$")
	if len(keys) <= 0 {
		return
	}

	for _, k := range keys {
		v, err := db.ReJSONGetString(ctx, k, ".")
		if err != nil {
			continue
		}
		record := k
		k = strings.ReplaceAll(k, "$", "@")
		kArr := strings.Split(strings.ReplaceAll(k, "BM@", "DM@"), "@")
		if len(kArr) <= 0 {
			continue
		}
		collection := strings.Join(kArr[:len(kArr)-1], "@")
		persistData := make(map[string]interface{})
		err = json.Unmarshal([]byte(v), &persistData)
		if err != nil {
			logger.ERROR(err)
			return
		}
		mgodb.SaveMongo(dbName, collection, record, persistData)
	}
	logger.INFO("Sync Bucket Management table Finish")
}

func syncDeviceMgntHandler(ctx context.Context) {
	keys := db.Keys(ctx, "all$")
	if len(keys) <= 0 {
		return
	}

	for _, k := range keys {
		kArr := strings.Split(k, "$")
		if len(kArr) < 2 {
			continue
		}
		userID := kArr[1]
		keysDM := db.Keys(ctx, "DM$"+userID)
		if len(keys) <= 0 {
			continue
		}

		for _, kDM := range keysDM {
			v, err := db.ReJSONGetString(ctx, kDM, ".")
			if err != nil {
				continue
			}
			persistData := make(map[string]interface{})
			err = json.Unmarshal([]byte(v), &persistData)
			if err != nil {
				logger.ERROR(err)
				continue
			}
			mgodb.SaveMongo(dbName, userID, kDM, persistData)
		}
	}
	logger.INFO("Sync Device Management table Finish")
}

func syncAPNDeviceMgntHandler(ctx context.Context) {
	keys := db.Keys(ctx, "APN$")
	if len(keys) <= 0 {
		return
	}
	for _, k := range keys {
		v, err := db.ReJSONGetString(ctx, k, ".")
		if err != nil {
			logger.ERROR(err)
			continue
		}
		k = strings.ReplaceAll(k, "$", "@")
		kArr := strings.Split(k, "@")
		if len(kArr) <= 0 {
			continue
		}
		collection := strings.Join(kArr[:len(kArr)-1], "@")
		record := kArr[len(kArr)-1]
		persistData := make(map[string]interface{})
		err = json.Unmarshal([]byte(v), &persistData)
		if err != nil {
			logger.ERROR(err)
			return
		}
		mgodb.SaveMongo(dbName, collection, record, persistData)
	}
	logger.INFO("Sync APN Management table Finish")
}

func isLeader(ctx context.Context, memberMQID string) bool {
	v, err := db.Get(ctx, memberMQID)
	if err != nil {
		logger.ERROR(err)
		return false
	}
	if v == "leader" {
		return true
	}
	return false
}
