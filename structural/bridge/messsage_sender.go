package bridge

import "fmt"

// MessageSender defines the bridge interface
type MessageSender interface {
	SendMessage(content string)
}

// EmailSender (Concrete Implementation)
type EmailSender struct{}

func (e *EmailSender) SendMessage(content string) {
	fmt.Println("Sending Email:", content)
}

// SMSSender (Concrete Implementation)
type SMSSender struct{}

func (s *SMSSender) SendMessage(content string) {
	fmt.Println("Sending SMS:", content)
}
