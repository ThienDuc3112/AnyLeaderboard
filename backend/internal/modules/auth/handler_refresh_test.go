package auth

import (
	"anylbapi/internal/constants"
	"anylbapi/internal/database"
	tu "anylbapi/internal/testutils"
	"anylbapi/internal/utils"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRefreshHandler_Success(t *testing.T) {
	t.Parallel()
	m := tu.NewMockQuerierer(t)
	service := newAuthService(m)

	// Mock behaviors
	mockUser := database.User{
		ID:       1,
		Username: "test_user",
		Email:    "test@test.com",
	}
	mockRefreshToken := database.RefreshToken{
		ID:              1,
		UserID:          1,
		RotationCounter: 1,
		IssuedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		ExpiresAt: pgtype.Timestamp{
			Time:  time.Now().Add(14 * 24 * time.Hour),
			Valid: true,
		},
		DeviceInfo: "TestDevice",
		IpAddress:  "127.0.0.1",
		RevokedAt: pgtype.Timestamp{
			Valid: false,
		},
	}
	mockUpdatedRefreshToken := mockRefreshToken
	mockUpdatedRefreshToken.RotationCounter += 1
	mockTokenStr, err := utils.MakeRefreshTokenJWT(mockRefreshToken, os.Getenv(constants.EnvKeySecret), mockRefreshToken.ExpiresAt.Time)
	if !assert.NoError(t, err, "Cannot create mock refresh token string") {
		t.FailNow()
	}

	m.EXPECT().GetRefreshToken(mock.Anything, mockRefreshToken.ID).Return(mockRefreshToken, nil)
	m.EXPECT().GetUserByID(mock.Anything, mockUser.ID).Return(mockUser, nil)
	m.EXPECT().UpdateRefreshToken(mock.Anything, mock.Anything).Return(mockUpdatedRefreshToken, nil)

	// Test inputs
	cookie := &http.Cookie{
		Name:  "refresh_token",
		Value: mockTokenStr,
	}
	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, "/refresh", nil)
	r.AddCookie(cookie)
	r.RemoteAddr = "127.0.0.1"
	r.Header.Set("User-Agent", "TestDevice")

	// Run test
	service.refreshHandler(w, r)

	// Assertions
	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)
	resBody, err := utils.ExtractBody[map[string]string](res.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, resBody["access_token"])

	cookies := res.Cookies()
	assert.NotEmpty(t, cookies)
	assert.Equal(t, "refresh_token", cookies[0].Name)
	assert.NotEmpty(t, cookies[0].Value)
}

func TestRefreshHandler_ExpiredToken(t *testing.T) {
	t.Parallel()
	service := newAuthService(nil)

	mockRefreshToken := database.RefreshToken{
		ID:              1,
		UserID:          1,
		RotationCounter: 1,
		IssuedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		ExpiresAt: pgtype.Timestamp{
			Time:  time.Now().Add(-1 * time.Hour),
			Valid: true,
		},
		DeviceInfo: "TestDevice",
		IpAddress:  "127.0.0.1",
		RevokedAt: pgtype.Timestamp{
			Valid: false,
		},
	}
	mockTokenStr, err := utils.MakeRefreshTokenJWT(mockRefreshToken, os.Getenv(constants.EnvKeySecret), mockRefreshToken.ExpiresAt.Time)
	if !assert.NoError(t, err, "Cannot create mock refresh token string") {
		t.FailNow()
	}
	// Test inputs
	cookie := &http.Cookie{
		Name:  "refresh_token",
		Value: mockTokenStr,
	}
	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, "/refresh", nil)
	r.AddCookie(cookie)

	// Run test
	service.refreshHandler(w, r)

	// Assertions
	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
	resBody, err := io.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(resBody), "error")

	cookies := res.Cookies()
	assert.NotEmpty(t, cookies)
	assert.Equal(t, "refresh_token", cookies[0].Name)
	assert.Empty(t, cookies[0].Value)
}

func TestRefreshHandler_InvalidToken(t *testing.T) {
	t.Parallel()
	service := newAuthService(nil)

	// Test inputs
	cookie := &http.Cookie{
		Name:  "refresh_token",
		Value: "adsjlfsasdf.asdfasdfasdf.asdfasdfasdf",
	}
	w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, "/refresh", nil)
	r.AddCookie(cookie)

	// Run test
	service.refreshHandler(w, r)

	// Assertions
	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusUnauthorized, res.StatusCode)
	resBody, err := io.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(resBody), "error")

	cookies := res.Cookies()
	assert.NotEmpty(t, cookies)
	assert.Equal(t, "refresh_token", cookies[0].Name)
	assert.Empty(t, cookies[0].Value)
}
