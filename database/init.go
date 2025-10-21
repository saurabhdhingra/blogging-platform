package database

import (

)

func (s *PostGresStore) InitTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS posts (
		id SERIAL PRIMARY KEY,
		title VARCHAR(50) NOT NULL,
		content TEXT NOT NULL,
		category VARCHAR(50) NOT NULL,
		tags TEXT[] NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE CURRENT_TIMESTAMP
	)`

	_, err := s.db.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("error creating posts table : %w", err)
	}
	log.Println("PostgreSQL table 'posts' initialized successfully.")
	return nil
}