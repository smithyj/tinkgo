package passport

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tinkgo/service/app/api/internal/logic/passport"
	"tinkgo/service/app/api/internal/svc"
)

func LogoutHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req passport.LogoutRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			panic(err)
		}
		l := passport.NewLogoutLogic(c, svcCtx)
		resp, err := l.Logout(&req)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, resp)
	}
}