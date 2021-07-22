package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
	"tinkgo/service/app/api/internal/svc"
)

func Logger(srvCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = fmt.Sprintf("%v?%v", path, raw)
		}

		c.Next()

		zap.L().WithOptions(zap.WithCaller(false)).Info(
			"",
			zap.String("host", c.Request.RemoteAddr),
			zap.String("method", c.Request.Method),
			zap.String("uri", c.Request.RequestURI),
			zap.String("proto", c.Request.Proto),
			zap.String("status", fmt.Sprintf("%v", c.Writer.Status())),
			zap.String("duration", time.Now().Sub(start).String()),
		)
	}
}
