package smtp

// Opt is a function that receives the Client struct being constructed
// and sets the value.
type Opt func(*Client)

// Opts used to configure the SMTP client

// Host set's the host the client should connect to
func Host(h string) Opt {
	return func(c *Client) {
		c.host = h
	}
}

// Port set's the port to connect to
func Port(p int) Opt {
	return func(c *Client) {
		c.port = p
	}
}

// Auth set's the authentication type
func Auth(a AuthMethod) Opt {
	return func(c *Client) {
		c.authtype = a
	}
}
