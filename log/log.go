package log

import (
	"fmt"
	"log"
)

const (
	//Min the least verbose log level and default level.
	//More or less only errors and important warnings.
	Min = 0
	//Verbose is more verbose usefull for trouble shooting with source code.
	Verbose = 1
	//DebugMsg use for debugging messages flow.
	DebugMsg = 2 //TODO change flag commont
	//Debug full information for debugging.
	Debug = 3
	dept  = 2
)

var logLevel = Min

//InitLog sets the logging level is not thread safe.
func InitLog(level int) {
	if level > Verbose {
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	} else {
		log.SetFlags(log.Ldate | log.Ltime)
	}
	logLevel = level
}

//Level returns the current logging level.
func Level() int {
	return logLevel
}

//Println log text depending on level if level equal to current
//log level or belove the text is logged. Works as fmt.Println.
func Println(level int, v ...interface{}) {
	if level <= logLevel {
		log.Output(dept, fmt.Sprintln(v...))
	}
}

//Print log text depending on level if level equal to current
//log level or belove the text is logged. Works as fmt.Print.
func Print(level int, v ...interface{}) {
	if level <= logLevel {
		log.Output(dept, fmt.Sprint(v...))
	}
}

//Printf log text depending on level if level equal to current
//log level or belove the text is logged. Works as fmt.Printf.
func Printf(level int, format string, v ...interface{}) {
	if level <= logLevel {
		log.Output(dept, fmt.Sprintf(format, v...))
	}
}

//PrintErr log the err to os. standard error.
func PrintErr(err error) {
	if Level() != Debug {
		log.Printf("Error: %v", err)
	} else {
		log.Printf("Error %+v", err)
	}
}

//ErrNo just create the string ErrNo no.
func ErrNo(no int) string {
	return fmt.Sprintf("ErrNo %v. ", no)
}
