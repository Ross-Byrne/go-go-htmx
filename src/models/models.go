package models

import "time"

type Post struct {
	ID        uint64 `gorm:"primaryKey"`
	Title     string `form:"title"`
	Text      string `form:"text"`
	AuthorID  uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}
