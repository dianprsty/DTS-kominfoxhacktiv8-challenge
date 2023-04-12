package models

import "time"

type Book struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `json:"name_book"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
