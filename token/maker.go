package token

import "time"

type Maker interface {
	// CreateToken create a new Token for a specific username and duration
	CreateToken(username string, duration time.Duration) (string, *Payload, error)

	//verifyingToken cacks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
