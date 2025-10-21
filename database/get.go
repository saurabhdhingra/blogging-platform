package database

import (

)

func (s *PostGresStore) GetPostByID(id int) (*Post, error) {
	query := `
	SELECT id, title, content, category, tags, created_at, updated_at
	FROM posts
	WHERE id = $1`

	post := new(Post)
	err := s.db.QueryRow(context.Background(), query, id).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.Category,
		&post.Tags,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error querying post by ID: %w", err)
	}
	return post, nil
}

func (s *PostGresStore) GetAllPosts(searchTerm string) ([]*Post, error) { 
	var (
		posts 	[]*Post
		query	string
		args	[]interface{}
	)

	query = `
	SELECT id, title, content, category, tags, created_at, updated_at
	FROM posts`

	if searchTerm != "" {
		query  += ` WHERE title ILIKE $1 OR content ILIKE $1 OR category ILIKE $1`
		args = append(args, "%" + searchTerm + "%")
	}

	query += ` ORDER BY created_at DESC`

	rows, err := s.db.Query(context.Background(), query, args...)
	if err != nil {
		return nil, fmt.Errorf("error querying all posts: %w", err)
	}
	defer rows.Close()

	for rows.Next(){
		post := new(Post)
		err := rows.Scan(
			&post.ID,
			&post.Title, 
			&post.Content,
			&post.Category,
			&post.Tags,
			&post.CreatedAt,
			&post.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("error scanning post row: %w", err)
		}
		posts = append(posts, post)
	}

	return posts, nil
}