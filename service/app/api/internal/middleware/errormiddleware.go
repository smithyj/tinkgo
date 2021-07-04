package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tinkgo/service/app/api/internal/svc"
	"tinkgo/service/pkg/errorx"
)

func ErrorMiddleware(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case *errorx.CodeError:
					c.JSON(http.StatusOK, err)
				default:
					e := err.(error)
					c.JSON(http.StatusInternalServerError, errorx.NewCodeErrorWithMsg(e.Error()))
				}
			}
		}()
		c.Next()
	}
}
