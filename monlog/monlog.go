package monlog

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Logger struct {
	prefix string
	level  Level
	out    io.Writer
}

func New(prefix string, level Level, out io.Writer) *Logger {
	return &Logger{
		prefix: prefix,
		level:  level,
		out:    out,
	}
}

func Default() *Logger {
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
	// atomic.StoreInt32((*int32)(&l.level), int32(level))
}

func (l *Logger) Info(message string) {
	m := l.Log(message, "INFO", nil)
	l.Print(m)
}

func (l *Logger) Debug(message string) {
	if l.level != Debug {
		return
	}
	m := l.Log(message, "DEBUG", nil)
	l.Print(m)
}

func (l *Logger) Error(message string, err error) {
	m := l.Log(message, "ERROR", err)
	l.Print(m)
}

func (l *Logger) Fatal(message string, err error) {
	m := l.Log(message, "FATAL", err)
	l.Print(m)
	panic(m)
}

func (l *Logger) Log(message, logtype string, err error) string {
	timestamp := time.Now().Format(time.RFC3339)
	log := fmt.Sprintf("%s [%s] %s: %s", l.prefix, timestamp, logtype, message)
	if err != nil {
		log += fmt.Sprintf(", error: %s", err)
	}
	return log
}

func (l *Logger) Print(message string) {
	fmt.Fprintln(l.out, message)
}
