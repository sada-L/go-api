package route

import (
	"github.com/gin-gonic/gin"
	"go-server/api/controller"
	"go-server/bootstrap"
	"gorm.io/gorm"
	"time"
)

func NewHomeRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	sc := controller.HomeController{
		Env: env,
	}
	group.GET("/", sc.Home)
}
