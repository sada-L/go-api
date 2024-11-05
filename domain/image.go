package domain

import (
	"context"
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	Filename string `json:"filename" example:"Image.jpg" form:"title" binding:"required"`
}

type ImageRepository interface {
	GetAll() ([]Image, error)
	GetById(id int) (Image, error)
	Create(image *Image) error
	CreateMany(image *[]Image) error
	Update(image *Image) (uint, error)
	DeleteById(id int) error
}

type ImageUsecase interface {
	GetAll(c context.Context) ([]Image, error)
	GetByID(c context.Context, id int) (Image, error)
	Create(c context.Context, image *Image) error
	CreateMany(c context.Context, image *[]Image) error
	Update(c context.Context, image *Image) (uint, error)
	DeleteByID(c context.Context, id int) error
}
