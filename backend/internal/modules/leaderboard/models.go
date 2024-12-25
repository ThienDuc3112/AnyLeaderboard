package leaderboard

import (
	"anylbapi/internal/database"
	"errors"
)

// ============ Request body type ============
type createLeaderboardReqBody struct {
	Name                 string         `json:"name" validate:"required,isLBName"`
	Description          string         `json:"description" validate:"max=256"`
	CoverImageURL        string         `json:"coverImageUrl" validate:"omitempty,http_url"`
	ExternalLinks        []externalLink `json:"externalLinks" validate:"dive"`
	AllowAnonymous       bool           `json:"allowAnonymous"`
	RequiredVerification bool           `json:"requiredVerification"`
	Fields               []field        `json:"fields" validate:"required,unique=Name,dive"`
}

type externalLink struct {
	DisplayValue string `json:"displayValue" validate:"required,max=32"`
	URL          string `json:"url" validate:"required,http_url"`
}

type field struct {
	Name       string   `json:"name" validate:"required,max=32"`
	Required   bool     `json:"required"`
	Hidden     bool     `json:"hidden"`
	FieldOrder int      `json:"fieldOrder" validate:"required"`
	Type       string   `json:"type" validate:"required,oneof=TEXT SHORT_TEXT NUMBER DURATION TIMESTAMP OPTION"`
	Options    []string `json:"options" validate:"required_if=Type OPTION,dive,max=32"`
	ForRank    bool     `json:"forRank" validate:"excluded_if=Type OPTION"`
}

// ============ Service param and return types ============
type createLeaderboardParam struct {
	createLeaderboardReqBody
	User database.User
}

// ============ Service errors ============
var (
	errMultipleForRankField     = errors.New("multiple for rank field")
	errUnableToInsertAllFields  = errors.New("unable to insert all fields")
	errNoForRankField           = errors.New("no for rank field found")
	errUnableToInsertAllOptions = errors.New("unable to insert all options")
	errNoPublicField            = errors.New("no public field exist")
)
