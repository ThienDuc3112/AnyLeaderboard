package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/models"
	"anylbapi/internal/modules/leaderboard"
	"anylbapi/internal/testutils"
	"anylbapi/internal/utils"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock successful input
var mockOptions = []string{"v1.0", "v1.1", "v1.2", "v2.0"}
var mockBody = models.LeaderboardStructure{
	Name: "Test Leaderboard",
	Fields: []models.Field{
		{Name: "Score", Type: "NUMBER", FieldOrder: 1, ForRank: true, Required: true},
		{Name: "IGT", Type: "DURATION", FieldOrder: 2, Required: true},
		{Name: "Version", Type: "OPTION", FieldOrder: 3, Options: mockOptions, Required: true},
	},
	ExternalLinks: []models.ExternalLink{
		{DisplayValue: "Discord", URL: "https://discord.gg/testServer"},
		{DisplayValue: "Test 2", URL: "https://test.com"},
	},
}

func TestCreateLeaderboardHandler_Success(t *testing.T) {
	t.Parallel()
	m := testutils.NewMockQuerierer(t)
	mc := testutils.NewMockCache(t)
	service := leaderboard.New(m, mc)
	handler := New(&service)

	// Mock behaviors
	m.EXPECT().BeginTx(mock.Anything).Return(m, nil)
	m.EXPECT().Rollback(mock.Anything).Return(nil)
	m.EXPECT().Commit(mock.Anything).Return(nil)
	m.EXPECT().CreateLeaderboard(mock.Anything, mock.Anything).Return(database.Leaderboard{ID: 1}, nil)
	m.EXPECT().CreateLeaderboardExternalLink(mock.Anything, mock.Anything).Return(int64(len(mockBody.ExternalLinks)), nil)
	m.EXPECT().CreateLeadeboardFields(mock.Anything, mock.Anything).Return(int64(len(mockBody.Fields)), nil)
	m.EXPECT().CreateLeadeboardOptions(mock.Anything, mock.Anything).Return(int64(len(mockOptions)), nil)

	mc.EXPECT().Delete(fmt.Sprintf("%s-%d", c.CachePrefixNoLeaderboard, 1))

	// Run test
	w, r, _, err := testutils.SetupPostJSONTestWithUser("/leaderboard", mockBody)
	assert.NoError(t, err)
	handler.createLeaderboardHandler(w, r)

	// Assertion
	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusCreated, res.StatusCode)
	resBody, err := utils.ExtractBody[map[string]interface{}](res.Body)

	assert.NoError(t, err)
	assert.NotEmpty(t, resBody["id"])

	m.AssertExpectations(t)
}

func TestCreateLeaderboardHandler_RequestBodyDecodeError(t *testing.T) {
	t.Parallel()
	m := testutils.NewMockQuerierer(t)
	mc := testutils.NewMockCache(t)
	service := leaderboard.New(m, mc)
	handler := New(&service)

	// Run test
	w, r, _, err := testutils.SetupPostJSONTestWithUser("/leaderboard", "Invalid_json_body")
	assert.NoError(t, err)
	handler.createLeaderboardHandler(w, r)

	// Assertion
	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	resBody, err := utils.ExtractBody[map[string]interface{}](res.Body)

	assert.NoError(t, err)
	assert.NotEmpty(t, resBody["error"])

	m.AssertExpectations(t)
}

func TestCreateLeaderboardHandler_NoFields(t *testing.T) {
	t.Parallel()
	m := testutils.NewMockQuerierer(t)
	mc := testutils.NewMockCache(t)
	service := leaderboard.New(m, mc)
	handler := New(&service)

	// Test inputs (no fields in body)
	body := mockBody
	body.Fields = []models.Field{}
	body.ExternalLinks = []models.ExternalLink{}
	w, r, _, err := testutils.SetupPostJSONTestWithUser("/leaderboard", body)
	assert.NoError(t, err)
	handler.createLeaderboardHandler(w, r)

	// Assertion
	res := w.Result()
	defer res.Body.Close()
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	resBody, err := utils.ExtractBody[map[string]interface{}](res.Body)

	assert.NoError(t, err)
	assert.NotEmpty(t, resBody["fields"])
}

func TestCreateLeaderboardHandler_NoUserContext(t *testing.T) {
	t.Parallel()
	m := testutils.NewMockQuerierer(t)
	mc := testutils.NewMockCache(t)
	service := leaderboard.New(m, mc)
	handler := New(&service)

	// Simulating the absence of a user context
	w, r, err := testutils.SetupPostJSONTest("/leaderboard", mockBody)
	assert.NoError(t, err)
	handler.createLeaderboardHandler(w, r)

	// Assertion
	res := w.Result()
	defer res.Body.Close()
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	resBody, err := utils.ExtractBody[map[string]interface{}](res.Body)

	assert.NoError(t, err)
	assert.NotEmpty(t, resBody["error"])
}

func TestCreateLeaderboardHandler_CreateLeaderboardError(t *testing.T) {
	t.Parallel()

	// Create the table with different test cases
	tests := []struct {
		name           string
		setupMocks     func(m *testutils.MockQuerierer, mc *testutils.MockCache)
		expectedStatus int
	}{
		{
			name: "InsertOptionsError",
			setupMocks: func(m *testutils.MockQuerierer, mc *testutils.MockCache) {
				// Mock behaviors for this test case
				m.EXPECT().BeginTx(mock.Anything).Return(m, nil)
				m.EXPECT().Rollback(mock.Anything).Return(nil)
				m.EXPECT().CreateLeaderboard(mock.Anything, mock.Anything).Return(database.Leaderboard{ID: 1}, nil)
				m.EXPECT().CreateLeaderboardExternalLink(mock.Anything, mock.Anything).Return(int64(len(mockBody.ExternalLinks)), nil)
				m.EXPECT().CreateLeadeboardFields(mock.Anything, mock.Anything).Return(int64(len(mockBody.Fields)), nil)
				m.EXPECT().CreateLeadeboardOptions(mock.Anything, mock.Anything).Return(0, assert.AnError)
			},
		},
		{
			name: "InsertFieldsError",
			setupMocks: func(m *testutils.MockQuerierer, mc *testutils.MockCache) {
				// Mock behaviors for this test case
				m.EXPECT().BeginTx(mock.Anything).Return(m, nil)
				m.EXPECT().Rollback(mock.Anything).Return(nil)
				m.EXPECT().CreateLeaderboard(mock.Anything, mock.Anything).Return(database.Leaderboard{ID: 1}, nil)
				m.EXPECT().CreateLeaderboardExternalLink(mock.Anything, mock.Anything).Return(int64(len(mockBody.ExternalLinks)), nil)
				m.EXPECT().CreateLeadeboardFields(mock.Anything, mock.Anything).Return(0, assert.AnError)
			},
		},
		{
			name: "InsertLinksError",
			setupMocks: func(m *testutils.MockQuerierer, mc *testutils.MockCache) {
				// Mock behaviors for this test case
				m.EXPECT().BeginTx(mock.Anything).Return(m, nil)
				m.EXPECT().Rollback(mock.Anything).Return(nil)
				m.EXPECT().CreateLeaderboard(mock.Anything, mock.Anything).Return(database.Leaderboard{ID: 1}, nil)
				m.EXPECT().CreateLeaderboardExternalLink(mock.Anything, mock.Anything).Return(0, assert.AnError)
			},
		}, {
			name: "CreateLeaderboardError",
			setupMocks: func(m *testutils.MockQuerierer, mc *testutils.MockCache) {
				// Mock behaviors for this test case
				m.EXPECT().BeginTx(mock.Anything).Return(m, nil)
				m.EXPECT().Rollback(mock.Anything).Return(nil)
				m.EXPECT().CreateLeaderboard(mock.Anything, mock.Anything).Return(database.Leaderboard{}, assert.AnError)
			},
		},
		// You can add more cases here with different setupMocks and expectedStatus.
	}

	// Iterate over the test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock Querier and Cache
			m := testutils.NewMockQuerierer(t)
			mc := testutils.NewMockCache(t)
			service := leaderboard.New(m, mc)
			handler := New(&service)

			// Set up mocks for the current test case
			tt.setupMocks(m, mc)

			// Run the test
			w, r, _, err := testutils.SetupPostJSONTestWithUser("/leaderboard", mockBody)
			assert.NoError(t, err)
			handler.createLeaderboardHandler(w, r)

			// Assertion
			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
			resBody, err := utils.ExtractBody[map[string]interface{}](res.Body)
			assert.NoError(t, err)
			assert.NotEmpty(t, resBody["error"])

			m.AssertExpectations(t)
		})
	}
}

func TestCreateLeaderboardHandler_Validation(t *testing.T) {
	t.Parallel()

	// Define the table of test cases
	tests := []struct {
		name               string
		setupBody          func() models.LeaderboardStructure
		expectedStatus     int
		expectedErrorField string
	}{
		{
			name: "MissingName",
			setupBody: func() models.LeaderboardStructure {
				body := copyMockBody()
				body.Name = "" // Missing name
				return body
			},
			expectedStatus:     http.StatusBadRequest,
			expectedErrorField: "name",
		},
		{
			name: "InvalidLBName",
			setupBody: func() models.LeaderboardStructure {
				body := copyMockBody()
				body.Name = "Invalid Name!" // Invalid LBName due to special character "!"
				return body
			},
			expectedStatus:     http.StatusBadRequest,
			expectedErrorField: "name",
		},
		{
			name: "MissingFields",
			setupBody: func() models.LeaderboardStructure {
				body := copyMockBody()
				body.Fields = []models.Field{} // No fields
				return body
			},
			expectedStatus:     http.StatusBadRequest,
			expectedErrorField: "fields",
		},
		{
			name: "InvalidFieldName",
			setupBody: func() models.LeaderboardStructure {
				body := copyMockBody()
				body.Fields[0].Name = "\"Invalid Name!\"" // Invalid field name
				return body
			},
			expectedStatus:     http.StatusBadRequest,
			expectedErrorField: "fields[0].name",
		},
		{
			name: "InvalidExternalLinkURL",
			setupBody: func() models.LeaderboardStructure {
				body := copyMockBody()
				body.ExternalLinks[0].URL = "invalid-url" // Invalid URL
				return body
			},
			expectedStatus:     http.StatusBadRequest,
			expectedErrorField: "externalLinks[0].url",
		},
	}

	// Iterate over the test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup the body based on the test case
			body := tt.setupBody()

			// Mock Querier and Cache
			m := testutils.NewMockQuerierer(t)
			mc := testutils.NewMockCache(t)
			service := leaderboard.New(m, mc)
			handler := New(&service)

			// Run the test
			w, r, _, err := testutils.SetupPostJSONTestWithUser("/leaderboard", body)
			assert.NoError(t, err)
			handler.createLeaderboardHandler(w, r)

			// Assertion
			res := w.Result()
			defer res.Body.Close()

			assert.Equal(t, tt.expectedStatus, res.StatusCode)

			// Check if the response contains the expected error message
			resBody, err := utils.ExtractBody[map[string]interface{}](res.Body)
			assert.NoError(t, err)

			assert.NotEmpty(t, resBody[tt.expectedErrorField])
		})
	}
}

func copyMockBody() models.LeaderboardStructure {
	res := mockBody
	res.ExternalLinks = make([]models.ExternalLink, 0)
	res.Fields = make([]models.Field, 0)
	res.ExternalLinks = append(res.ExternalLinks, mockBody.ExternalLinks...)
	for _, field := range mockBody.Fields {
		newField := field
		if len(field.Options) > 0 {
			newField.Options = make([]string, 0)
			newField.Options = append(newField.Options, field.Options...)
		}
		res.Fields = append(res.Fields, newField)
	}

	return res
}
