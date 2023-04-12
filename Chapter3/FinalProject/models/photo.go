package models

import "time"

type (
	Photo struct {
		ID        uint      `json:"id" gorm:"primary_key"`
		Title     string    `gorm:"not null;" json:"title"`
		Caption   string    `gorm:"not null;" json:"caption"`
		PhotoUrl  string    `gorm:"not null;" json:"photo_url"`
		UserId    uint      `gorm:"not null;unique" json:"user_id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Comments  []Comment `json:"-"`
		User      User      `json:"-"`
	}
)
