package mock

import "anylbapi/internal/database"

var _ database.Querierer = (*MockedQueries)(nil)
