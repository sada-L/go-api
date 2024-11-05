package controller

import (
	"github.com/gin-gonic/gin"
	"go-server/bootstrap"
	. "go-server/domain"
	"go-server/internal/logger"
	"go.uber.org/zap"
	"net/http"
)

type RefreshTokenController struct {
	RefreshTokenUsecase RefreshTokenUsecase
	Env                 *bootstrap.Env
}

// RefreshToken
// @Summary Обновление токенов.
// @Description Обновить токены с помощью токена обновления.
// @Tags Auth
// @Produce json
// @Param refreshToken formData string true "Refresh token"
// @Success 200 {object} domain.RefreshTokenResponse "JWT Tokens"
// @Failure 400 {object} domain.Error
// @Failure 401 {object} domain.Error
// @Failure 500 {object} domain.Error
// @Router /api/refresh [post]
func (rtc *RefreshTokenController) RefreshToken(c *gin.Context) {
	var request RefreshTokenRequest

	err := c.ShouldBind(&request)
	if err != nil {
		logger.Error("failed to bind request", zap.Error(err))
		c.JSON(http.StatusBadRequest, Error{Message: err.Error()})
		return
	}

	id, err := rtc.RefreshTokenUsecase.ExtractIDFromToken(request.RefreshToken, rtc.Env.RefreshTokenSecret)
	if err != nil {
		logger.Error("user not found", zap.Error(err))
		c.JSON(http.StatusUnauthorized, Error{Message: "User not found"})
		return
	}

	user, err := rtc.RefreshTokenUsecase.GetUserByID(c, id)
	if err != nil {
		logger.Error("user not found", zap.Error(err))
		c.JSON(http.StatusUnauthorized, Error{Message: "User not found"})
		return
	}

	accessToken, err := rtc.RefreshTokenUsecase.CreateAccessToken(&user, rtc.Env.AccessTokenSecret, rtc.Env.AccessTokenExpiryHour)
	if err != nil {
		logger.Error("failed to create access token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, Error{Message: err.Error()})
		return
	}

	refreshToken, err := rtc.RefreshTokenUsecase.CreateRefreshToken(&user, rtc.Env.RefreshTokenSecret, rtc.Env.RefreshTokenExpiryHour)
	if err != nil {
		logger.Error("failed to create refresh token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, Error{Message: err.Error()})
		return
	}

	refreshTokenResponse := RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, refreshTokenResponse)
}
