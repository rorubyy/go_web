package main

import (
	"github.com/gin-gonic/gin"
	"web/controller"
	// "github.com/kevinjone25/go_web/controller"
)

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("template/*")
	server.GET("/home", controller.HomePageHandler)
	server.Run()
}
