package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(fmt.Sprintf("{\"error\":\"%v\"}", msg)))
}

func RespondWithJSON(w http.ResponseWriter, code int, payload any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	data, err := json.Marshal(&payload)
	if err != nil {
		RespondWithError(w, 500, "Unable to marshal data")
		return
	}
	w.Write(data)
}

func RespondEmpty(w http.ResponseWriter) {
	w.WriteHeader(204)
}
