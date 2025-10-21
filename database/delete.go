package database

func (s *PostgresStore) DeletePost(id int) error {
	query := `DELETE FROM posts WHERE id = $1`
	cmdTag, err := s.db.Exec(context.Background(), query, id)

	if err != nil {
		return fmt.Errorf("error executing delete query: %w", err)
	}
	if cmdTag.RowsAffected() == 0 { 
		return pgx.ErrNoRows
	}
	return nil
}