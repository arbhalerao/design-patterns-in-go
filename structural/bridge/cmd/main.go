package main

import "github.com/arbhalerao/design-patterns-in-go/structural/bridge"

func main() {
	emailSender := &bridge.EmailSender{}
	smsSender := &bridge.SMSSender{}

	// ✅ Create Notifications using the Bridge Pattern
	regularEmail := bridge.NewRegularNotification(emailSender)
	urgentSMS := bridge.NewUrgentNotification(smsSender)

	regularEmail.Send("this is a regular email!")
	urgentSMS.Send("Urgent: Action required immediately!")
}
