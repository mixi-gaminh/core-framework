package queue

import (
	"log"

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

// TryLenArray - TryLenArray
func TryLenArray(data []string, length int) bool {
	if len(data) < length {
		log.Println("Data: ", data, " has length ", len(data))
		return false
	}
	return true
}
