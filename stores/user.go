package stores

import (
	. "go-api/models"
	"gorm.io/gorm"
)

type (
	UserStore interface {
		Get() ([]User, error)
		Create(user *User) (int, error)
		Update(user *User) (int, error)
		DeleteById(id int) error
	}

	userStore struct {
		*gorm.DB
	}
)

func (s *userStore) Get() ([]User, error) {
	var users []User
	if err := s.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
func (s *userStore) Create(user *User) (int, error) {
	var id int

	if err := s.DB.Create(&user).Error; err != nil {
		return 0, err
	}
	id = user.Id

	return id, nil
}

func (s *userStore) Update(user *User) (int, error) {
	var id int

	if err := s.DB.Save(&user).Error; err != nil {
		return 0, err
	}
	id = user.Id

	return id, nil
}

func (s *userStore) DeleteById(id int) error {
	if err := s.DB.Delete(&User{}, id).Error; err != nil {
		return err
	}
	return nil
}
