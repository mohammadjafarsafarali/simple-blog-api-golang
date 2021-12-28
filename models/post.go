package models

import "time"

type Post struct {
	ID          uint       `json:"id"  gorm:"primary_key"`
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	PublishedAt *time.Time `json:"publishedAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
}
