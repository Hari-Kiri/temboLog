package temboLog

import (
	"log"
	"os"
)

// Fatal message will print Your provided message as error log then stop the application.
func FatalLogging(message string, error any) {
	fatalLog := log.New(os.Stderr, "[FATAL] ", 1)
	fatalLog.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	fatalLog.Fatalln(message, message)
}

// Error message will print Your provided error message
func ErrorLogging(message string, error any) {
	errorLog := log.New(os.Stderr, "[ERROR] ", 1)
	errorLog.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	errorLog.Println(message, error)
}

// Info message will print Your standard provided message
func InfoLogging(message ...any) {
	infoLog := log.New(os.Stdout, "[INFO]", 0)
	infoLog.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	infoLog.Println(message...)
}
