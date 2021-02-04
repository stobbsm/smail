package smtp

import (
	"context"
	"log"
	"net/smtp"

	"github.com/stobbsm/smail/services"
	"github.com/stobbsm/smail/services/credentials"
)

// AuthMethod represents recognized authentication methods
// for SMTP authentication
type AuthMethod string

// Supported authentication methods
const (
	CramMD5 AuthMethod = `CRAM-MD5`
	Plain   AuthMethod = `PLAIN`
)

// Defaults used when a configuration isn't set
const (
	DefaultPort     int = 587
	DefaultAuthType     = Plain
)

// Client implements what's needed for an SMTP client to exist
type Client struct {
	port     int
	host     string
	authtype AuthMethod
	auth     smtp.Auth
	creds    *credentials.Credentials
}

// NewClient returns a new SMTP client
func NewClient(creds *credentials.Credentials, test bool, opts ...Opt) (*Client, error) {
	var c = &Client{creds: creds}
	for _, opt := range opts {
		opt(c)
	}

	if test {
		log.Println("TODO: Implement testing")
	}

	return c, nil
}

// Authenticate implements the Autenticatable interface
// Creates an SMTP authentication
func (c *Client) Authenticate() error {
	switch c.authtype {
	case CramMD5:
		c.auth = smtp.CRAMMD5Auth(c.creds.GetUser(), c.creds.PlainPass())
	case Plain:
		c.auth = smtp.PlainAuth("", c.creds.GetUser(), c.creds.PlainPass(), c.host)
	default:
		log.Printf("Error: %s, received %v", ErrUnknownAuthMethod.Error(), c.authtype)
		return ErrUnknownAuthMethod
	}
	return nil
}

// Send implements the Sender interface
// Receives a context to halt execution
// Receives a channel to listen for messages on
func (c *Client) Send(ctx context.Context, msg services.Messenger) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	}
}
