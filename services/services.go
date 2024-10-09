package services

import "go-api/stores"

type Services struct {
	User UserService
}

func New(s *stores.Stores) *Services {
	return &Services{
		User: &userService{stores: s},
	}
}
