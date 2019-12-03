package main

import (
	"github.com/gin-gonic/gin"
	"site-health-check/modules/site-healthy/controllers"
)

func main() {
	r := gin.Default()

	handler := controllers.SiteHealthyControllerHandler()
	r.LoadHTMLGlob("views/*")
	r.GET("/", handler.Index)
	r.POST("/", handler.Post)

	r.Run(":8080")
}
