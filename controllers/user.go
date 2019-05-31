package controllers

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"github.com/trensy/blog/forms"
	"github.com/trensy/blog/models"
	"log"
)

func Login(c *gin.Context) {
	c.HTML(200, "hello.html", pongo2.Context{"name":"world"})
}

func DoLogin(c *gin.Context) {
	var login forms.Login

	if loginerr := c.ShouldBind(&login);loginerr == nil {
		username :=login.Username
		pwd := login.Pwd
		isGet, err := models.CheckLogin(username, pwd)

		if isGet {
			log.Println("is login")
		}else {
			log.Println("err",err)
		}
	}else{
		log.Println("loginerr",loginerr)
	}
    //username := c.Query("username")
	//pwd := c.Query("pwd")

	c.JSON(200, gin.H{"hello":"world", "username":gin.H{"wang":"kaihui"}})
}