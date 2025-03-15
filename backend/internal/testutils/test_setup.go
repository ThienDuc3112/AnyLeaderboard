package testutils

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"bytes"
	"context"
	"encoding/json"
	"math"
	"math/rand"
	"net/http"
	"net/http/httptest"
	time "time"

	"github.com/jackc/pgx/v5/pgtype"
)

func SetupPostJSONTest(path string, payload any) (*httptest.ResponseRecorder, *http.Request, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, nil, err
	}

	r := httptest.NewRequest(http.MethodPost, path, bytes.NewReader(body))
	r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	return w, r, nil
}

func SetupPostJSONTestWithUser(path string, payload any) (*httptest.ResponseRecorder, *http.Request, database.User, error) {
	w, r, err := SetupPostJSONTest(path, payload)
	var mockUser database.User
	if err == nil {
		mockDate := time.Date(2024, time.Month(1+rand.Int31n(12)), int(1+rand.Int31n(28)), 0, 0, 0, 0, time.Local)
		mockUser = database.User{
			ID:          rand.Int31n(math.MaxInt32),
			CreatedAt:   pgtype.Timestamptz{Time: mockDate},
			UpdatedAt:   pgtype.Timestamptz{Time: mockDate},
			Username:    "tester",
			DisplayName: "The best tester to ever lived",
			Email:       "tester@professional.com",
			Description: "The best tester to ever lived",
		}
		newCtx := context.WithValue(r.Context(), c.MidKeyUser, mockUser)
		r = r.WithContext(newCtx)
	}
	return w, r, mockUser, err
}
