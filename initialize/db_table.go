package initialize

import (
	"base-gin/global"
	"base-gin/model"
)

func DBTables()  {
	db := global.GVA_DB
	db.AutoMigrate(model.SysUser{})
	global.GVA_LOG.Debug("register table successfully")
}
