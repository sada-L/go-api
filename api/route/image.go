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

func NewImageRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, group *gin.RouterGroup) {
	ir := repository.NewImageRepository(db)
	ic := &controller.ImageController{
		ImageUsecase: usecase.NewImageUsecase(ir, timeout),
	}
	group.GET("/image/:id", ic.GetImage)
	group.POST("/image/single", ic.UploadImage)
	group.POST("/image/multi", ic.UploadMultipleImages)
	group.POST("/image/multi/zip", ic.UploadZipFiles)
	group.DELETE("/image/:id", ic.DeleteImageById)
}
