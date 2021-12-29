package models

import "time"

type Post struct {
	Model
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	PublishedAt *time.Time `json:"publishedAt,omitempty"`
}
