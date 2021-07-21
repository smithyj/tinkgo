package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tinkgo/service/app/api/internal/logic/account"
	"tinkgo/service/app/api/internal/svc"
)

func UpdateProfileHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req account.UpdateProfileRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			panic(err)
		}
		l := account.NewUpdateProfileLogic(c, svcCtx)
		resp, err := l.UpdateProfile(&req)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, resp)
	}
}
