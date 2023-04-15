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

func (ms *MyGramService) GetAllComment(photoID uint) ([]entities.Comment, error) {
	var comments []entities.Comment
	if err := ms.DB.Debug().Where("photo_id = ?", photoID).Find(&comments).Error; err != nil {
		return []entities.Comment{}, err
	}
	return comments, nil
}

func (ms *MyGramService) GetOneComment(id uint) (entities.Comment, error) {
	var comment entities.Comment
	if err := ms.DB.Debug().Where("id = ?", id).Take(&comment).Error; err != nil {
		return entities.Comment{}, err
	}
	return comment, nil
}

func (ms *MyGramService) CreateComment(comment entities.Comment) (entities.Comment, error) {
	if err := ms.DB.Debug().Create(&comment).Error; err != nil {
		return entities.Comment{}, err
	}
	return comment, nil
}

func (ms *MyGramService) UpdateComment(comment entities.Comment) (entities.Comment, error) {
	res := ms.DB.Debug().Model(&comment).Where("id = ? AND user_id = ? AND photo_id = ?", comment.ID, comment.UserID, comment.PhotoID).Updates(&comment)
	if res.Error != nil {
		return entities.Comment{}, res.Error
	}
	if res.RowsAffected == 0 {
		return entities.Comment{}, errors.New("no data updated")
	}
	return comment, nil
}

func (ms *MyGramService) DeleteComment(commentId, userId uint) error {
	result := ms.DB.Debug().Where("user_id = ? AND id = ?", userId, commentId).Delete(&entities.Comment{})
	if result.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("Post with userId %v and commentId %v not available", userId, commentId))
	}
	return nil
}

func (ms *MyGramService) GetAllSocialMedia(userID uint) ([]entities.Socialmedia, error) {
	var socialMedia []entities.Socialmedia
	if err := ms.DB.Debug().Where("user_id = ?", userID).Find(&socialMedia).Error; err != nil {
		return []entities.Socialmedia{}, err
	}
	return socialMedia, nil
}

func (ms *MyGramService) GetOneSocialMedia(id, userID uint) (entities.Socialmedia, error) {
	var socialMedia entities.Socialmedia
	if err := ms.DB.Debug().Where("id = ? AND user_id = ?", id, userID).Take(&socialMedia).Error; err != nil {
		return entities.Socialmedia{}, err
	}
	return socialMedia, nil
}

func (ms *MyGramService) CreateSocialMedia(socialMedia entities.Socialmedia) (entities.Socialmedia, error) {
	if err := ms.DB.Debug().Create(&socialMedia).Error; err != nil {
		return entities.Socialmedia{}, err
	}
	return socialMedia, nil
}

func (ms *MyGramService) UpdateSocialMedia(socialMedia entities.Socialmedia) (entities.Socialmedia, error) {
	res := ms.DB.Debug().Model(&socialMedia).Where("id = ? AND user_id = ?", socialMedia.ID, socialMedia.UserID).Updates(&socialMedia)
	if res.Error != nil {
		return entities.Socialmedia{}, res.Error
	}
	if res.RowsAffected == 0 {
		return entities.Socialmedia{}, errors.New("no data updated")
	}
	return socialMedia, nil
}

func (ms *MyGramService) DeleteSocialMedia(smId, userId uint) error {
	result := ms.DB.Debug().Where("user_id = ? AND id = ?", userId, smId).Delete(&entities.Socialmedia{})
	if result.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("Post with userId %v and smId %v not available", userId, smId))
	}
	return nil
}
