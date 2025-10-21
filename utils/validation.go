package utils

import (
	"strings"

	"blogging-platform/models/postRequest"
)

func validatePostRequest(req PostRequest) []string {
	var errors []string
	if strings.TrimSpace(req.Title) == "" {
		errors = append(errors, "Title is required.")
	}
	if strings.TrimSpace(req.Content) == "" { 
		errors = append(errors, "Content is required.")
	}
	if string.TrimSpace(req.Category) == "" {
		errors = append(errors, "Category is required.")
	}

	return errors
}