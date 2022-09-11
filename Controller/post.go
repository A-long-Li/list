/**
 *@Description
 *@author          liyajun
 *@create          2022-09-09 16:59
 */

package Controller

import (
	"github.com/gin-gonic/gin"
	"list/Models"
)

func CreateTodo(c *gin.Context) {
	var todo Models.Todo
	err := c.BindJSON(&todo)
	if err != nil {
		panic(err)
	}
	if err = Models.CreateTodo(&todo); err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, todo)
	}
}
