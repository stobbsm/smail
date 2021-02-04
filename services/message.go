package services

import "time"

// Message is the basic structure of a message
// Sender's and receivers take this message and dispatch it
type Message struct {
	content  string
	topic    string
	sent     time.Time
	received time.Time
}

// Messenger needs to be implemented to send a message
type Messenger interface {
	Content() string
	Topic() string
	Msg(...Addresser)
}

// Addresser interface allows you to set addresses
// This can be email, or some other unique identifier
type Addresser interface {
	// Address should take a list of addresses, and return a Sender
	Address(...string) Sender
}
