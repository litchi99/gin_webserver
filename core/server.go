package core

import (
	"base-gin/global"
	"base-gin/initialize"
	"fmt"
	"log"
	"syscall"
	"time"
)

type server interface {
	ListenAndServe() error
}
//启动服务器各个模块,如数据库的连接等
func RunWindowsServer()  {
	Router := initialize.Routers()
	address := fmt.Sprintf(":%d",global.GVA_CONFIG.System.Addr)
	s := initServer(address,Router)
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Debug("server run success on ", address)
	log.Println(syscall.Getpid(), address)  // 打印端口地址 :8081, 在endless包内
	global.GVA_LOG.Error(s.ListenAndServe())
}
