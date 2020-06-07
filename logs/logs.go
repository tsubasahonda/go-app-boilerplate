package logs

import (
	"fmt"
	"log"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stderr, "", log.Ldata|log.Ltime|log.Lmicroseconds|log.LUTC)
}

// Info logger
func Info(msg string, a ...interface{}) {
	logger.Printf("[INFO] %s\n", fmt.Sprintf(msg, a...))
}

// Warn loger
func Warn(msg string, a ...interface{}) {
	logger.Printf("[WARN] %s\n", fmt.Sprintf(msg, a...))
}

// Error logger
func Error(msg string, a ...interface{}) {
	logger.Printf("[ERROR] %s\n", fmt.Sprintf(msg, a...))
}