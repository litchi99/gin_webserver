package initialize

import (
	"base-gin/route"
	"github.com/gin-gonic/gin"
)
// 路由组
func Routers() *gin.Engine {
	var Router = gin.Default()
	ApiGroup := Router.Group("")
	route.InitBaseRoute(ApiGroup)
	return Router
}