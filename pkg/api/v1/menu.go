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

	if dish, ok := model.GetDish(getDishRequest.Id); ok {
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
	var ok bool

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
		if err, ok = model.AddDish(&dish); !ok {
			result = gin.H{
				"code": response.DBError,
				"msg": err.Error(),
			}
		} else {
			result = response.SuccessJson
		}
	}

	context.JSON(http.StatusOK, result)
}

func UpdateDish(context *gin.Context) {
	// TODO: optimize model func
	var err error
	var code int
	var updateDishRequest field.UpdateDishRequest
	var dish4Update model.Dish
	var dish4Check model.Dish

	id := context.Param("id")

	if err = context.ShouldBindJSON(&updateDishRequest); err != nil {
		code = response.InvalidParams
		response.CustomResponse(context, code, err.Error())
	} else {
		findNum := model.DB.First(&dish4Update, "id = ?", id).RowsAffected
		if findNum != 1 {
			code = response.InvalidParams
			response.CustomResponse(context, code, "Invalid id, resource not found")
		} else {
			findNum = model.DB.Where("name = ?", updateDishRequest.Name).First(&dish4Check).RowsAffected
			if findNum == 1 {
				code = response.InvalidParams
				response.CustomResponse(context, code, "Duplicated dish name")
			} else {
				dish4Update.Name = updateDishRequest.Name
				dish4Update.Price = updateDishRequest.Price
				dish4Update.Description = updateDishRequest.Description
				dish4Update.WayToCook = updateDishRequest.WayToCook
				dish4Update.Cost = updateDishRequest.Cost
				if err = model.DB.Save(&dish4Update).Error; err != nil {
					code = response.InvalidParams
					response.CustomResponse(context, code, err.Error())
				} else {
					response.Success(context)
				}
			}
		}
	}
}

func DeleteDish(context *gin.Context) {

}

func GetMenu(context *gin.Context) {

}
