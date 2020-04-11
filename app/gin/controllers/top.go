package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Top トップページの処理
func Top(c *gin.Context) {
	c.HTML(http.StatusOK, "top.html", gin.H{})
}
