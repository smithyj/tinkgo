package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"

	"tinkgo/internal/app-api/pkg/core"
	"tinkgo/pkg/tinkgo/api"
)

func NewRouter(server *api.Server, ctx *core.Context) {
	engine := server.Engine()
	{
		// 中间件
		engine.Use(cors.Default())
	}
	{
		passport := engine.Group("/passport")
		passport.POST("/register", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "ok",
				"data": gin.H{
					"request": c.Request.RequestURI,
				},
			})
		})
		passport.POST("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "ok",
				"data": gin.H{
					"request": c.Request.RequestURI,
				},
			})
		})
		passport.POST("/logout", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "ok",
				"data": gin.H{
					"request": c.Request.RequestURI,
				},
			})
		})
	}
	{
		account := engine.Group("/account")
		account.GET("/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "ok",
				"data": gin.H{
					"request": c.Request.RequestURI,
				},
			})
		})
		account.PUT("/:id", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "ok",
				"data": gin.H{
					"request": c.Request.RequestURI,
				},
			})
		})
	}
}
