package entities

import "time"

type Socialmedia struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"not null" json:"name"`
	SocialMediaUrl string    `gorm:"not null" json:"social_media_url"`
	UserID         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
