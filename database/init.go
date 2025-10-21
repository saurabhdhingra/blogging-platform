package database

import (
	"os"

	"github.com/jackc/pgx/v5"
)

type PostGresStore struct {
	db *pgx.Conn
}

func NewPostGresStore() (*PostgresStore, error) {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == ""  {
		connStr = "postgres://<user>:<password>@localhost:5432/blog_db?sslmode=disable"
		log.Printf("WARNING: Using default DB connection string. Set DATABASE_URL env var for production.")
	}
}