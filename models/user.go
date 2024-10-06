package models

import "gorm.io/gorm"

// Модель пользователя
type User struct {
	gorm.Model
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
