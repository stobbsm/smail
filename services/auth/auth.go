package auth

// Authenticatable must be implemented for each type that can authenticate itself
// Used for login in other services.
type Authenticatable interface {
	Authenticate() error
}
