package services

import (
	. "go-api/models"
	"go-api/stores"
)

type (
	ImageService interface {
		GetImages() ([]Image, error)
		GetImageById(id int) (Image, error)
		CreateImage(a *Image) (int, error)
		UpdateImage(a *Image) (int, error)
		DeleteImage(id int) error
	}

	imageService struct {
		stores *stores.Stores
	}
)

func (s *imageService) GetImages() ([]Image, error) {
	r, err := s.stores.Image.Get()
	return r, err
}

func (s *imageService) GetImageById(id int) (Image, error) {
	r, err := s.stores.Image.GetById(id)
	return r, err
}

func (s *imageService) CreateImage(a *Image) (int, error) {
	r, err := s.stores.Image.Create(a)
	return r, err
}

func (s *imageService) UpdateImage(a *Image) (int, error) {
	r, err := s.stores.Image.Update(a)
	return r, err
}

func (s *imageService) DeleteImage(id int) error {
	err := s.stores.Image.DeleteById(id)
	return err
}
