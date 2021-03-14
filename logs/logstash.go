package logs

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/spf13/viper"
)

// LogStashClient - LogStashClient
type LogStashClient struct{}

var appName, documentType string
var connection net.Conn
var connectionError error

// CreateConnection - CreateConnection
func (l *LogStashClient) CreateConnection() {
	appName = viper.GetString(`log.appName`)
	documentType = viper.GetString(`log.document_type`)
	if connectionError != nil {
		go l.CreateConnection()
		return
	}
	connection, err = net.Dial("tcp", viper.GetString(`log.domain`))
	if err != nil {
		fmt.Println("connectionError: ", connectionError)
		return
	}
}

// CloseConnection - CloseConnection
func (l *LogStashClient) CloseConnection() {
	if connectionError == nil {
		connection.Close()
	}
}

// LogInfo - LogInfo
func (l *LogStashClient) LogInfo(action string, message interface{}) {
	level := "INFO"
	l.SendLog(action, message, level)
}

// LogError - LogError
func (l *LogStashClient) LogWarning(action string, message interface{}) {
	level := "WARNING"
	l.SendLog(action, message, level)
}

// LogError - LogError
func (l *LogStashClient) LogError(action string, message interface{}) {
	level := "ERROR"
	l.SendLog(action, message, level)
}

// SendLog - SendLog
func (l *LogStashClient) SendLog(action string, message interface{}, level string) {
	defer connection.Close()
	mlog := map[string]interface{}{
		"action":        action,
		"appname":       "Ứng dụng " + appName,
		"message":       fmt.Sprintf("%v", message),
		"level":         level,
		"type":          appName,
		"document_type": documentType,
	}
	jLog, err := json.Marshal(mlog)
	if err != nil {
		log.Println(err)
		return
	}
	connection.Write([]byte(string(jLog) + "\n"))
}
