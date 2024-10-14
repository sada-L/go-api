package repository

import (
	. "go-server/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) Create(user *User) error {
	if err := r.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetByID(id string) (User, error) {
	var user User
	if err := r.DB.First(&user, id).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) GetByEmail(email string) (User, error) {
	var user User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
