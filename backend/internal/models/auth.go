package models

import (
	"time"
)

type RefreshToken struct {
	ID              int32
	UserID          int32
	RotationCounter int32
	IssuedAt        time.Time
	ExpiresAt       time.Time
	DeviceInfo      string
	IpAddress       string
	RevokedAt       *time.Time
}

var r RefreshToken
