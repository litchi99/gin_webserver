package main

import (
	"base-gin/core"
	"base-gin/global"
	"base-gin/initialize"
	"fmt"
	_ "github.com/fvbock/endless"
)

func main() {
	//// 创建一个默认的路由引擎
	//r := gin.Default()
	//// GET：请求方式；/hello：请求的路径
	//// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	//r.GET("/hello", func(c *gin.Context) {
	//	// c.JSON：返回JSON格式的数据
	//	c.JSON(200, gin.H{
	//		"message": "Hello world!",
	//	})
	//})
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	//r.Run()

	// endless
	//s := endless.NewServer(":8081", r)
	//s.ReadHeaderTimeout = 10 * time.Millisecond
	//s.WriteTimeout = 10 * time.Second
	//s.MaxHeaderBytes = 1 << 20
	//s.ListenAndServe()
	// net/http
	//s := &http.Server{
	//	Addr: ":8081",
	//	Handler: r,
	//	ReadTimeout: 5*time.Second,
	//	WriteTimeout: 5*time.Second,
	//}
	//s.ListenAndServe()
	fmt.Println("vp: ", global.GVA_VP.Get("mysql.password"))
	fmt.Println("config: ", global.GVA_CONFIG.Mysql.Password)
	fmt.Println("vp: ", global.GVA_VP)
	initialize.Mysql()
	core.RunWindowsServer()
}
