package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*.html")

	router.GET("/", func(c *gin.Context) {
		html := template.Must(template.ParseFiles("templates/layouts/base.html", "templates/contents/index.html"))
		router.SetHTMLTemplate(html)
		c.HTML(http.StatusOK, "base.html", gin.H{})
	})

	router.Run()
}
