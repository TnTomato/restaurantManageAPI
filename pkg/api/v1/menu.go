package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"restaurantManageAPI/pkg/field"
	"restaurantManageAPI/pkg/model"
	"restaurantManageAPI/pkg/util/response"
)

func GetDish(context *gin.Context) {
	var result gin.H

	getDishRequest := new(field.GetDishRequest)
	getDishRequest.Id = context.Param("id")

	if dish, ok := model.FindDishById(getDishRequest.Id); ok {
		getDishResponse := field.GetDishResponse{
			Id:   dish.Id,
			Name: dish.Name,
			Price: dish.Price,
			Description: dish.Description,
			WayToCook: dish.WayToCook,
			Cost: dish.Cost,
		}
		result = gin.H{
			"code": response.OK,
			"msg": response.ResponseMsg(response.OK),
			"data": getDishResponse,
		}
	} else {
		result = gin.H{
			"code": response.NotFound,
			"msg": response.ResponseMsg(response.NotFound),
		}
	}

	context.JSON(http.StatusOK, result)
}

func AddDish(context *gin.Context) {
	var result gin.H
	var err error
	var status int

	var addDishRequest field.AddDishRequest

	if err = context.ShouldBindJSON(&addDishRequest); err != nil {
		result = gin.H{
			"code": response.InvalidParams,
			"msg": err.Error(),
		}
	} else {
		dish := model.Dish{
			Name:        addDishRequest.Name,
			Price:       addDishRequest.Price,
			Description: addDishRequest.Description,
			WayToCook:   addDishRequest.WayToCook,
			Cost:        addDishRequest.Cost,
		}
		
		switch err, status = model.CreateDish(&dish); status {
		case response.OK:
			result = response.SuccessJson
		case response.DuplicatedName:
			result = gin.H{
				"code": status,
				"msg": err.Error(),
			}
		case response.DBError:
			result = gin.H{
				"code": status,
				"msg": err.Error(),
			}
		}
	}

	context.JSON(http.StatusOK, result)
}

func UpdateDish(context *gin.Context) {
	var result gin.H
	var err error
	var status int
	var updateDishRequest field.UpdateDishRequest


	id := context.Param("id")

	if err = context.ShouldBindJSON(&updateDishRequest); err != nil {
		result = gin.H{
			"code": response.InvalidParams,
			"msg": err.Error(),
		}
	} else {
		switch err, status = model.UpdateDishById(id, updateDishRequest); status {
		case response.OK:
			result = response.SuccessJson
		case response.NotFound:
			result = gin.H{
				"code": status,
				"msg": response.ResponseMsg(status),
			}
		case response.InvalidParams:
			result = gin.H{
				"code": status,
				"msg": err.Error(),
			}
		}
	}

	context.JSON(http.StatusOK, result)
}

func DeleteDish(context *gin.Context) {
	var result gin.H
	var err error
	var status int

	deleteDishRequest := new(field.DeleteDishRequest)
	deleteDishRequest.Id = context.Param("id")

	switch err, status = model.DeleteDishById(deleteDishRequest.Id); status {
	case response.OK:
		result = response.SuccessJson
	case response.NotFound:
		result = gin.H{
			"code": status,
			"msg": response.ResponseMsg(status),
		}
	case response.InvalidParams:
		result = gin.H{
			"code": status,
			"msg": err.Error(),
		}
	}

	context.JSON(http.StatusOK, result)
}

func GetMenu(context *gin.Context) {

}
