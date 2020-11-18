package model

import (
	"errors"
	"restaurantManageAPI/pkg/field"
	"restaurantManageAPI/pkg/util/response"
)

type MenuType struct {
	BaseModel

	Name string `gorm:"comment:'Name of the type'" json:"name"`
}

type Menu struct {
	BaseModel

	TypeId string `gorm:"type:varchar(24);comment:'id of menu_type'" json:"type_id"`
	DishId string `gorm:"type:varchar(24);comment:'id of dish'" json:"dish_id"`
	Order  int8   `gorm:"comment:'The rank of the dish under the same type'" json:"order"`
}

type Dish struct {
	BaseModel

	Name        string `gorm:"comment:'Name of the dish'" json:"name"`
	Price       uint16 `gorm:"comment:'Price of the dish'" json:"price"`
	Description string `gorm:"comment:'Description of the dish'" json:"description"`
	WayToCook   string `gorm:"comment:'How to cook'" json:"way_to_cook"`
	Cost        uint16 `gorm:"comment:'The cost price of the dish'" json:"cost"`
}

func FindDishById(id string) (dish Dish, state bool) {
	rowAffected := DB.Where("id = ?", id).First(&dish).RowsAffected
	if rowAffected == 1 {
		state = true
	} else {
		state = false
	}
	return
}

func CreateDish(d *Dish) (err error, status int) {
	var foundDish Dish
	if rowsAffected := DB.Where(
		"name = ? and is_enable = ?", d.Name, 1).First(
			&foundDish).RowsAffected; rowsAffected >= 1 {
		err = errors.New("duplicated name")
		status = response.DuplicatedName
	} else {
		if err := DB.Create(d).Error; err != nil {
			status = response.DBError
		} else {
			status = response.OK
		}
	}
	return
}

func UpdateDishById(id string, request field.UpdateDishRequest) (err error, status int) {
	var dish4Update Dish
	var dish4Check Dish

	foundNum := DB.First(&dish4Update, "id = ?", id).RowsAffected
	if foundNum != 1 {
		status = response.NotFound
	} else {
		foundNum = DB.Where("name = ? and is_enable = ?", request.Name, 1).First(&dish4Check).RowsAffected
		if foundNum == 1 && dish4Check.Id != id{
			status = response.InvalidParams
			err = errors.New("duplicated name")
		} else {
			dish4Update.Name = request.Name
			dish4Update.Price = request.Price
			dish4Update.Description = request.Description
			dish4Update.WayToCook = request.WayToCook
			dish4Update.Cost = request.Cost
			if err = DB.Save(&dish4Update).Error; err != nil {
				status = response.InvalidParams
			} else {
				status = response.OK
			}
		}
	}
	return
}

func DeleteDishById(id string) (err error, status int) {
	var dish4Update Dish

	foundNum := DB.First(&dish4Update, "id = ? and is_enable = ?", id, 1).RowsAffected
	if foundNum != 1 {
		status = response.NotFound
	} else {
		dish4Update.IsEnable = 0
		if err = DB.Save(&dish4Update).Error; err != nil {
			status = response.InvalidParams
		} else {
			status = response.OK
		}
	}
	return
}
