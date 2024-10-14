package route

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"go-server/bootstrap"
	"gorm.io/gorm"
	"time"
)

func NewSwaggerRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
