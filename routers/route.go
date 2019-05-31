package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/trensy/blog/controllers"
	"gitlab.com/go-box/pongo2gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/assets", "./web/assets")

	//模板引擎
	r.HTMLRender = pongo2gin.Default()
	//log.Print("rout start")
	//route.Log("hello world")
	r.GET("/ping", controllers.Ping)

	r.GET("/login", controllers.Login)

	r.GET("/dologin", controllers.DoLogin)

	return r
}
