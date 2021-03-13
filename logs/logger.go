package logs

func INFO(v ...interface{}) {
	InfoLogger.Println(v...)
	//log.Println(v...)
}

func WARNING(v ...interface{}) {
	WarningLogger.Println(v...)
	//log.Println(v...)
}

func ERROR(v ...interface{}) {
	ErrorLogger.Println(v...)
	//log.Println(v...)
}
