package auth

import (
	"anylbapi/internal/database"
	tu "anylbapi/internal/testutils"
	"anylbapi/internal/utils"
	"database/sql"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

var hash, _ = bcrypt.GenerateFromPassword([]byte("correct_password"), bcrypt.DefaultCost)

func TestLoginHandler_Success(t *testing.T) {
	t.Parallel()
	m := new(tu.MockedQueries)
	service := newAuthService(m)

	// Mock behaviors
	mockUser := database.User{
		ID:       1,
		Username: "test_user",
		Email:    "test@test.com",
		Password: string(hash),
	}
	m.On("GetUserByEmail", mock.Anything, "test@test.com").Return(mockUser, nil)

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

	m.AssertExpectations(t)
}

func TestLoginHandler_MissingFields(t *testing.T) {
	t.Parallel()
	m := new(tu.MockedQueries)
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
}

func TestLoginHandler_UserNotFound(t *testing.T) {
	t.Parallel()
	m := new(tu.MockedQueries)
	service := newAuthService(m)

	// Mock behaviors
	m.On("GetUserByEmail", mock.Anything, "nonexistent@test.com").Return(database.User{}, sql.ErrNoRows)

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
}

func TestLoginHandler_IncorrectPassword(t *testing.T) {
	t.Parallel()
	m := new(tu.MockedQueries)
	service := newAuthService(m)

	// Mock behaviors
	mockUser := database.User{
		ID:       1,
		Username: "test_user",
		Email:    "test@test.com",
		Password: string(hash),
	}
	m.On("GetUserByEmail", mock.Anything, "test@test.com").Return(mockUser, nil)

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
}

func TestLoginHandler_DatabaseFailure(t *testing.T) {
	t.Parallel()
	m := new(tu.MockedQueries)
	service := newAuthService(m)

	// Mock behaviors
	m.On("GetUserByEmail", mock.Anything, "test@test.com").Return(database.User{}, errors.New("database error"))

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
}
