package controllers

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Top 表紙のページを表示する
func Top(c *gin.Context, router *gin.Engine) {
	// template の layout を利用する
	// https://qiita.com/taizo/items/38e4af19e2f42208a00a
	html := template.Must(template.ParseFiles("templates/layouts/base.html", "templates/contents/index.html"))
	router.SetHTMLTemplate(html)
	c.HTML(http.StatusOK, "base.html", gin.H{})
}
