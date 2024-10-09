package services

import (
	. "go-api/models"
	"go-api/stores"
)

type (
	UserService interface {
		GetUsers() []User
		CreateUser(a *User)
		UpdateUser(a *User)
		DeleteUser(id int)
	}

	userService struct {
		stores *stores.Stores
	}
)

func (s *userService) GetUsers() []User {
	return s.stores.User.Get()
}

func (s *userService) CreateUser(a *User) {
	s.stores.User.Create(a)
}

func (s *userService) UpdateUser(a *User) {
	s.stores.User.Update(a)
}

func (s *userService) DeleteUser(id int) {
	s.stores.User.DeleteById(id)
}
