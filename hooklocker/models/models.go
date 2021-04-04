package models

// SubscriptionRequest object representation of hook (un)subscription
type SubscriptionRequest struct {
	ID           string
	ChannelID    string
	LeaseSeconds int
	Topic        string
	Secret       string
	Active       bool
}
