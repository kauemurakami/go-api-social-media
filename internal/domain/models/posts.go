package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID         uuid.UUID `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	AuthorID   uuid.UUID `json:"author_id"`
	AuthorNick string    `json:"author_nick"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"created_at"`
}
