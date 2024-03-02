package domain

import "github.com/google/uuid"

type BlogPost struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Body   string    `json:"body"`
	Author User      `json:"author"`
}
