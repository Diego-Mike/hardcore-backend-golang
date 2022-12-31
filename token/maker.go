package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	// creates a new token for a specific user
	CreateToken(username string, duration time.Duration) (string, error)
	// checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
