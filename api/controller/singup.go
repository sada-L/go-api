package controller

import (
	"github.com/gin-gonic/gin"
	"go-server/bootstrap"
	. "go-server/domain"
	"go-server/internal/logger"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type SignupController struct {
	SignupUsecase SignupUsecase
	Env           *bootstrap.Env
}

// Signup
// @Summary  Регистрация.
// @Description Регистарация по логину и паролю.
// @Tags Auth
// @Produce json
// @Param name formData string true "Username"
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Success 200 {object} domain.SignupResponse "JWT tokens"
// @Failure 401 {object} domain.Error
// @Failure 500 {object} domain.Error
// @Router /signup [post]
func (ctl *SignupController) Signup(c *gin.Context) {
	var request SignupRequest

	err := c.ShouldBind(&request)
	if err != nil {
		logger.Error("failed to bind request", zap.Error(err))
		c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
		return
	}

	_, err = ctl.SignupUsecase.GetUserByEmail(c, request.Email)
	if err == nil {
		logger.Error("user already exists", zap.Error(err))
		c.JSON(http.StatusConflict, Error{Message: "User already exists with the given email"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		logger.Error("failed with password", zap.Error(err))
		c.JSON(http.StatusInternalServerError, Error{Message: err.Error()})
		return
	}

	request.Password = string(encryptedPassword)

	user := User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err = ctl.SignupUsecase.Create(c, &user)
	if err != nil {
		logger.Error("failed to create user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, Error{Message: err.Error()})
		return
	}

	accessToken, err := ctl.SignupUsecase.CreateAccessToken(&user, ctl.Env.AccessTokenSecret, ctl.Env.AccessTokenExpiryHour)
	if err != nil {
		logger.Error("failed to create access token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, Error{Message: err.Error()})
		return
	}

	refreshToken, err := ctl.SignupUsecase.CreateRefreshToken(&user, ctl.Env.RefreshTokenSecret, ctl.Env.RefreshTokenExpiryHour)
	if err != nil {
		logger.Error("failed to create refresh token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, Error{Message: err.Error()})
		return
	}

	signupResponse := SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, signupResponse)
}
