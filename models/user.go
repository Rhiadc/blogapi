package models

import "time"

type User struct {
	ID        uint32    `gorm:"primary_key; NOT NULL AUTO_INCREMENT" json:"id"`
	Nickname  string    `gorm:"type:varchar(20);not null;unique_index" json:"nickname"`
	Email     string    `gorm:"type:varchar(40);not null;unique_index" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Posts     []Post    `gorm:"ForeignKey:UserId" json:"posts"`
	Comments  []Comment `gorm:"ForeignKey:UserId"`
}
