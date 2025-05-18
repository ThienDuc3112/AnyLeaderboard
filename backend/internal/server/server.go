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
)

type Server struct {
	db *pgxpool.Pool
}

func NewServer() *http.Server {

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
