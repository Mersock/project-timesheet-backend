package response

import (
	"github.com/google/uuid"
	"time"
)

// SignUpRes -.
type SignUpRes struct {
	ID int64 `json:"id"`
}

// SignInRes -.
type SignInRes struct {
	SessionID            uuid.UUID `json:"session_id"`
	AccessToken          string    `json:"access_token"`
	AccessTokenExpireAt  time.Time `json:"access_token_expires_at"`
	RefreshToken         string    `json:"refresh_token"`
	RefreshTokenExpireAt time.Time `json:"refresh_token_expires_at"`
}