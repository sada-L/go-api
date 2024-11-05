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

type LoginController struct {
	LoginUsecase LoginUsecase
	Env          *bootstrap.Env
}

// Login
// @Summary Авторизация.
// @Description Авторизация по логину и паролю.
// @Tags Auth
// @Produce json
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Success 200 {object} domain.LoginResponse "JWT tokens"
// @Failure 401 {object} domain.Error
// @Failure 500 {object} domain.Error
// @Router /api/login [post]
func (lc *LoginController) Login(c *gin.Context) {
	var request LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		logger.Error("failed to bind request", zap.Error(err))
		c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
		return
	}

	user, err := lc.LoginUsecase.GetUserByEmail(c, request.Email)
	if err != nil {
		logger.Error("user not found", zap.Error(err))
		c.JSON(http.StatusNotFound, Error{Message: "User not found with the given email"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		logger.Error("Invalid credentials", zap.Error(err))
		c.JSON(http.StatusUnauthorized, Error{Message: "Invalid credentials"})
		return
	}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(&user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		logger.Error("failed to create access token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, Error{Message: err.Error()})
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(&user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		logger.Error("failed to create refresh token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, Error{Message: err.Error()})
		return
	}

	loginResponse := LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)
}
