package stores

import (
	. "go-api/models"
	"gorm.io/gorm"
)

type (
	UserStore interface {
		Get() []User
		Create(user *User)
		Update(user *User)
		DeleteById(id int)
	}

	userStore struct {
		*gorm.DB
	}
)

func (s *userStore) Get() []User {
	var users []User
	s.DB.Find(&users)
	return users
}
func (s *userStore) Create(user *User) {
	s.DB.Create(&user)
}

func (s *userStore) Update(user *User) {
	s.DB.Save(&user)
}

func (s *userStore) DeleteById(id int) {
	s.DB.Delete(&User{}, id)
}
