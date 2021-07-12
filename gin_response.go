package gtils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func FailResponse(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{
		"code": code,
		"msg":  err.Error(),
		"data": "",
	})
}

func SucResponse(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  msg,
		"data": data,
	})
}
