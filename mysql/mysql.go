package mysql

import (
	"ShowWeb/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitMysql() *gorm.DB {
	dsn := "hjy:admin@tcp(47.98.142.20:3306)/hjytest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("failed to connect database , err: " + err.Error())
	}

	db.AutoMigrate(&models.Menu{}, &models.Plug{}, &models.DownHistory{}, &models.User{})
	db.LogMode(true)

	return db
}

func GetDb() *gorm.DB {
	DB = InitMysql()
	return DB
}
