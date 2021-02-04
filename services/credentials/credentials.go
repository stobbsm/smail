package credentials

// Credentials storage, kept private and outside of other systems
type Credentials struct {
	username string
	password string
}

// NewCredentials returns a new set of credentials to be used
// by consumers
func NewCredentials(user, pass string) *Credentials {
	return &Credentials{user, pass}
}

// GetUser returns the user as a plaintext string
func (creds *Credentials) GetUser() string {
	return creds.username
}

// PlainPass returns the password as plaintext
func (creds *Credentials) PlainPass() string {
	return creds.password
}
