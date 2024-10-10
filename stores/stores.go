package stores

import (
	"gorm.io/gorm"
)

type Stores struct {
	DB    *gorm.DB
	User  UserStore
	Image ImageStore
}

func New(db *gorm.DB) *Stores {
	return &Stores{
		DB:    db,
		User:  &userStore{db},
		Image: &imageStore{db},
	}
}
