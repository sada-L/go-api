package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"go-api/db"
	. "go-api/models"
)

// @Summary Создать пользователя
// @Description Создает нового пользователя
// @Tags users
// @Accept json
// @Produce json
// @Param user body User true "User Info"
// @Success 200 {object} User
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Create(&user)
	c.JSON(http.StatusOK, user)
}

// @Summary Получить всех пользователей
// @Description Возвращает список всех пользователей
// @Tags users
// @Produce json
// @Success 200 {array} User
// @Router /users [get]
func GetUsers(c *gin.Context) {
	var users []User
	db.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// @Summary Получить пользователя по ID
// @Description Возвращает пользователя по его ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Failure 404 {string} string "User not found"
// @Router /users/{id} [get]
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := db.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Summary Обновить пользователя
// @Description Обновляет информацию о пользователе
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body User true "User Info"
// @Success 200 {object} User
// @Failure 404 {string} string "User not found"
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	if err := db.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

// @Summary Удалить пользователя
// @Description Удаляет пользователя по ID
// @Tags users
// @Param id path int true "User ID"
// @Success 200 {string} string "User deleted"
// @Failure 404 {string} string "User not found"
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&User{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
