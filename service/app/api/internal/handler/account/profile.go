package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tinkgo/service/app/api/internal/logic/account"
	"tinkgo/service/app/api/internal/svc"
)

func ProfileHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req account.ProfileRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			_ = c.Error(err)
			return
		}
		l := account.NewProfileLogic(c, svcCtx)
		resp, err := l.Profile(&req)
		if err != nil {
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, resp)
	}
}