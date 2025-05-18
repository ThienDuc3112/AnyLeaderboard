package main

import (
	"anylbapi/internal/server"
	"fmt"
	"net/http"
)

func main() {
	server := server.NewServer()

	certFile, keyFile := "/app/certs/localserver.crt", "/app/certs/localserver.key"

	err := server.ListenAndServeTLS(certFile, keyFile)
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}
}
