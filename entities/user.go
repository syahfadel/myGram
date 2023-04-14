package entities

import (
	"myGram/helpers"
	"time"

	"github.com/asaskevich/govalidator"

	"gorm.io/gorm"
)

type User struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Username     string `gorm:"not null;uniqueIndex" valid:"required" json:"username"`
	Email        string `gorm:"not null;uniqueIndex" form:"email" valid:"required,email"`
	Password     string `gorm:"not null" form:"password" valid:"required,length(6|50)"`
	Age          uint   `gorm:"not null" valid:"required,range(9|100)"`
	Photos       []Photo
	Socialmedias []Socialmedia
	Comments     []Comment
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errUpdate := govalidator.ValidateStruct(u)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
