package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kantaroso/kanta-workspace/controllers"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*.html")

	router.GET("/", func(c *gin.Context) { controllers.Top(c, router) })

	router.Run()
}
