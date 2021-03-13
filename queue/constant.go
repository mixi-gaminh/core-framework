package queue

import (
	mongodb "github.com/mixi-gaminh/core-framework/repository/mongodb"
	redisdb "github.com/mixi-gaminh/core-framework/repository/redisdb"
	logger "github.com/mixi-gaminh/core-framework/logs"

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

// TryLenArray - TryLenArray
func TryLenArray(data []string, length int) bool {
	if len(data) < length {
		logger.INFO("Data: ", data, " has length ", len(data))
		return false
	}
	return true
}
