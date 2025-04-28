package models

import "time"

type UserPreview struct {
	Id          int       `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	Username    string    `json:"username"`
	DisplayName string    `json:"displayName"`
	Description string    `json:"description"`
}
