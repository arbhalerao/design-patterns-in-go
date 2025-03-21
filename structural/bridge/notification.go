package bridge

// Notification defines the abstraction
type Notification struct {
	sender MessageSender // Bridge to Implementation
}

func (n *Notification) Send(content string) {
	n.sender.SendMessage(content)
}
