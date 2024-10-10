package services

import (
	. "go-api/models"
	"go-api/stores"
)

type (
	UserService interface {
		GetUserById(id int) (User, error)
		GetUsers() ([]User, error)
		CreateUser(a *User) (int, error)
		UpdateUser(a *User) (int, error)
		DeleteUser(id int) error
	}

	userService struct {
		stores *stores.Stores
	}
)

func (s *userService) GetUserById(id int) (User, error) {
	r, err := s.stores.User.GetById(id)
	return r, err
}

func (s *userService) GetUsers() ([]User, error) {
	r, err := s.stores.User.GetAll()
	return r, err
}

func (s *userService) CreateUser(a *User) (int, error) {
	r, err := s.stores.User.Create(a)
	return r, err
}

func (s *userService) UpdateUser(a *User) (int, error) {
	r, err := s.stores.User.Update(a)
	return r, err
}

func (s *userService) DeleteUser(id int) error {
	err := s.stores.User.DeleteById(id)
	return err
}
