package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restaurantManageAPI/pkg/util/response"
)

func NotFoundHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
		status := context.Writer.Status()
		if status == 404 {
			context.JSON(http.StatusNotFound, gin.H{
				"code": response.NotFound,
				"msg": "Resource not found, please check the url",
			})
		}
	}
}