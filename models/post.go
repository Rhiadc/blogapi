package models

import (
	"errors"
	"time"
)

type Post struct {
	ID          uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `gorm:"size:255; not null;unique" json:"title"`
	Content     string    `gorm:"size:255; not null;" json:"content"`
	UserId      uint32    `gorm:"not null" json:"user_id"`
	User        User      `json:"user"`
	PublishedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"published_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Comments    []Comment `gorm:"ForeignKey:PostId"`
}

var (
	ErrPostEmptyTitle   = errors.New("post.title can't be empty")
	ErrPostEmptyContent = errors.New("post.content can't be empty")
)

func ValidatePost(p Post) error {
	if p.Title == "" {
		return ErrPostEmptyTitle
	} else if p.Content == "" {
		return ErrPostEmptyContent
	}
	return nil
}
