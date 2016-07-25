package cerrors

import (
	"log"
)

const (
	LOG_Default = 0
	LOG_Verbose = 1
	LOG_Debug   = 2
)

var logLevel = LOG_Default

func InitLog(level int) {
	if level == LOG_Debug {
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	} else {
		log.SetFlags(log.Ldate | log.Ltime)
	}
	logLevel = level
}
func LogLevel() int {
	return logLevel
}
func IsVerbose() bool {
	return logLevel != LOG_Default
}
