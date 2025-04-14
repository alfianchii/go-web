package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	RowID uuid.UUID `json:"row_id"`
	UserID int `json:"user_id"`
	Token string `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
	IPAddress string `json:"ip_address"`
	IsBlacklisted bool `json:"is_blacklisted"`
}