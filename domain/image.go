package domain

import "context"

type Image struct {
	Id       int    `json:"id"`
	Filename string `json:"filename" example:"Image.jpg" form:"title" binding:"required"`
}

type ImageRepository interface {
	GetAll() ([]Image, error)
	GetById(id int) (Image, error)
	Create(image *Image) (int, error)
	Update(image *Image) (int, error)
	DeleteById(id int) error
}

type ImageUsecase interface {
	GetAll(c context.Context) ([]Image, error)
	GetByID(c context.Context, id int) (Image, error)
	Create(c context.Context, image *Image) (int, error)
	Update(c context.Context, image *Image) (int, error)
	DeleteByID(c context.Context, id int) error
}
