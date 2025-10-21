package handlers


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

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5" 
)

type PostHandler struct {
	store *PostgresStore
}

func NewPostHandler(store *PostgresStore) *PostHandler {
	return &PostHandler(store: store)
}

