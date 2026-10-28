package models

import "time"

type URL struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    OriginalURL string    `gorm:"not null" json:"original_url"`
    ShortCode   string    `gorm:"uniqueIndex;not null" json:"short_code"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}