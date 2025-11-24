package models

import "time"

type ArticlePhoto struct {
	ID        uint      `gorm:"primaryKey"`
	ArticleID uint      `gorm:"not null"`
	URL       string    `gorm:"not null"`
	Article   Article   `gorm:"constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
