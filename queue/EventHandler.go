package queue

import (
	"encoding/json"
	"log"

	"github.com/centrifugal/centrifuge-go"
)

var dbName string

// QueueConstructor -  QueueConstructor
func (q *Queue) QueueConstructor(_centrifugoURL, _dbName string) {
	q.CentrifugoWSURL = _centrifugoURL
	dbName = _dbName
}

// OnConnect - OnConnect
func (q *Queue) OnConnect(c *centrifuge.Client, e centrifuge.ConnectEvent) {
	log.Printf("Connected with id: %s\n", e.ClientID)
}

// OnDisconnect - OnDisconnect
func (q *Queue) OnDisconnect(c *centrifuge.Client, e centrifuge.DisconnectEvent) {
	log.Printf("Disconnected: %s\n", e.Reason)
	log.Println("Retry Connect...")
	err := ctfugo.Connect()
	if err != nil {
		log.Println(err)
		return
	}
}

// CreateConnectionToCentrifugo -  Ket noi Centrifugo
func (q *Queue) CreateConnectionToCentrifugo() (*centrifuge.Client, error) {
	ctfugo = centrifuge.New(q.CentrifugoWSURL, centrifuge.DefaultConfig())
	ctfgHandler := &Queue{}
	ctfugo.OnConnect(ctfgHandler)
	ctfugo.OnDisconnect(ctfgHandler)

	err := ctfugo.Connect()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ctfugo, nil
}

// Gửi Message tới Centrifugo
func publishMessageToCentrifugo(channel string, dataBytes []byte) error {
	ret, err := ctfugo.Publish(channel, dataBytes)
	if err != nil {
		log.Println("DEBUG ERROR publishMessageToCentrifugo:", err)
		return err
	}
	log.Println("DEBUG INFO publishMessageToCentrifugo:", ret)
	return nil
}

// publishDoneActionEvent - publishDoneActionEvent
func publishDoneActionEvent(userID, appID, bucketID, recordID, action string) {
	// Init prefix message
	prefixMsg := "Done action with record "
	switch action {
	case "SAVE":
		prefixMsg = "Saved record "
	case "UPDATE":
		prefixMsg = "Updated record "
	case "DELETE":
		prefixMsg = "Deleted record "
	}

	// Init publish channel & body message
	channel := userID + "$" + appID + "$" + bucketID
	bodyMap := map[string]interface{}{
		"record_id": recordID,
		"action":    action,
		"messsage":  prefixMsg + recordID,
	}
	bodyMsg, err := json.Marshal(bodyMap)
	if err != nil {
		log.Println(err)
		return
	}

	// Publish message to Centrifugo
	log.Println("Publish Message To Centrifugo:\nChannel: ", channel, "\nMessage: ", bodyMap)
	if err := publishMessageToCentrifugo(channel, bodyMsg); err != nil {
		log.Println(err)
		return
	}
}
