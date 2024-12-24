package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
