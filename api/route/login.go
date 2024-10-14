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

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db)
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}
