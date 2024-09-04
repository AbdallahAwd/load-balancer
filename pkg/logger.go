package pkg

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "", log.LstdFlags)

func Log(format string, v ...interface{}) {
	logger.Printf(format, v...)
}
