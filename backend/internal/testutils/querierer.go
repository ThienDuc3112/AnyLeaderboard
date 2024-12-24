package testutils

import "anylbapi/internal/database"

var _ database.Querierer = (*MockedQueries)(nil)
