package stores

import (
	. "go-api/models"
	"gorm.io/gorm"
)

type (
	ImageStore interface {
		Get() ([]Image, error)
		GetById(id int) (Image, error)
		Create(image *Image) (int, error)
		Update(image *Image) (int, error)
		DeleteById(id int) error
	}

	imageStore struct {
		*gorm.DB
	}
)

func (s *imageStore) Get() ([]Image, error) {
	var images []Image
	if err := s.DB.Find(&images).Error; err != nil {
		return nil, err
	}

	return images, nil
}

func (s *imageStore) GetById(id int) (Image, error) {
	var image Image
	if err := s.DB.First(&image, id).Error; err != nil {
		return image, err
	}

	return image, nil
}

func (s *imageStore) Create(image *Image) (int, error) {
	var id int

	if err := s.DB.Create(image).Error; err != nil {
		return 0, err
	}
	id = image.Id

	return id, nil
}

func (s *imageStore) Update(image *Image) (int, error) {
	var id int

	if err := s.DB.Save(&image).Error; err != nil {
		return 0, err
	}
	id = image.Id

	return id, nil
}

func (s *imageStore) DeleteById(id int) error {
	if err := s.DB.First(&Image{}, id).Error; err != nil {
		return err
	}

	if err := s.DB.Delete(&Image{}, id).Error; err != nil {
		return err
	}

	return nil
}
