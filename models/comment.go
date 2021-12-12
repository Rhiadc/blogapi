package models

import "errors"

type Comment struct {
	ID      uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Comment string `gorm:"size:255;not null" json:"comment"`
	UserId  uint32 `gorm:"not null" json:"user_id"`
	PostId  uint64 `gorm:"not null" json:"post_id"`
	User    User   `json:"user"`
	Post    Post   `json:"post"`
}

var (
	ErrCommentEmptyComment = errors.New("Comment can't be empty")
)

func ValidateComment(c Comment) error {
	if c.Comment == "" {
		return ErrCommentEmptyComment
	}
	return nil
}
