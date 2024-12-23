package helper

import (
	"encoding/json"
	"net/http"
)

func ExtractBody[T any](r *http.Request) (T, error) {
	var res T
	if err := json.NewDecoder(r.Body).Decode(&res); err != nil {
		return res, err
	}

	return res, nil
}
