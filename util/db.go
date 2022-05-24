package util

import (
	"demo01/model"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

func InitMySql() (*gorm.DB,error){
	dsn := "root:root@tcp(127.0.0.1:3306)/demo?charset=utf8&parseTime=True&loc=Local"
	Db,err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	Db.SingularTable(true)
	Db.AutoMigrate(model.DemoOrder{})
	return Db, Db.DB().Ping()
}

func Close() {
	Db.Close()
}