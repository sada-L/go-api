package usecase

import (
	"context"
	"go-server/domain"
	"time"
)

type imageUsecase struct {
	imageRepository domain.ImageRepository
	contextTimeout  time.Duration
}

func NewImageUsecase(imageRepository domain.ImageRepository, timeout time.Duration) domain.ImageUsecase {
	return &imageUsecase{
		imageRepository: imageRepository,
		contextTimeout:  timeout,
	}
}

func (u *imageUsecase) GetAll(c context.Context) ([]domain.Image, error) {
	_, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.imageRepository.GetAll()
}

func (u *imageUsecase) GetByID(c context.Context, id int) (domain.Image, error) {
	_, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.imageRepository.GetById(id)
}

func (u *imageUsecase) Create(c context.Context, image *domain.Image) (int, error) {
	_, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.imageRepository.Create(image)
}

func (u *imageUsecase) Update(c context.Context, image *domain.Image) (int, error) {
	_, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.imageRepository.Update(image)
}

func (u *imageUsecase) DeleteByID(c context.Context, id int) error {
	_, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.imageRepository.DeleteById(id)
}
