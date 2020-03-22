package controllers

import "github.com/gin-gonic/gin"

// Controllers 各エンドポイントの処理のまとめ
type Controllers interface {
	Top(*gin.Context, *gin.Engine)
}
