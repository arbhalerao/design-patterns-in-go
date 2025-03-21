package bridge

// RegularNotification (Refined Abstraction)
type RegularNotification struct {
	Notification
}

func NewRegularNotification(sender MessageSender) *RegularNotification {
	return &RegularNotification{Notification{sender}}
}

// UrgentNotification (Refined Abstraction)
type UrgentNotification struct {
	Notification
}

func NewUrgentNotification(sender MessageSender) *UrgentNotification {
	return &UrgentNotification{Notification{sender}}
}
