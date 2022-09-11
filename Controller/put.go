/**
 *@Description
 *@author          liyajun
 *@create          2022-09-09 16:59
 */

package Controller

import (
	"github.com/gin-gonic/gin"
	"list/DAO"
	"list/Models"
)

func UpdateTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(200, gin.H{
			"error": "id not found",
		})
		return
	}
	list, err := Models.GetById(id)
	if err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = c.BindJSON(&list)
	if err != nil {
		return
	}
	if err = DAO.DB.Save(&list).Error; err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, &list)
	}
}
