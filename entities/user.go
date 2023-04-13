package entities

import (
	// "github.com/asaskevich/govalidator"
	// "gorm.io/gorm"
	"time"
)

type User struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Username     string `gorm:"not null;uniqueIndex" json:"username"`
	Email        string `gorm:"not null;uniqueIndex" form:"email" valid:"required,email"`
	Password     string `gorm:"not null" form:"password" valid:"required,length(6|50)"`
	Age          uint   `gorm:"not null" valid:"required,range(8|100)"`
	Photos       []Photo
	Socialmedias []Socialmedia
	Comments     []Comment
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
