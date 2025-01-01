package main

import "github.com/MonkeFix/monkey-logger/monlog"

func main() {
	l := monlog.New()

	l.SetPrefix("monkey-fix")
	l.SetLevel(monlog.Test)

	l.Info("Something noteworthy happened!")
	l.Debug("Useful debugging information.")

	l.SetLevel(monlog.Debug)

	l.Debug("Useful debugging information.")
}
