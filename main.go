package main

import (
	"github.com/gin-gonic/gin"
	"myweb/handler"
	"myweb/initsql"
	"myweb/middleware"
)

func main() {
	//设置模式
	gin.SetMode(gin.ReleaseMode)
	//路由
	router := gin.Default()

	//中间件的顺序也决定了使用范围，use之后的可被覆盖到，之前的不行
	router.Use(middleware.MiddleWare())

	test1Group := router.Group("/test1")
	test1Group.POST("/hello1", handler.HandlerTest1)
	test2Group := router.Group("/test2")
	//还可以指定路由函数进行注册
	test2Group.GET("/hello2",middleware.SpecificMiddleWare() ,handler.HandlerTest2)

	//初始化数据库
	initsql.InitSql()

	router.Run(":8080")

	// 1.全局中间件 router.Use(gin.Logger()) router.Use(gin.Recovery())

	// 2.单路由的中间件，可以加任意多个 router.GET("/benchmark", MyMiddelware(), benchEndpoint)

	// 3.群组路由的中间件 authorized := router.Group("/", MyMiddelware()) // 或者这样用： authorized := router.Group("/") authorized.Use(MyMiddelware()) { authorized.POST("/login", loginEndpoint) }
}
