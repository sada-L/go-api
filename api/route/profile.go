package route

import (
	"github.com/gin-gonic/gin"
	"go-server/api/controller"
	"go-server/bootstrap"
	"go-server/repository"
	"go-server/usecase"
	"gorm.io/gorm"
	"time"
)

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}
