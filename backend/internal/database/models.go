// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type FieldType string

const (
	FieldTypeTEXT      FieldType = "TEXT"
	FieldTypeNUMBER    FieldType = "NUMBER"
	FieldTypeDURATION  FieldType = "DURATION"
	FieldTypeTIMESTAMP FieldType = "TIMESTAMP"
	FieldTypeOPTION    FieldType = "OPTION"
)

func (e *FieldType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = FieldType(s)
	case string:
		*e = FieldType(s)
	default:
		return fmt.Errorf("unsupported scan type for FieldType: %T", src)
	}
	return nil
}

type NullFieldType struct {
	FieldType FieldType
	Valid     bool // Valid is true if FieldType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullFieldType) Scan(value interface{}) error {
	if value == nil {
		ns.FieldType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.FieldType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullFieldType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.FieldType), nil
}

type Leaderboard struct {
	ID                  int32
	Name                string
	Description         string
	CreatedAt           pgtype.Timestamptz
	UpdatedAt           pgtype.Timestamptz
	CoverImageUrl       pgtype.Text
	AllowAnonymous      bool
	RequireVerification bool
	UniqueSubmission    bool
	Creator             int32
	NameLanguage        string
	DescriptionLanguage string
	NameTsv             interface{}
	DescriptionTsv      interface{}
	SearchTsv           interface{}
}

type LeaderboardEntry struct {
	ID            int32
	CreatedAt     pgtype.Timestamptz
	UpdatedAt     pgtype.Timestamptz
	UserID        pgtype.Int4
	Username      string
	LeaderboardID int32
	SortedField   float64
	CustomFields  []byte
	Verified      bool
	VerifiedAt    pgtype.Timestamp
	VerifiedBy    pgtype.Int4
}

type LeaderboardExternalLink struct {
	ID            int32
	LeaderboardID int32
	DisplayValue  string
	Url           string
}

type LeaderboardFavourite struct {
	UserID        int32
	LeaderboardID int32
}

type LeaderboardField struct {
	Lid        int32
	FieldName  string
	FieldValue FieldType
	FieldOrder int32
	ForRank    bool
	Hidden     bool
	Required   bool
}

type LeaderboardOption struct {
	Lid       int32
	FieldName string
	Option    string
}

type LeaderboardVerifier struct {
	LeaderboardID int32
	UserID        int32
	AddedAt       pgtype.Timestamptz
}

type RefreshToken struct {
	ID              int32
	UserID          int32
	RotationCounter int32
	IssuedAt        pgtype.Timestamptz
	ExpiresAt       pgtype.Timestamptz
	DeviceInfo      string
	IpAddress       string
	RevokedAt       pgtype.Timestamptz
}

type User struct {
	ID          int32
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
	Username    string
	DisplayName string
	Email       string
	Password    string
	Description string
}
