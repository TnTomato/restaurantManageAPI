package field

import (
	"github.com/go-playground/validator/v10"
)

type GetDishRequest struct {
	Id int `json:"id" binding:"required"`
}

type GetDishResponse struct {
	Id int `json:"id"`

}

type AddDishRequest struct {
	Name        string `json:"name" binding:"required,NameValidator"`
	Price       uint16 `json:"price" binding:"required"`
	Description string `json:"description"`
	WayToCook   string `json:"way_to_cook"`
	Cost        uint16 `json:"cost"`
}

func NameValidator(fl validator.FieldLevel) bool {
	if name, ok := fl.Field().Interface().(string); ok {
		if name == "什么哩个东西" {
			return false
		}
	}
	return true
}
