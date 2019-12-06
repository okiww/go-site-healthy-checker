package main

import (
	"github.com/gin-gonic/gin"
	"site-health-check/common/infra/socket"
	"site-health-check/modules/site-healthy/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	handler := controllers.SiteHealthyControllerHandler()
	r.POST("/post", handler.Post)
	r.LoadHTMLGlob("views/*")
	r.GET("/ws", func(c *gin.Context) {
		socket.Wshandler(c.Writer, c.Request)
	})

	r.GET("/", handler.Index)
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	return r
}


func main() {
	r := SetupRouter()
	r.Run(":8080")
}

