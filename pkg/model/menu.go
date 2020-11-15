package model

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

	Name        string `gorm:"unique;comment:'Name of the dish'" json:"name"`
	Price       uint16 `gorm:"comment:'Price of the dish'" json:"price"`
	Description string `gorm:"comment:'Description of the dish'" json:"description"`
	WayToCook   string `gorm:"comment:'How to cook'" json:"way_to_cook"`
	Cost        uint16 `gorm:"comment:'The cost price of the dish'" json:"cost"`
}

func GetDish(id string) (dish Dish) {
	//TODO: https://gorm.io/docs/query.html
	DB.Where("id = ?", id).First(&dish)
	return
}

func (d *Dish) Create() (error, bool) {
	tx := DB.Begin()
	if err := tx.Create(d).Error; err != nil {
		tx.Rollback()
		return err, false
	} else {
		tx.Commit()
		return nil, true
	}
}

func (d *Dish) String() string {
	return d.Name
}
