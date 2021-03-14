package logs

import (
	"io"
	"log"
	"os"
)

// LogStashClient - LogStashClient
type LogStashClient struct{}

const ERROR_LOG_FILE = "error.log"
const INFO_LOG_FILE = "info.log"
const WARNING_LOG_FILE = "warning.log"

// IsDevelopment - IsDevelopment
var IsDevelopment bool = true

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)
var err error
var errorLog, infoLog, warningLog *os.File

func Constructor(_isDevelopment bool) {
	IsDevelopment = _isDevelopment
	infoLog, err = os.OpenFile(INFO_LOG_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	warningLog, err = os.OpenFile(WARNING_LOG_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	errorLog, err = os.OpenFile(ERROR_LOG_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err)
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
