package logs

import (
	"log"
	"os"
)

var (
	Info  *log.Logger
	Error *log.Logger
)

func init() {
	fileInfo, err := os.OpenFile("Infologs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	fileError, err := os.OpenFile("Errorlogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	Info = log.New(fileInfo, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(fileError, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

}
