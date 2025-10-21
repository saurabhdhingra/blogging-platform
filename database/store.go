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

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}	

	store := &PostGresStore{db: conn}
	if err := store.InitTable(); err != nil {
		return nil, fmt.Errorf("failed to initialize posts table: %w", err)
	}

	return store, nil
}