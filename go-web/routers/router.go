package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hearecho/go-pro/go-web/pkg/resp"
	"github.com/hearecho/go-pro/go-web/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	//路由逻辑
	r.GET("/test", func(context *gin.Context) {
		r := resp.R{}
		r = r.Ok().SetData("test").SetPath(context.Request.URL.Path)
		context.JSON(200,r)
	})
	return r
}
