package monlog

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	DEBUG = iota
	TEST
	PROD
)

type Logger struct {
	prefix string
	level  int
	out    io.Writer
}

func NewLogger() *Logger {
	return &Logger{
		prefix: "monkey-logger",
	}
}

func (l *Logger) Info(message string) {
	timestamp := time.Now().Format(time.RFC3339)
	log := fmt.Sprintf("%s [%s] %s: %s", l.prefix, timestamp, "DEBUG", message)
	fmt.Fprintln(os.Stdout, log)
}
