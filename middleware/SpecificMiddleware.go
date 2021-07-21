package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func SpecificMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("specific before middleware")
		//设置request变量到Context的Key中,通过Get等函数可以取得
		c.Set("request1", "test")
		//发送request之前
		c.Next()

		//发送requst之后

		// 这个c.Write是ResponseWriter,我们可以获得状态等信息
		status := c.Writer.Status()
		fmt.Println("specific after middleware,", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}
