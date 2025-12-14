package model

import "time"

// Struct giống như Class trong C#
// `json:"..."` để map khi trả API
// `gorm:"..."` để config DB (Primary Key, Index)
type ShortLink struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	OriginalURL string    `gorm:"not null" json:"original_url"`
	ShortCode   string    `gorm:"uniqueIndex;not null" json:"short_code"`
	ClickCount  int       `gorm:"default:0" json:"click_count"`
	CreatedAt   time.Time `json:"created_at"`
}
