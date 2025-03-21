package main

import "github.com/abha1erao/design-patterns-in-go/structural/adapter"

func main() {
	// Using New Logger Directly
	newLogger := &adapter.ConsoleLogger{}
	newLogger.Log("INFO", "This is a new log message")

	// Using Old Logger with Adapter
	oldLogger := &adapter.OldLogger{}
	adapter := adapter.NewLoggerAdapter(oldLogger)
	adapter.Log("WARN", "This is an adapted log message")
}
