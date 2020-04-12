package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kantaroso/kanta-workspace/lib/accesslog"
	"net/http"
)

// Top トップページの処理
func Top(c *gin.Context) {
	accesslog.Register(c.Request)
	accessCount := accesslog.GetAccessCount()
	c.HTML(http.StatusOK, "top.html", gin.H{
		"pv": accessCount,
	})
}
