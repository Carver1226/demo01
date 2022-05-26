package util

import (
	"demo01/model"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

func InitMySql() *gorm.DB{
	dsn := "root:root@tcp(127.0.0.1:3306)/demo?charset=utf8&parseTime=True&loc=Local"
	Db,err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	Db.SingularTable(true)
	Db.AutoMigrate(model.DemoOrder{})
	return Db
}