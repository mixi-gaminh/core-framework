package logs

import "os"

func INFO(v ...interface{}) {
	InfoLogger.Println(v...)
}

func WARNING(v ...interface{}) {
	WarningLogger.Println(v...)
}

func ERROR(v ...interface{}) {
	ErrorLogger.Println(v...)
}

func FATAL(v ...interface{}) {
	ErrorLogger.Println(v...)
	os.Exit(1)
}

func PANIC(v ...interface{}) {
	ErrorLogger.Println(v...)
	os.Exit(-1)
}
