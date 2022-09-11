/**
 *@Description
 *@author          liyajun
 *@create          2022-09-09 17:00
 */

package Controller

import (
	"github.com/gin-gonic/gin"
	"list/Models"
	"net/http"
)

func GetAllList(c *gin.Context) {

	list, err := Models.GetAllList()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, list)
	}
}
