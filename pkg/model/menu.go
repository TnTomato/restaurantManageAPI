package model

type MenuType struct {
	BaseModel

	Name string `gorm:"comment:'Name of the type'" json:"name"`
}

type Menu struct {
	BaseModel

	TypeId int `gorm:"comment:'id of menu_type'" json:"type_id"`
	DishId int `gorm:"comment:'id of dish'" json:"dish_id"`
	Order int `gorm:"comment:'The rank of the dish under the same type'" json:"order"`
}

type Dish struct {
	BaseModel

	Name        string `gorm:"unique;comment:'Name of the dish'" json:"name"`
	Price       uint16 `gorm:"comment:'Price of the dish'" json:"price"`
	Description string `gorm:"comment:'Description of the dish'" json:"description"`
	WayToCook   string `gorm:"comment:'How to cook'" json:"way_to_cook"`
	Cost        uint16 `gorm:"comment:'The cost price of the dish'" json:"cost"`
}

func GetDish(id int) (dish Dish) {
	//TODO: https://gorm.io/docs/query.html
	DB.First(&dish, id)
	return
}

func (d *Dish) Create() (error, bool) {
	// TODO: The auto-increase primary key keep increasing after db error happen with no new row inserted
	result := DB.Create(d)
	if result.Error != nil {
		return result.Error, false
	} else {
		return nil, true
	}
}

func (d *Dish) String() string {
	return d.Name
}
