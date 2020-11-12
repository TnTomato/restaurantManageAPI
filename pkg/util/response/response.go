package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func EmptySuccessResp(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": ResponseOK,
		"msg":  ResponseMsg(ResponseOK),
	})
}

func SuccessResp(context *gin.Context, ) {
	context.JSON(http.StatusOK, gin.H{
		"code": ResponseOK,
		"msg":  ResponseMsg(ResponseOK),
		"data": nil,
	})
}
