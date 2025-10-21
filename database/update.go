package database

func (s *PostGresStore) UpdatePost(id int, req *PostRequest) (*Post, error) { 
	query := `
	UPDATE posts
	SET title = $1, content = $2, category = $3, tags = $4, updated_at = NOW()
	WHERE id = $5
	RETURNING id, title, content, category, tags, created_at, updated_at`

	updatedPost := new(Post)
	err := s.db.QueryRow(
		context.Background(),
		query,
		req.Title,
		req.Content,
		req.Category,
		req.Tags,
		id,
	).Scan(
		&updatedPost.ID,
		&updatedPost.Title,
		&updatedPost.Content,
		&updatedPost.Category,
		&updatedPost.Tags,
		&updatedPost.CreatedAt,
		&updatedPost.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows { 
			return nil, nil
		}
		return nil, fmt.Errorf("error updating post: %w", err)
	}
	return updatedPost, nil
}