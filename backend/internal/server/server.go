package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

type Server struct {
	db *sql.DB
}

func NewServer() *http.Server {
	isProduction := os.Getenv("ENVIRONMENT") == "PRODUCTION"
	if !isProduction {
		godotenv.Load(".env.local")
	}

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	db, err := sql.Open("pgx", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	newServer := Server{
		db: db,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      newServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Register server on port %d \n", port)

	return server
}
