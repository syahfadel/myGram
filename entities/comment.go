package entities

import "time"

type Comment struct {
	ID        uint      `gorm:primaryKey json:"id"`
	UserID    uint      `json:"user_id"`
	PhotoID   uint      `json:"photo_id"`
	Message   string    `gorm:"not null" json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at`
}
