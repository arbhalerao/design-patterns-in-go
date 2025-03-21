package adapter

import "fmt"

// Old Logging System (Incompatible with New Logger)
type OldLogger struct{}

func (o *OldLogger) WriteLog(message string) {
	fmt.Println("Old Logger: " + message)
}
