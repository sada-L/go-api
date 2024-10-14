package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"go-server/api/middleware"
	"go-server/bootstrap"
	"gorm.io/gorm"
	"time"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, gin *gin.Engine) {
	//конфиг в разработке
	gin.Use(cors.New(cors.Config{
		AllowAllOrigins:            true,
		AllowOrigins:               nil,
		AllowOriginFunc:            nil,
		AllowOriginWithContextFunc: nil,
		AllowMethods:               []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowPrivateNetwork:        true,
		AllowHeaders:               []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials:           true,
		ExposeHeaders:              []string{"Content-Length", "Content-Type"},
		MaxAge:                     12 * time.Hour,
		AllowWildcard:              true,
		AllowBrowserExtensions:     true,
		CustomSchemas:              nil,
		AllowWebSockets:            false,
		AllowFiles:                 true,
		OptionsResponseStatusCode:  204,
	}))

	gin.Use(gzip.Gzip(gzip.DefaultCompression))
	gin.LoadHTMLGlob("./template/*")

	// Public APIs
	publicRouter := gin.Group("")
	NewHomeRouter(env, timeout, db, publicRouter)
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)
	NewSwaggerRouter(env, timeout, db, publicRouter)

	// Private APIs
	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))

	NewProfileRouter(env, timeout, db, protectedRouter)
	NewImageRouter(env, timeout, db, protectedRouter)
}
