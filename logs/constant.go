package logs

import (
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// LogStashClient - LogStashClient
type LogStashClient struct{}

var (
	ERROR_LOG_FILE   string = "error.log"
	INFO_LOG_FILE    string = "info.log"
	WARNING_LOG_FILE string = "warning.log"
)

// IsDevelopment - IsDevelopment
var IsDevelopment bool = true

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

var err error

//var err error
var errorLog, infoLog, warningLog *os.File

func Constructor(_isDevelopment bool) {
	IsDevelopment = _isDevelopment
	if _, err := os.Stat("LOG/INFO_LOG"); os.IsNotExist(err) {
		os.MkdirAll("LOG/INFO_LOG", 0777)
	}
	if _, err := os.Stat("LOG/ERROR_LOG"); os.IsNotExist(err) {
		os.MkdirAll("LOG/ERROR_LOG", 0777)
	}
	if _, err := os.Stat("LOG/WARNING_LOG"); os.IsNotExist(err) {
		os.MkdirAll("LOG/WARNING_LOG", 0777)
	}

	currentTime := time.Now()
	timeString := currentTime.Format("2006-01-02 15:04:05")
	timeString = strings.ReplaceAll(timeString, " ", "_")

	INFO_LOG_FILE = "LOG/INFO_LOG/" + timeString + "_" + INFO_LOG_FILE
	WARNING_LOG_FILE = "LOG/WARNING_LOG/" + timeString + "_" + WARNING_LOG_FILE
	ERROR_LOG_FILE = "LOG/ERROR_LOG/" + timeString + "_" + ERROR_LOG_FILE

	infoLog, err = os.OpenFile(INFO_LOG_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	warningLog, err = os.OpenFile(WARNING_LOG_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	errorLog, err = os.OpenFile(ERROR_LOG_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func NewLogger() {
	InfoLogger = log.New(infoLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(warningLog, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(errorLog, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	if IsDevelopment {
		InfoLogger.SetOutput(io.MultiWriter(os.Stdout, infoLog))
		WarningLogger.SetOutput(io.MultiWriter(os.Stdout, warningLog))
		ErrorLogger.SetOutput(io.MultiWriter(os.Stdout, errorLog))
	}
	InfoLogger.SetOutput(infoLog)
	WarningLogger.SetOutput(warningLog)
	ErrorLogger.SetOutput(errorLog)
}

func Close() {
	errorLog.Close()
	infoLog.Close()
	warningLog.Close()
}
