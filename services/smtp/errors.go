package smtp

import "errors"

// ErrUnknownAuthMethod error, for when an authentication method isn't understood
var ErrUnknownAuthMethod = errors.New(`Unknown Authentication method`)
