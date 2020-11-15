package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	dishField "restaurantManageAPI/pkg/field"
	dishModel "restaurantManageAPI/pkg/model"
	"restaurantManageAPI/pkg/util/response"
)

var code int
var err error
var ok bool

func GetDish(context *gin.Context) {
	getDishRequest := new(dishField.GetDishRequest)
	getDishRequest.Id = context.Param("id")
	dish := dishModel.GetDish(getDishRequest.Id)
	getDishResponse := dishField.GetDishResponse{
		Id: dish.Id,
		Name: dish.Name,
	}
	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg": "OK",
		"data": getDishResponse,
	})
}

func AddDish(context *gin.Context) {
	var addDishRequest dishField.AddDishRequest

	if err = context.ShouldBindJSON(&addDishRequest); err != nil {
		log.Println(err)
		code = response.ResponseInvalidParams
		context.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  err.Error(),
		})
	} else {
		dish := dishModel.Dish{
			Name:        addDishRequest.Name,
			Price:       addDishRequest.Price,
			Description: addDishRequest.Description,
			WayToCook:   addDishRequest.WayToCook,
			Cost:        addDishRequest.Cost,
		}
		if err, ok = dish.Create(); !ok {
			code = response.DBError
			context.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  err.Error(),
			})
		} else {
			response.EmptySuccessResp(context)
		}
	}
}

func UpdateDish(context *gin.Context) {

}

func DeleteDish(context *gin.Context) {

}

func GetMenu(context *gin.Context) {

}
