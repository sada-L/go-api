package middleware

import (
	"github.com/gin-gonic/gin"
	"go-server/domain"
	"go-server/internal/logger"
	"go-server/internal/tokenutil"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := tokenutil.IsAuthorized(authToken, secret)
			if authorized {
				userID, err := tokenutil.ExtractIDFromToken(authToken, secret)
				if err != nil {
					logger.Error("failed to tale id", zap.Error(err))
					c.JSON(http.StatusUnauthorized, domain.Error{Message: err.Error()})
					c.Abort()
					return
				}
				c.Set("x-user-id", userID)
				c.Next()
				return
			}
			logger.Error("token is invalid", zap.Error(err))
			c.JSON(http.StatusUnauthorized, domain.Error{Message: err.Error()})
			c.Abort()
			return
		}
		logger.Error("incorrect authorization header")
		c.JSON(http.StatusUnauthorized, domain.Error{Message: "Not authorized"})
		c.Abort()
	}
}
