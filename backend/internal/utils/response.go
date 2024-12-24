package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
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

func RespondToInvalidBody(w http.ResponseWriter, err error, trans ut.Translator) {
	resp := map[string]any{}
	for _, fieldErr := range err.(validator.ValidationErrors) {
		resp[fieldErr.Field()] = fieldErr.Translate(trans)
	}
	RespondWithJSON(w, 400, resp)
}

func RespondEmpty(w http.ResponseWriter) {
	w.WriteHeader(204)
}
