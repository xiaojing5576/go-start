package initsql

import (
	"github.com/jinzhu/gorm"
	//即使用【import _ 包路径】只是引用该包，仅仅是为了调用init()函数，所以无法通过包名来调用包中的其他函数
	_ "github.com/go-sql-driver/mysql"
	"myweb/model"
)

func InitSql()  {
	db,err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/activiti?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	defer db.Close()
	if !db.HasTable(&model.User{}){
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.User{}).Error; err != nil {
			panic(err)
		}
	}
}
