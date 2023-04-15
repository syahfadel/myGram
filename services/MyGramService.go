package services

import (
	"errors"
	"fmt"
	"myGram/entities"

	"gorm.io/gorm"
)

type MyGramService struct {
	DB *gorm.DB
}

func (ms *MyGramService) CreateUser(user entities.User) (entities.User, error) {
	if err := ms.DB.Debug().Create(&user).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (ms *MyGramService) Login(user entities.User) (entities.User, error) {
	if err := ms.DB.Debug().Where("email = ?", user.Email).Take(&user).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (ms *MyGramService) GetAllPhoto() ([]entities.Photo, error) {
	var photos []entities.Photo
	if err := ms.DB.Debug().Find(&photos).Error; err != nil {
		return []entities.Photo{}, err
	}
	return photos, nil
}

func (ms *MyGramService) GetOnePhoto(id uint) (entities.Photo, error) {
	var photo entities.Photo
	if err := ms.DB.Debug().Where("id = ?", id).Take(&photo).Error; err != nil {
		return entities.Photo{}, err
	}
	return photo, nil
}

func (ms *MyGramService) CreatePhoto(photo entities.Photo) (entities.Photo, error) {
	if err := ms.DB.Debug().Create(&photo).Error; err != nil {
		return entities.Photo{}, err
	}
	return photo, nil
}

func (ms *MyGramService) UpdatePhoto(photo entities.Photo) (entities.Photo, error) {
	res := ms.DB.Debug().Model(&photo).Where("id = ? AND user_id = ?", photo.ID, photo.UserID).Updates(&photo)
	if res.Error != nil {
		return entities.Photo{}, res.Error
	}
	if res.RowsAffected == 0 {
		return entities.Photo{}, errors.New("no data updated")
	}
	return photo, nil
}

func (ms *MyGramService) DeletePhoto(photoId, userId uint) error {
	result := ms.DB.Debug().Where("user_id = ? AND id = ?", userId, photoId).Delete(&entities.Photo{})
	if result.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("Post with userId %v and photoId %v not available", userId, photoId))
	}
	return nil
}

func (ms *MyGramService) GetAllComment() error {
	//todo
	return nil
}

func (ms *MyGramService) GetOneComment() error {
	//todo
	return nil
}

func (ms *MyGramService) CreateComment() error {
	//todo
	return nil
}

func (ms *MyGramService) UpdateComment() error {
	//todo
	return nil
}

func (ms *MyGramService) DeleteComment() error {
	//todo
	return nil
}

func (ms *MyGramService) GetAllSocialMedia() error {
	//todo
	return nil
}

func (ms *MyGramService) GetOneSocialMedia() error {
	//todo
	return nil
}

func (ms *MyGramService) CreateSocialMedia() error {
	//todo
	return nil
}

func (ms *MyGramService) UpdateSocialMedia() error {
	//todo
	return nil
}

func (ms *MyGramService) DeleteSocialMedia() error {
	//todo
	return nil
}
