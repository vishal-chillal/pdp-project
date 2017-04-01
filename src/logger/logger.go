package logger
import (
"log"
"os"
"fmt"
)

type LogLevel int
const(
	DEBUG LogLevel = iota+1
	INFO
)

var (
	l        *log.Logger // PRIVATE LOGGER INSTANCE
	logLevel LogLevel    // PRIVATE LOGGING LEVEL
)

type LoggerConfig struct{
	logFile string
	logLevel int
}

func llog(format string, v ...interface{}) {
	l.Output(3, fmt.Sprintf(format, v...))
}


func Init(config *LoggerConfig) bool{	
	//TODO filename is yet to habdle
	fileHandle := os.Stderr
	l = log.New(fileHandle, "", log.Ldate|log.Ltime|log.Lshortfile)
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	log.Println("logger initialised successfully")
	return true
}

func Debug(format string, v ...interface{}) {
	llog("[INFO] "+format, v...)
	//TODO handle logong level
}

func SetLogLevel(level LogLevel) {
	//TODO set propper levels
}

// func main(){
// 	// logConfig := &LoggerConfig{}
// 	// Init(logConfig)
// 	Debug("My name is %s","Vishal")

// }