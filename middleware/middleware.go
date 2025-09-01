package middleware

import (
	"net/http"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"github.com/rushairer/gouno"
	"{{.ModulePath}}/config"
)

func TimeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(config.GlobalConfig.WebServerConfig.RequestTimeout),
		timeout.WithResponse(
			func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gouno.ErrRequestTimeoutResponse)
			},
		),
	)
}

func RecoveryMiddleware() gin.HandlerFunc {
	return gin.CustomRecovery(
		func(ctx *gin.Context, err any) {
			ctx.JSON(http.StatusOK, gouno.ErrInternalServerErrorResponse)
		},
	)
}
