/**
 *@Description
 *@author          liyajun
 *@create          2022-09-11 15:09
 */

package Routers

import (
	"github.com/gin-gonic/gin"
	"list/Controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//加载静态文件
	r.Static("/static", "./static")
	//加载模板文件
	r.LoadHTMLGlob("templates/*")
	//主页
	r.GET("/", Controller.IndexHandler)

	v1grp := r.Group("v1")
	{
		//添加事项
		v1grp.POST("/todo", Controller.CreateTodo)
		//删除事项
		v1grp.DELETE("/todo/:id", Controller.DeleteTodo)
		//修改事项
		v1grp.PUT("/todo/:id", Controller.UpdateTodo)
		//查看事项
		//查看所有
		v1grp.GET("/todo", Controller.GetAllList)
		//查看某一个
		v1grp.GET("/todo/:id", func(c *gin.Context) {

		})
	}
	return r
}
