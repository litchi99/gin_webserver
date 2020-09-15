package route

import (
	"base-gin/api/v1"
	"github.com/gin-gonic/gin"
)
// 路由初始化,相对访问路径与接口句柄的绑定,接口类型(GET/POST/PUT...)的定义
// route service dao
func InitBaseRoute(Router *gin.RouterGroup) (R gin.IRouter)  {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("hello",v1.HelloWorld)

	}
	return BaseRouter
}
