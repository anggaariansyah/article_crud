package models

import "time"

type ArticlePhoto struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ArticleID uint      `json:"article_id"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
