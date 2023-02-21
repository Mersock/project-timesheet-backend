package token

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var ErrExpiredToken = errors.New("token expire")
var ErrInvalidToken = errors.New("invalid token")

type Payload struct {
	ID       uuid.UUID `json:"id"`
	UserID   int       `json:"user_id"`
	Username string    `json:"username"`
	Role     string    `json:"role"`
	IssuedAt time.Time `json:"issued_at"`
	ExpireAt time.Time `json:"expired_at"`
}

// NewPayload -.
func NewPayload(userID int, username string, role string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:       tokenID,
		UserID:   userID,
		Username: username,
		Role:     role,
		IssuedAt: time.Now(),
		ExpireAt: time.Now().Add(duration),
	}
	return payload, nil
}

// Valid -.
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpireAt) {
		return ErrExpiredToken
	}

	return nil
}
