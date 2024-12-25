package server

import (
	"anylbapi/internal/constants"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

type Server struct {
	// db *sql.DB
	db *pgxpool.Pool
}

func NewServer() *http.Server {
	isProduction := os.Getenv(constants.EnvKeyEnvironment) == "PRODUCTION"
	if !isProduction {
		godotenv.Load(".env.local")
	}

	port, _ := strconv.Atoi(os.Getenv(constants.EnvKeyPort))
	db, err := pgxpool.New(context.Background(), os.Getenv(constants.EnvKeyDbUrl))
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

	log.Printf("Register server on port %d\n", port)

	return server
}
