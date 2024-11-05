package repository

import (
	. "go-server/domain"
	"gorm.io/gorm"
)

type imageRepository struct {
	DB *gorm.DB
}

func NewImageRepository(db *gorm.DB) ImageRepository {
	return &imageRepository{DB: db}
}

func (s *imageRepository) GetAll() ([]Image, error) {
	var images []Image
	if err := s.DB.Find(&images).Error; err != nil {
		return nil, err
	}

	return images, nil
}

func (s *imageRepository) GetById(id int) (Image, error) {
	var image Image
	if err := s.DB.First(&image, id).Error; err != nil {
		return image, err
	}

	return image, nil
}

func (s *imageRepository) Create(image *Image) error {
	if err := s.DB.Create(image).Error; err != nil {
		return err
	}

	return nil
}

func (s *imageRepository) CreateMany(image *[]Image) error {
	if err := s.DB.Create(image).Error; err != nil {
		return err
	}
	return nil
}

func (s *imageRepository) Update(image *Image) (uint, error) {
	if err := s.DB.Save(&image).Error; err != nil {
		return 0, err
	}
	id := image.ID

	return id, nil
}

func (s *imageRepository) DeleteById(id int) error {
	if err := s.DB.First(&Image{}, id).Error; err != nil {
		return err
	}

	if err := s.DB.Delete(&Image{}, id).Error; err != nil {
		return err
	}

	return nil
}
