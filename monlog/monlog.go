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

func NewLogger() *Logger {
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
}

func (l *Logger) Info(message string) {
	timestamp := time.Now().Format(time.RFC3339)
	log := fmt.Sprintf("%s [%s] %s: %s", l.prefix, timestamp, l.level.String(), message)
	fmt.Fprintln(l.out, log)
}
