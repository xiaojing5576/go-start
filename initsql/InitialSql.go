package initsql

import (
	"github.com/jinzhu/gorm"
	"time"

	//即使用【import _ 包路径】只是引用该包，仅仅是为了调用init()函数，所以无法通过包名来调用包中的其他函数
	_ "github.com/go-sql-driver/mysql"
	"myweb/model"
)

var _db *gorm.DB

func InitSql() {
	_db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/activiti?charset=utf8mb4&parseTime=True&loc=Local")
	_db.DB().SetMaxOpenConns(100)          //设置数据库连接池最大连接数
	_db.DB().SetMaxIdleConns(20)           //设置连接数据库最大连接数
	_db.DB().SetConnMaxLifetime(time.Hour) //一个连接被复用的最长时间
	if err != nil {
		panic(err)
	}
	//defer _db.Close()
	if !_db.HasTable(&model.User{}) {
		if err := _db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&model.User{}).Error; err != nil {
			panic(err)
		}
	}
}

//获取gorm db对象，其他包需要执行数据库查询的时候，只要通过tools.getDB()获取db对象即可。
//不用担心协程并发使用同样的db对象会共用同一个连接，
// db对象在调用他的方法的时候会从数据库连接池中获取新的连接
// 注意：使用连接池技术后，千万不要使用完db后调用db.Close关闭数据库连接，
// 这样会导致整个数据库连接池关闭，导致连接池没有可用的连接

func GetDB() *gorm.DB {
	return _db
}
