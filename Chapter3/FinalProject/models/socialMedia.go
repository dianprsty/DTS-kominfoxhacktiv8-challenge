package models

import "time"

type (
	SocialMedia struct {
		ID             uint      `json:"id" gorm:"primary_key"`
		Name           string    `gorm:"not null;" json:"name"`
		SocialMediaUrl string    `gorm:"not null;" json:"social_media_url"`
		UserId         uint      `gorm:"not null;unique" json:"user_id"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
		User           User      `json:"-"`
	}
)
