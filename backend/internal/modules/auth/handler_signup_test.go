package auth

import (
	"anylbapi/internal/database"
	tu "anylbapi/internal/testutils"
	"anylbapi/internal/utils"
	"database/sql"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSignUpHandler_Success(t *testing.T) {
	t.Parallel()
	m := new(tu.MockedQueries)
	service := newAuthService(m)

	// Mock behaviours
	m.On("GetUserByUsername", mock.Anything, mock.Anything).Return(database.User{}, sql.ErrNoRows)
	m.On("GetUserByEmail", mock.Anything, mock.Anything).Return(database.User{}, sql.ErrNoRows)
	m.On("CreateUser", mock.Anything, mock.Anything).Return(nil)

	// Test inputs
	body := signUpReqBody{
		Username:    "test_user",
		DisplayName: "the greatest tester that ever lived",
		Email:       "test@test.com",
		Password:    "longPasswordIsBetterThanComplicatedOne",
	}

	// Run test
	w, r, err := tu.SetupPostJSONTest("/signup", body)
	if err != nil {
		assert.Failf(t, "Setup failed", "Unable to setup the test: %v", err)
	}
	service.signUpHandler(w, r)

	// ================ Assertion ================
	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusCreated, res.StatusCode)
	resBody, err := io.ReadAll(r.Body)
	if err == nil {
		t.Logf("response body: %s", string(resBody))
	}

	m.AssertExpectations(t)
}

func TestSignUpHandler_MissingFields(t *testing.T) {
	t.Parallel()
	m := new(tu.MockedQueries)
	service := newAuthService(m)

	// Mock behaviours

	// Test inputs
	body := signUpReqBody{}

	// Run test
	w, r, err := tu.SetupPostJSONTest("/signup", body)
	if err != nil {
		assert.Failf(t, "Setup failed", "Unable to setup the test: %v", err)
	}
	service.signUpHandler(w, r)

	// ================ Assertion ================
	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	resBody, err := utils.ExtractBody[map[string]string](res.Body)
	assert.Equal(t, err, nil)
	if err == nil {
		assert.NotEqual(t, resBody["username"], "")
		assert.NotEqual(t, resBody["displayName"], "")
		assert.NotEqual(t, resBody["email"], "")
		assert.NotEqual(t, resBody["password"], "")
	}

	m.AssertExpectations(t)
}

func TestSignUpHandler_DuplicateUsername(t *testing.T) {
	t.Parallel()
	m := new(tu.MockedQueries)
	service := newAuthService(m)

	// Mock behaviours
	m.On("GetUserByUsername", mock.Anything, "test_user").Return(database.User{}, nil) // Username exists

	// Test inputs
	body := signUpReqBody{
		Username:    "test_user",
		DisplayName: "tester",
		Email:       "test@test.com",
		Password:    "password123",
	}

	// Run test
	w, r, err := tu.SetupPostJSONTest("/signup", body)
	assert.NoError(t, err)
	service.signUpHandler(w, r)

	// ================ Assertion ================
	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	resBody, err := utils.ExtractBody[map[string]string](res.Body)
	assert.NoError(t, err)
	assert.Equal(t, "Username is taken", resBody["username"])

	m.AssertExpectations(t)
}

func TestSignUpHandler_DuplicateEmail(t *testing.T) {
	t.Parallel()
	m := new(tu.MockedQueries)
	service := newAuthService(m)

	// Mock behaviours
	m.On("GetUserByUsername", mock.Anything, mock.Anything).Return(database.User{}, sql.ErrNoRows) // Username does not exist
	m.On("GetUserByEmail", mock.Anything, "test@test.com").Return(database.User{}, nil)            // Email exists

	// Test inputs
	body := signUpReqBody{
		Username:    "unique_user",
		DisplayName: "tester",
		Email:       "test@test.com",
		Password:    "password123",
	}

	// Run test
	w, r, err := tu.SetupPostJSONTest("/signup", body)
	assert.NoError(t, err)
	service.signUpHandler(w, r)

	// ================ Assertion ================
	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	resBody, err := utils.ExtractBody[map[string]string](res.Body)
	assert.NoError(t, err)
	assert.Equal(t, "Email is already used", resBody["email"])

	m.AssertExpectations(t)
}

func TestSignUpHandler_DatabaseFailureOnUsername(t *testing.T) {
	t.Parallel()
	m := new(tu.MockedQueries)
	service := newAuthService(m)

	// Mock behaviours
	m.On("GetUserByUsername", mock.Anything, "test_user").Return(database.User{}, sql.ErrConnDone) // Simulate DB connection error

	// Test inputs
	body := signUpReqBody{
		Username:    "test_user",
		DisplayName: "tester",
		Email:       "test@test.com",
		Password:    "password123",
	}

	// Run test
	w, r, err := tu.SetupPostJSONTest("/signup", body)
	assert.NoError(t, err)
	service.signUpHandler(w, r)

	// ================ Assertion ================
	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
}

func TestSignUpHandler_UserCreationFailure(t *testing.T) {
	t.Parallel()
	m := new(tu.MockedQueries)
	service := newAuthService(m)

	// Mock behaviours
	m.On("GetUserByUsername", mock.Anything, mock.Anything).Return(database.User{}, sql.ErrNoRows)
	m.On("GetUserByEmail", mock.Anything, mock.Anything).Return(database.User{}, sql.ErrNoRows)
	m.On("CreateUser", mock.Anything, mock.Anything).Return(assert.AnError) // Simulate creation failure

	// Test inputs
	body := signUpReqBody{
		Username:    "test_user",
		DisplayName: "tester",
		Email:       "test@test.com",
		Password:    "password123",
	}

	// Run test
	w, r, err := tu.SetupPostJSONTest("/signup", body)
	assert.NoError(t, err)
	service.signUpHandler(w, r)

	// ================ Assertion ================
	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
}
