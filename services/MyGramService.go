package services

import (
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

func (ms *MyGramService) GetAllPhoto() error {
	//todo
	return nil
}

func (ms *MyGramService) GetOnePhoto() error {
	//todo
	return nil
}

func (ms *MyGramService) CreatePhoto() error {
	//todo
	return nil
}

func (ms *MyGramService) UpdatePhoto() error {
	//todo
	return nil
}

func (ms *MyGramService) DeletePhoto() error {
	//todo
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
