package main

import(
  "github.com/gin-gonic/gin"
	"github.com/trensy/blog/models"
	"github.com/trensy/blog/routers"
	"net/http"
	"time"
)

func main() {
	gin.SetMode(gin.DebugMode)
	//gin.ForceConsoleColor()
	routeInit := routers.InitRouter()
	addr := ":8080"


	models.Init()

	server := &http.Server{
		Addr:addr,
		Handler:routeInit,
		ReadTimeout:10 * time.Second,
		WriteTimeout:10 * time.Second,
		MaxHeaderBytes:1<<20,
	}
	server.ListenAndServe()
}

