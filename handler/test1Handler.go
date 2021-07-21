package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"myweb/model"
	"myweb/vo"
	"net/http"
	"time"
)

func HandlerTest1(c *gin.Context) {
	db,err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/activiti?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	defer db.Close()
	//测试参数绑定
	var login vo.Login
	if err := c.Bind(&login) ;err != nil {
		log.Println(err)
	}
	cCp := c.Copy()
	go func() {
		log.Println(login)
		log.Println(cCp.Request.URL.Path)
	}()

	u2 := model.User{"topgoer.com", "女", "足球",time.Now()}
	u1 := model.User{"枯藤","男","篮球",time.Now()}
	// 创建记录
	result := db.Create(&u1)
	if result.Error != nil{
		log.Println("插入数据错误：{}",result.Error)
	}
	result = db.Create(&u2)
	if result.Error != nil{
		log.Println("插入数据错误：{}",result.Error)
	}
	log.Println(result.Error)
	c.JSON(http.StatusOK, gin.H{"context":"hello ,world"})
}

func HandlerTest2(c *gin.Context){
	c.JSON(http.StatusOK,"hello go")
}
