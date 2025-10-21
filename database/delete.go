package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5" // Recommended router for clean REST handling
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5" 

	"blogging-platform/database/init"
)

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