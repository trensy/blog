package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/trensy/blog/utils/Base"
)

type Route struct {
	*Base
}

func InitRouter() *gin.Engine {
	r := gin.New()

}
