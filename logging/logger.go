package logging

import (
	"log"
	"os"
)

var Logger *log.Logger

func Init(service string) {
	Logger = log.New(os.Stdout, "["+service+"] ", log.LstdFlags|log.Lshortfile)
}
