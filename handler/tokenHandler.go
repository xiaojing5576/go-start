package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"myweb/model"
	"net/http"
)

func LoginPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := model.QueryUserWithParam(username, password)
	if &user != nil {
		session := sessions.Default(c)
		session.Set("loginUser", username)
		//在gin框架中，无论是Set一个session，还是Delete一个session，都要调用Save()方法进行保存
		session.Save()

		c.JSON(http.StatusOK, gin.H{"result": "login succ"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "user not exsit!"})
	}
}
