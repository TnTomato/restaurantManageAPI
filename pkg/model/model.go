package model

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/mgo.v2/bson"
)

var DB *gorm.DB

type BaseModel struct {
	Id        string `gorm:"primary_key;type:varchar(24);comment:'The unique id'" json:"id"`
	CreatedAt int    `gorm:"comment:'The datetime when it is created'" json:"created_at"`
	UpdatedAt int    `gorm:"comment:'The datetime when it is updated'" json:"updated_at"`
	IsEnable  int8   `gorm:"default:1;comment:'To indicate logically deleted rows'" json:"is_enable"`
}

func (model *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("Id", bson.NewObjectId().Hex())
	scope.SetColumn("CreatedAt", time.Now().Unix())
	return nil
}

func (model *BaseModel) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now().Unix())
	return nil
}

func init() {
	var err error

	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")
	mysqlDB := os.Getenv("MYSQL_DB")
	mysqlCharset := os.Getenv("MYSQL_CHARSET")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDB, mysqlCharset)

	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

	DB.SingularTable(true)

	DB.DropTableIfExists(MenuType{}, Menu{}, Dish{})
	DB.CreateTable(MenuType{}, Menu{}, Dish{})
}
