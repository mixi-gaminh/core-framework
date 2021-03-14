package queue

import (
	logger "github.com/mixi-gaminh/core-framework/logs"
	mongodb "github.com/mixi-gaminh/core-framework/repository/mongodb"
	redisdb "github.com/mixi-gaminh/core-framework/repository/redisdb"

	"github.com/centrifugal/centrifuge-go"
)

// QueueHandler - QueueHandler
// var QueueHandler Queue

// Queue - Queue
type Queue struct {
	CentrifugoWSURL string
}

var mgodb mongodb.Mgo
var db redisdb.Cache
var ctfugo *centrifuge.Client

var dbName string

// QueueConstructor -  QueueConstructor
func (q *Queue) QueueConstructor(_centrifugoURL, _dbName string) {
	q.CentrifugoWSURL = _centrifugoURL
	dbName = _dbName
	logger.Constructor(logger.IsDevelopment)
	logger.NewLogger()
	logger.INFO("Queue Constructor Successfull")
}

// TryLenArray - TryLenArray
func TryLenArray(data []string, length int) bool {
	if len(data) < length {
		logger.INFO("Data: ", data, " has length ", len(data))
		return false
	}
	return true
}
