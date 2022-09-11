/**
 *@Description
 *@author          liyajun
 *@create          2022-09-09 16:56
 */

package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
