package utils

import (
	"net/http"
	"time"
)

func CreateCookie(name, value, domain string, expiresAt time.Time) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Value:    value,
		Expires:  expiresAt,
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Domain:   domain,
	}
}
