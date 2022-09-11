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

func DeleteTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(200, gin.H{
			"error": "id not found",
		})
		return
	}
	err := Models.DeleteById(id)
	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, nil)
	}
}
