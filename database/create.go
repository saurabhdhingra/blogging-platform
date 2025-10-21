package database

import (

)

func (s *PostgresStore) CreatePost(post *Post) error {
	query := `
	INSERT INTO posts (title, content, category, tags)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id, created_at, updated_at`

	now := time.Now().UTC()
	post.CreatedAt = now
	post.UpdatedAt = now
	
	return s.db.QueryRow(
		context.Background(),
		query,
		post.Title, 
		post.Content,
		post.Category,
		post.Tags,
	).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)
}
