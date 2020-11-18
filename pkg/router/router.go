package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	v1 "restaurantManageAPI/pkg/api/v1"
	"restaurantManageAPI/pkg/field"
	"restaurantManageAPI/pkg/middleware"
)

var Router *gin.Engine

func init() {
	Router = gin.New()
	Router.Use(gin.Logger(), gin.Recovery())
	Router.Use(middleware.NotFoundHander())
	gin.SetMode(os.Getenv("GIN_MODE"))

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("NameValidator", field.NameValidator)
	}

	dishV1 := Router.Group("/v1/menu")
	dishV1.POST("/dish", v1.AddDish)
	dishV1.GET("/dish/:id", v1.GetDish)
	dishV1.PUT("/dish/:id", v1.UpdateDish)
	dishV1.DELETE("/dish/:id", v1.DeleteDish)
	dishV1.GET("/menu", v1.GetMenu)
}
