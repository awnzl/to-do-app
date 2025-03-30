package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/awnzl/to-do-app/db"
	"github.com/awnzl/to-do-app/internal/api"
	"github.com/awnzl/to-do-app/internal/repository/postgres"
	"github.com/awnzl/to-do-app/internal/service"
)

func main() {
	cfg, err := getDBConfig()
	if err != nil {
		log.Fatalln("get db config", err)
	}

	connectedDB, err := db.NewConnection(cfg)
	if err != nil {
		log.Fatalln("connect to the db", err)
	}

	if err := db.MigrateWithLock(context.Background(), connectedDB, cfg.MigrateURL(), "./migrations"); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	router := setupAPI(connectedDB)

	log.Printf("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func setupAPI(db *sqlx.DB) chi.Router {
	// Initialize repository
	repo := postgres.NewTodoRepo(db)

	// Initialize transaction manager
	txManager := postgres.NewTxManager(db)

	// Initialize service
	todoService := service.NewTodoService(repo, txManager)

	// Create router
	return api.NewRouter(todoService)
}

func getDBConfig() (db.Config, error) {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return db.Config{}, err
	}
	return db.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  "disable",
	}, nil
}
