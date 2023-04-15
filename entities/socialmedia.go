package entities

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Socialmedia struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"not null" valid:"required" json:"name"`
	SocialMediaUrl string    `gorm:"not null" valid:"required" json:"social_media_url"`
	UserID         uint      `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (sm *Socialmedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(sm)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return nil
}

func (sm *Socialmedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(sm)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
