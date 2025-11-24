package models

import "time"

type Article struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `gorm:"unique;not null" json:"title"`
	Content   string         `gorm:"not null" json:"content"`
	Photos    []ArticlePhoto `gorm:"constraint:OnDelete:CASCADE" json:"photos"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
