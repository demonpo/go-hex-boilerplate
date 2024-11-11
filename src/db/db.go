package db

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"goHexBoilerplate/ent"
	"log"
	"os"
)

type DB struct {
	DB *ent.Client
}

// new database
func NewDB() *DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	client, err := ent.Open("postgres", conn)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return &DB{
		DB: client,
	}
}
