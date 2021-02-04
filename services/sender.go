package services

import "context"

// Sender interface for sending a message
// Must receive a channel of the Message to send
type Sender interface {
	Send(context.Context, Messenger) error
}
