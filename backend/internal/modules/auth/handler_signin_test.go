package auth

import (
	"anylbapi/internal/database"
	tu "anylbapi/internal/testutils"
	"anylbapi/internal/utils"
	"net/http"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

var hash, _ = bcrypt.GenerateFromPassword([]byte("correct_password"), bcrypt.DefaultCost)

func TestLoginHandler_Success_Email_login(t *testing.T) {
	t.Parallel()
	m := tu.NewMockQuerierer(t)
	service := newAuthService(m)

	// Mock behaviors
	mockUser := database.User{
		ID:       1,
		Username: "test_user",
		Email:    "test@test.com",
		Password: string(hash),
	}
	mockRefreshToken := database.RefreshToken{
		ID:              1,
		UserID:          1,
		RotationCounter: 0,
		IssuedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		ExpiresAt: pgtype.Timestamp{
			Time:  time.Now().Add(14 * 24 * time.Hour),
			Valid: true,
		},
		DeviceInfo: "",
		IpAddress:  "",
		RevokedAt: pgtype.Timestamp{
			Valid: false,
		},
	}
	m.EXPECT().GetUserByEmail(mock.Anything, "test@test.com").Return(mockUser, nil)
	m.EXPECT().CreateNewRefreshToken(mock.Anything, mock.Anything).Return(mockRefreshToken, nil)

	// Test inputs
	body := loginReqBody{
		Username: "test@test.com",
		Password: "correct_password",
	}

	// Run test
	w, r, err := tu.SetupPostJSONTest("/login", body)
	assert.NoError(t, err)
	service.loginHandler(w, r)

	// Assertion
	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)
	resBody, err := utils.ExtractBody[map[string]string](res.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, resBody["access_token"])
	if assert.NotEmpty(t, res.Cookies()) {
		cookie := res.Cookies()[len(res.Cookies())-1]
		assert.Equal(t, "refresh_token", cookie.Name)
		assert.Equal(t, true, cookie.HttpOnly)
		assert.Equal(t, true, cookie.Secure)
		assert.WithinDuration(t, mockRefreshToken.ExpiresAt.Time, cookie.Expires, time.Second)
		assert.NotEmpty(t, cookie.Value)
	}

	m.AssertExpectations(t)
}

func TestLoginHandler_Success_Username_login(t *testing.T) {
	t.Parallel()
	m := tu.NewMockQuerierer(t)
	service := newAuthService(m)

	// Mock behaviors
	mockUser := database.User{
		ID:       1,
		Username: "test_user",
		Email:    "test@test.com",
		Password: string(hash),
	}
	mockRefreshToken := database.RefreshToken{
		ID:              1,
		UserID:          1,
		RotationCounter: 0,
		IssuedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		ExpiresAt: pgtype.Timestamp{
			Time:  time.Now().Add(14 * 24 * time.Hour),
			Valid: true,
		},
		DeviceInfo: "",
		IpAddress:  "",
		RevokedAt: pgtype.Timestamp{
			Valid: false,
		},
	}
	m.EXPECT().GetUserByUsername(mock.Anything, "test_user").Return(mockUser, nil)
	m.EXPECT().CreateNewRefreshToken(mock.Anything, mock.Anything).Return(mockRefreshToken, nil)

	// Test inputs
	body := loginReqBody{
		Username: "test_user",
		Password: "correct_password",
	}

	// Run test
	w, r, err := tu.SetupPostJSONTest("/login", body)
	assert.NoError(t, err)
	service.loginHandler(w, r)

	// Assertion
	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)
	resBody, err := utils.ExtractBody[map[string]string](res.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, resBody["access_token"])
	if assert.NotEmpty(t, res.Cookies()) {
		cookie := res.Cookies()[len(res.Cookies())-1]
		assert.Equal(t, "refresh_token", cookie.Name)
		assert.Equal(t, true, cookie.HttpOnly)
		assert.Equal(t, true, cookie.Secure)
		assert.WithinDuration(t, mockRefreshToken.ExpiresAt.Time, cookie.Expires, time.Second)
		assert.NotEmpty(t, cookie.Value)
	}

	m.AssertExpectations(t)
}

func TestLoginHandler_MissingFields(t *testing.T) {
	t.Parallel()
	m := tu.NewMockQuerierer(t)
	service := newAuthService(m)

	// Test inputs
	body := loginReqBody{}

	// Run test
	w, r, err := tu.SetupPostJSONTest("/login", body)
	assert.NoError(t, err)
	service.loginHandler(w, r)

	// Assertion
	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	if assert.NotEmpty(t, res.Cookies()) {
		cookie := res.Cookies()[len(res.Cookies())-1]
		assert.Equal(t, "refresh_token", cookie.Name)
		assert.Equal(t, true, cookie.HttpOnly)
		assert.Equal(t, true, cookie.Secure)
		assert.Zero(t, cookie.Value)
	}
}

func TestLoginHandler_UserNotFound(t *testing.T) {
	t.Parallel()
	m := tu.NewMockQuerierer(t)
	service := newAuthService(m)

	// Mock behaviors
	m.EXPECT().GetUserByEmail(mock.Anything, "nonexistent@test.com").Return(database.User{}, pgx.ErrNoRows)

	// Test inputs
	body := loginReqBody{
		Username: "nonexistent@test.com",
		Password: "any_password",
	}

	// Run test
	w, r, err := tu.SetupPostJSONTest("/login", body)
	assert.NoError(t, err)
	service.loginHandler(w, r)

	// Assertion
	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
	if assert.NotEmpty(t, res.Cookies()) {
		cookie := res.Cookies()[len(res.Cookies())-1]
		assert.Equal(t, "refresh_token", cookie.Name)
		assert.Zero(t, cookie.Value)
	}
}

func TestLoginHandler_IncorrectPassword(t *testing.T) {
	t.Parallel()
	m := tu.NewMockQuerierer(t)
	service := newAuthService(m)

	// Mock behaviors
	mockUser := database.User{
		ID:       1,
		Username: "test_user",
		Email:    "test@test.com",
		Password: string(hash),
	}
	m.EXPECT().GetUserByEmail(mock.Anything, "test@test.com").Return(mockUser, nil)

	// Test inputs
	body := loginReqBody{
		Username: "test@test.com",
		Password: "wrong_password",
	}

	// Run test
	w, r, err := tu.SetupPostJSONTest("/login", body)
	assert.NoError(t, err)
	service.loginHandler(w, r)

	// Assertion
	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
	if assert.NotEmpty(t, res.Cookies()) {
		cookie := res.Cookies()[len(res.Cookies())-1]
		assert.Equal(t, "refresh_token", cookie.Name)
		assert.Zero(t, cookie.Value)
	}
}

func TestLoginHandler_DatabaseFailure(t *testing.T) {
	t.Parallel()
	m := tu.NewMockQuerierer(t)
	service := newAuthService(m)

	// Mock behaviors
	m.EXPECT().GetUserByEmail(mock.Anything, "test@test.com").Return(database.User{}, assert.AnError)

	// Test inputs
	body := loginReqBody{
		Username: "test@test.com",
		Password: "any_password",
	}

	// Run test
	w, r, err := tu.SetupPostJSONTest("/login", body)
	assert.NoError(t, err)
	service.loginHandler(w, r)

	// Assertion
	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	if assert.NotEmpty(t, res.Cookies()) {
		cookie := res.Cookies()[len(res.Cookies())-1]
		assert.Equal(t, "refresh_token", cookie.Name)
		assert.Zero(t, cookie.Value)
	}
}
