package constants

import "time"

const (
	EnvKeyEnvironment = "ENVIRONMENT"
	EnvKeyDbUrl       = "DB_URL"
	EnvKeyPort        = "PORT"
	EnvKeySecret      = "SECRET"
	EnvKeyFrontendUrl = "FRONTEND_URL"

	AccessTokenDuration  = 30 * time.Minute
	RefreshTokenDuration = 14 * 24 * time.Hour
)

const (
	PathValueLeaderboardId = "lid"
	PathValueEntryId       = "eid"
	PathValueUserId        = "uid"
)

const (
	EntryFieldPrefix      = "field"
	EntryDisplayNameField = "\"displayName"
)
