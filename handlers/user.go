package handlers

import (
	"github.com/labstack/echo/v4"
	. "go-api/models"
	"go-api/services"
	"go-api/utils"
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
// @Summary Fetch a list of all users.
// @Description Fetch a list of all users.
// @Tags User
// @Accept */*
// @Param Authorization header string true "'Bearer _YOUR_TOKEN_'"
// @Security Bearer Authentication
// @Produce json
// @Success 200 {object} []User
// @Failure 500 {object} utils.Error
// @Router /api/v1/user [get]
func (h *userHandler) GetUser(c echo.Context) error {
	r := h.UserService.GetUsers()
	return c.JSON(http.StatusOK, r)
}

// CreateUser
// @Summary Создание пользователя.
// @Description Create a user.
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
		return c.JSON(http.StatusBadRequest, utils.Error{Message: err.Error()})
	}

	h.UserService.CreateUser(u)
	return c.JSON(http.StatusOK, u)
}

// UpdateUser
// @Summary Update a user.
// @Description Update a user.
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
		return c.JSON(http.StatusBadRequest, utils.Error{Message: "args is invalid"})
	}

	h.UserService.UpdateUser(u)
	return c.JSON(http.StatusOK, u)
}

// DeleteUserById
// @Summary Delete a user by ID.
// @Description Delete a user by ID.
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
		return c.JSON(http.StatusBadRequest, utils.Error{Message: "ID is invalid"})
	}

	h.UserService.DeleteUser(id)
	return c.JSON(http.StatusOK, "OK")
}
