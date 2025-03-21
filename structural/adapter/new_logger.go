package adapter

import "fmt"

// New Logging System (Uses Levels)
type NewLogger interface {
	Log(level, message string)
}

type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(level, message string) {
	fmt.Printf("[%s] %s\n", level, message)
}
