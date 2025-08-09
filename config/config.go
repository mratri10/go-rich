package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {
	databaseUrl := os.Getenv("DATABASE_URL")

	if databaseUrl == "" {
		databaseUrl = "postgres://richman:glory_100jt@167.99.76.27:5432/richdb?sslmode=disable"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error

	DB, err = pgxpool.New(ctx, databaseUrl)

	if err != nil {
		log.Fatal("DB connection error: ", err)
	}

	log.Println("Connected to PosgreSQL")
}
