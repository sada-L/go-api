package controller

import (
	"github.com/gin-gonic/gin"
	"go-server/domain"
	"go-server/internal/logger"
	"go.uber.org/zap"
	"net/http"
)

type ProfileController struct {
	ProfileUsecase domain.ProfileUsecase
}

// Fetch
// @Summary Просмотреть профиль.
// @Description Просморт профиля пользователя по токену доступа.
// @Tags Auth
// @Produce json
// @Param Authorization header string true "'Bearer _YOUR_TOKEN_'"
// @Security Bearer Authentication
// @Success 200 {object} domain.Profile "Profile"
// @Failure 401 {object} domain.Error
// @Failure 500 {object} domain.Error
// @Router /profile [get]
func (pc *ProfileController) Fetch(c *gin.Context) {
	userID := c.GetString("x-user-id")

	profile, err := pc.ProfileUsecase.GetProfileByID(c, userID)
	if err != nil {
		logger.Error("failed to get profile", zap.Error(err))
		c.JSON(http.StatusInternalServerError, domain.Error{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}
