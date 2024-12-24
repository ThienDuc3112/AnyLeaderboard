package main

import (
	"anylbapi/internal/server"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	server := server.NewServer()

	certFile, keyFile, err := getTLSFilePaths()
	if err != nil {
		panic(fmt.Sprintf("Cannot get tls file: %s", err))
	}

	err = server.ListenAndServeTLS(certFile, keyFile)
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}
}

func getTLSFilePaths() (certPath string, keyPath string, err error) {
	// Get the ENVIRONMENT variable
	environment := os.Getenv("ENVIRONMENT")

	var projectRoot string

	if environment == "PRODUCTION" {
		// Assume the binary is in /bin/, move up one level to the project root
		executablePath, err := os.Executable()
		if err != nil {
			return "", "", fmt.Errorf("failed to get executable path: %v", err)
		}
		projectRoot = filepath.Join(filepath.Dir(executablePath), "..")
	} else {
		// Assume running from /cmd/api/main.go, move up two levels to the project root
		_, filePath, _, ok := runtime.Caller(0)
		if !ok {
			return "", "", fmt.Errorf("failed to get caller information")
		}
		projectRoot = filepath.Join(filepath.Dir(filePath), "..", "..")
	}

	// Clean up the path
	projectRoot = filepath.Clean(projectRoot)

	// Construct the paths for the certificate and key files
	certPath = filepath.Join(projectRoot, "localserver.crt")
	keyPath = filepath.Join(projectRoot, "localserver.key")

	return certPath, keyPath, nil
}
