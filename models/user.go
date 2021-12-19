package models

import (
	"errors"
	"time"
)

type User struct {
	ID        uint64    `gorm:"primary_key; NOT NULL AUTO_INCREMENT" json:"id"`
	Nickname  string    `gorm:"type:varchar(20);not null;unique_index" json:"nickname"`
	Email     string    `gorm:"type:varchar(40);not null;unique_index" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Posts     []Post    `gorm:"ForeignKey:UserId" json:"posts,omitempty"`
	Comments  []Comment `gorm:"ForeignKey:UserId" json:"comments,omitempty"`
}

type UserResponse struct {
	ID        uint64    `json:"id"`
	Nickname  string    `json:"nickname"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Posts     []Post    `json:"posts,omitempty"`
	Comments  []Comment `json:"comments,omitempty"`
}

var (
	ErrUserEmptyNickname = errors.New("user.nickname can't be empty")
	ErrUserEmptyEmail    = errors.New("user.email can't be empty")
	ErrUserEmptyPassword = errors.New("user.password can't be empty")
)

func ValidateUser(u User) error {
	if u.Nickname == "" {
		return ErrUserEmptyNickname
	} else if u.Email == "" {
		return ErrUserEmptyEmail
	} else if u.Password == "" {
		return ErrUserEmptyPassword
	}
	return nil
}
