package token

import "time"

type Maker interface {
	CreateToken(userID int, username string, role string, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
