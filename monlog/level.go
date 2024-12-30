package monlog

import "fmt"

type Level int8

const (
	Debug Level = iota
	Test
	Prod
)

// String returns an ASCII representation of the log level.
func (l Level) String() string {
	switch l {
	case Debug:
		return "DEBUG"
	case Test:
		return "TEST"
	case Prod:
		return "PROD"
	default:
		return fmt.Sprintf("LEVEL_%d", l)
	}
}
