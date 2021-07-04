package handler

import (
	"github.com/gin-contrib/cors"
	"tinkgo/pkg/tinkgo/httpx"
	"tinkgo/service/app/api/internal/handler/account"
	"tinkgo/service/app/api/internal/handler/passport"
	"tinkgo/service/app/api/internal/middleware"
	"tinkgo/service/app/api/internal/svc"
)

func NewRouter(server *httpx.Server, svcCtx *svc.ServiceContext) {
	engine := server.Engine()
	// 全局中间件
	{
		engine.Use(cors.Default())
		engine.Use(middleware.ErrorMiddleware(svcCtx))
	}
	// 通行证
	{
		g := engine.Group("/passport")
		g.POST("/register", passport.RegisterHandler(svcCtx))
		g.POST("/login", passport.LoginHandler(svcCtx))
		g.POST("/logout", passport.LogoutHandler(svcCtx))
	}
	// 账号
	{
		g := engine.Group("/account")
		g.GET("/:id", account.ProfileHandler(svcCtx))
		g.PUT("/:id", account.UpdateProfileHandler(svcCtx))
	}
}
