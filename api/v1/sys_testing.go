package v1

import (
	"base-gin/global"
	"base-gin/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)
// 接口功能的实现, service
func HelloWorld(c *gin.Context)  {
	var user model.SysUser
	sqlTest := global.GVA_DB.Where("username = ?", "admin").First(&user)
	fmt.Println("flag: ",sqlTest.Value)
	fmt.Println(user.NickName)
	fmt.Println(global.GVA_DB.HasTable(&user))
	c.JSON(200, gin.H{
		"message": "Hello world!",
		"info": user.NickName,
	})
	//readFile("../log/text.txt")
}
func readFile(filename string) {
	var (
		err error
		content []byte
	)
	fileObj,err := os.Open(filename)
	if err != nil {
		fmt.Println("os open error:",err)
		return
	}
	defer fileObj.Close()
	content,err = ioutil.ReadAll(fileObj)
	if err != nil {
		fmt.Println("ioutil.ReadAll error:",err)
		return
	}
	fmt.Println(string(content))
}