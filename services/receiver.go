package services

import "context"

// Receiver interface to get a Message
// Returns a channel of messages to receive
type Receiver interface {
	Receive(context.Context) (chan<- Message, error)
}
