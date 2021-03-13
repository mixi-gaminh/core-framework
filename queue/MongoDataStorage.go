package queue

import (
	"context"
	"encoding/json"
	"strings"

	logger "github.com/mixi-gaminh/core-framework/logs"
)

// Save - Save
func (q *Queue) Save(ctx context.Context, msg []string) {
	msgArr := strings.Split(msg[0], ",")
	if !TryLenArray(msgArr, 2) {
		logger.ERROR("MongoDB Consumer Cannot Storing: \nKey: " + msg[0] + "\nValue: " + msg[1])
		return
	}

	headerArr := strings.Split(msg[0], ",")
	if !TryLenArray(headerArr, 2) {
		return
	}

	collection := headerArr[0]
	record := headerArr[1]

	bodyData := msg[1]
	persistData := make(map[string]interface{})

	err := json.Unmarshal([]byte(bodyData), &persistData)
	if err != nil {
		logger.ERROR(err)
		return
	}
	kArr := strings.Split(collection, "@")
	if len(kArr) < 4 {
		return
	}
	keyRecordInHash := kArr[1] + "$" + kArr[2] + "$" + kArr[3] + "$" + record
	err = mgodb.SaveMongoMQ(dbName, collection, record, persistData)
	if err != nil {
		return
	}
	err = mgodb.SaveMongoMQ(dbName, "all@"+kArr[1], keyRecordInHash, persistData)
	if err != nil {
		return
	}

	db.Delete(ctx, keyRecordInHash)
	db.HDel(ctx, "all$"+kArr[1], keyRecordInHash)

	logger.INFO("MongoDB Consumer Stored: \nData: " + bodyData + "\nRecord ID: " + record + "\nCollection: " + collection)
}

// Update - Update
func (q *Queue) Update(ctx context.Context, msg []string) {
	msgArr := strings.Split(msg[0], ",")
	if !TryLenArray(msgArr, 2) {
		logger.ERROR("MongoDB Consumer Cannot Storing: \nKey: " + msg[0] + "\nValue: " + msg[1])
		return
	}

	headerArr := strings.Split(msg[0], ",")
	if !TryLenArray(headerArr, 2) {
		return
	}

	collection := headerArr[0]
	record := headerArr[1]
	bodyData := msg[1]
	persistData := make(map[string]interface{})

	err := json.Unmarshal([]byte(bodyData), &persistData)
	if err != nil {
		logger.ERROR(err)
		return
	}
	kArr := strings.Split(collection, "@")
	if len(kArr) < 4 {
		return
	}

	keyRecordInHash := kArr[1] + "$" + kArr[2] + "$" + kArr[3] + "$" + record

	err = mgodb.SaveMongoMQ(dbName, collection, record, persistData)
	if err != nil {
		return
	}
	err = mgodb.SaveMongoMQ(dbName, "all@"+kArr[1], keyRecordInHash, persistData)
	if err != nil {
		return
	}

	db.Delete(ctx, keyRecordInHash)
	db.HDel(ctx, "all$"+kArr[1], keyRecordInHash)

	logger.INFO("MongoDB Consumer Updated: \nData: " + bodyData + "\nRecord ID: " + record + "\nCollection: " + collection)
	go publishDoneActionEvent(kArr[1], kArr[2], kArr[3], record, "UPDATE")
}

//Delete - Delete
func (q *Queue) Delete(ctx context.Context, msg []string) {
	collection := msg[0]
	record := msg[1]
	kArr := strings.Split(collection, "@")
	if len(kArr) < 4 {
		return
	}
	keyRecordInHash := kArr[1] + "$" + kArr[2] + "$" + kArr[3] + "$" + record

	mgodb.DeleteInMongoMQ(dbName, collection, record)
	logger.INFO("MongoDB Consumer Deleted: \nRecord: " + record + "\nCollection: " + collection)

	mgodb.DeleteInMongoMQ(dbName, "all@"+kArr[1], keyRecordInHash)
	logger.INFO("MongoDB Consumer Deleted: \nRecord: " + keyRecordInHash + "\nCollection: " + "all@" + kArr[1])

	go publishDoneActionEvent(kArr[1], kArr[2], kArr[3], record, "DELETE")
}

//Drop - Drop
func (q *Queue) Drop(ctx context.Context, msg []string) {
	key := msg[0]
	if key == "Drop" {
		collection := msg[1]
		mgodb.DropCollectionInMongoMQ(dbName, collection)
		logger.INFO("MongoDB Consumer Dropped: \nCollection: " + collection)
	} else if key == "DropMany" {
		collections := strings.Split(msg[1], ",")
		var listCollection []string
		listCollection = append(listCollection, collections...)

		mgodb.DropManyCollectionInMongoMQ(dbName, listCollection)
		logger.INFO("MongoDB Consumer Dropped Many: \nCollection: " + msg[1])
	}
}
