package models

import "time"

type (
	Comment struct {
		ID        uint      `json:"id" gorm:"primary_key"`
		UserId    uint      `gorm:"not null;unique" json:"user_id"`
		PhotoId   uint      `gorm:"not null;unique" json:"photo_id"`
		Message   string    `gorm:"not null;" json:"message"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		User      User      `json:"-"`
		Photo     Photo     `json:"-"`
	}
)
