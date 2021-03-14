package logs

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/spf13/viper"
)

var appName, documentType string
var connection net.Conn
var connectionError error

// CreateConnection - CreateConnection
func (l *LogStashClient) CreateConnection() {
	appName = viper.GetString(`log.appName`)
	documentType = viper.GetString(`log.document_type`)

	connection, err = net.Dial("tcp", viper.GetString(`log.domain`))
	if connectionError != nil {
		log.Println("connectionError: ", connectionError)
		os.Exit(-1)
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
