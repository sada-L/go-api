package handlers

import (
	"github.com/labstack/echo/v4"
	"go-api/logger"
	. "go-api/models"
	"go-api/services"
	"go-api/utils"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type (
	UserHandler interface {
		GetUser(c echo.Context) error
		CreateUser(c echo.Context) error
		UpdateUser(c echo.Context) error
		DeleteUserById(c echo.Context) error
	}

	userHandler struct {
		services.UserService
	}
)

// GetUser
// @Summary Получить всех пользователей.
// @Description Возвращает список всех пользователей.
// @Tags User
// @Accept */*
// @Param Authorization header string true "'Bearer _YOUR_TOKEN_'"
// @Security Bearer Authentication
// @Produce json
// @Success 200 {object} []models.User
// @Failure 500 {object} utils.Error
// @Router /api/v1/user [get]
func (h *userHandler) GetUser(c echo.Context) error {
	r, err := h.UserService.GetUsers()

	if err != nil {
		logger.Error("failed to get user", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, r)
}

// CreateUser
// @Summary Создание пользователя.
// @Description Создает нового пользователя.
// @Tags User
// @Accept */*
// @Param Authorization header string true "'Bearer _YOUR_TOKEN_'"
// @Param user body User true "User Info"
// @Security Bearer Authentication
// @Produce json
// @Success 200 {integer} integer "Created ID"
// @Failure 400 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /api/v1/user [post]
func (h *userHandler) CreateUser(c echo.Context) error {
	var u *User

	if err := c.Bind(&u); err != nil {
		logger.Error("failed to bind user", zap.Error(err))
		return c.JSON(http.StatusBadRequest, utils.Error{Message: err.Error()})
	}

	r, err := h.UserService.CreateUser(u)
	if err != nil {
		logger.Error("failed to create user", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, r)
}

// UpdateUser
// @Summary Обновить пользователся.
// @Description Обновляет пользователя.
// @Tags User
// @Accept */*
// @Param Authorization header string true "'Bearer _YOUR_TOKEN_'"
// @Param user body User true "User Info"
// @Security Bearer Authentication
// @Produce json
// @Failure 400 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /api/v1/user [put]
func (h *userHandler) UpdateUser(c echo.Context) error {
	var u *User

	if err := c.Bind(&u); err != nil {
		logger.Error("failed to bind user", zap.Error(err))
		return c.JSON(http.StatusBadRequest, utils.Error{Message: "args is invalid"})
	}

	r, err := h.UserService.UpdateUser(u)

	if err != nil {
		logger.Error("failed to update user", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, r)
}

// DeleteUserById
// @Summary Удалить пользователся по ID.
// @Description Удаляет пользователя по ID.
// @Tags User
// @Accept */*
// @Security Bearer Authentication
// @Produce json
// @Param id path int true "User id"
// @Param Authorization header string true "'Bearer _YOUR_TOKEN_'"
// @Success 200 {string} string "OK"
// @Failure 400 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /api/v1/user/{id} [delete]
func (h *userHandler) DeleteUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error("failed to parse id", zap.Error(err))
		return c.JSON(http.StatusBadRequest, utils.Error{Message: "ID is invalid"})
	}

	err = h.UserService.DeleteUser(id)
	if err != nil {
		logger.Error("failed to delete user", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, "OK")
}