package monlog

import (
	"fmt"
	"io"
	"os"
	"sync/atomic"
	"time"
)

type Logger struct {
	prefix string
	level  Level
	out    io.Writer
}

func New() *Logger {
	return &Logger{
		prefix: "monkey-logger",
		level:  Debug,
		out:    os.Stdout,
	}
}

func (l *Logger) SetPrefix(prefix string) {
	l.prefix = prefix
}

func (l *Logger) SetLevel(level Level) {
	l.level = level
	atomic.StoreInt32((*int32)(&l.level), int32(level))
}

func (l *Logger) Info(message string) {
	l.Log(message, "INFO", nil)
}

func (l *Logger) Debug(message string) {
	if l.level != Debug {
		return
	}
	l.Log(message, "DEBUG", nil)
}

func (l *Logger) Error(message string, err error) {
	l.Log(message, "ERROR", err)
}

func (l *Logger) Fatal(message string, err error) {
	l.Log(message, "FATAL", err)
	panic(err)
}

func (l *Logger) Log(message, logtype string, err error) {
	timestamp := time.Now().Format(time.RFC3339)
	log := fmt.Sprintf("%s [%s] %s: %s", l.prefix, timestamp, logtype, message)
	if err != nil {
		log += fmt.Sprintf(", error: %s", err)
	}
	fmt.Fprintln(l.out, log)
}
