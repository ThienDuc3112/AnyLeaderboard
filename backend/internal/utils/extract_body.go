package utils

import (
	"encoding/json"
	"io"
)

func ExtractBody[T any](body io.ReadCloser) (T, error) {
	var res T
	if err := json.NewDecoder(body).Decode(&res); err != nil {
		return res, err
	}

	return res, nil
}
