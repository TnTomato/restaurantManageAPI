package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var SuccessJson = gin.H{
	"code": 200,
	"msg": "OK",
}

func Success(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": OK,
		"msg":  ResponseMsg(OK),
		"data": nil,
	})
}

func CustomResponse(context *gin.Context, code int, msg string, data ...interface{}) {
	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
