package models

import (
	"encoding/json"
	"time"
)

type Leaderboard struct {
	ID                   int
	Name                 string
	Description          string
	Creator              int
	CoverImageUrl        string
	CreatedAt            time.Time
	UpdatedAt            time.Time
	AllowAnonymous       bool
	RequiredVerification bool
	UniqueSubmission     bool
}

type LeaderboardFull struct {
	ID                   int            `json:"id"`
	Name                 string         `json:"name"`
	Description          string         `json:"description"`
	Creator              string         `json:"creator"`
	CoverImageUrl        string         `json:"coverImageUrl,omitempty"`
	CreatedAt            time.Time      `json:"createdAt"`
	EntriesCount         int            `json:"entriesCount"`
	AllowAnonymous       bool           `json:"allowAnonymous"`
	RequiredVerification bool           `json:"requiredVerification"`
	UniqueSubmission     bool           `json:"uniqueSubmission"`
	ExternalLink         []ExternalLink `json:"externalLinks"`
	Fields               []Field        `json:"fields"`
	Data                 []Entry        `json:"data,omitempty"`
}

type SubmitLBStructure struct {
	Name                 string         `json:"name" validate:"required,isLBName"`
	Description          string         `json:"description" validate:"max=256"`
	CoverImageURL        string         `json:"coverImageUrl" validate:"omitempty,http_url"`
	ExternalLinks        []ExternalLink `json:"externalLinks" validate:"max=5,unique=DisplayValue,dive"`
	AllowAnonymous       bool           `json:"allowAnonymous" validate:"excluded_if=UniqueSubmission true"`
	RequiredVerification bool           `json:"requiredVerification"`
	UniqueSubmission     bool           `json:"uniqueSubmission"`
	Fields               []Field        `json:"fields" validate:"required,min=1,max=10,unique=Name,unique=FieldOrder,dive"`
}

type LeaderboardPreview struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	CoverImageUrl  string    `json:"coverImageUrl,omitempty"`
	EntriesCount   int       `json:"entriesCount"`
	CreatedAt      time.Time `json:"createdAt"`
	AllowAnonymous bool      `json:"allowAnonymous"`
	// CreatorId      int       `json:"-"`
	// Rank           float32   `json:"-"` // Only for searching purpose
}

type Entry struct {
	Id          int             `json:"id"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
	Username    string          `json:"username"`
	Fields      json.RawMessage `json:"fields"`
	Verified    bool            `json:"verified"`
	VerifiedAt  *time.Time      `json:"verifiedAt,omitempty"`
	VerifiedBy  string          `json:"verifiedBy,omitempty"`
	SortedField float64         `json:"-"`
}

type ExternalLink struct {
	DisplayValue string `json:"displayValue" validate:"required,max=32"`
	URL          string `json:"url" validate:"required,http_url"`
}

type Field struct {
	Name       string   `json:"name" validate:"required,max=32,isSafeName"`
	Required   bool     `json:"required"`
	Hidden     bool     `json:"hidden"`
	FieldOrder int      `json:"fieldOrder" validate:"required"`
	Type       string   `json:"type" validate:"required,oneof=TEXT NUMBER DURATION TIMESTAMP OPTION"`
	Options    []string `json:"options" validate:"required_if=Type OPTION,omitempty,unique,min=1,dive,min=1,max=32,isSafeName"`
	ForRank    bool     `json:"forRank" validate:"excluded_if=Type OPTION"`
}
