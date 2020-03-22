package main

import (
	"github.com/gin-gonic/gin"
	// controller の 処理はモジュール化
	// https://qiita.com/tkj06/items/a5f79417935100045650
	"github.com/kantaroso/kanta-workspace/controllers"
)

func main() {
	router := gin.Default()
	// gin teemplateの使い方
	// https://qiita.com/lanevok/items/dbf591a3916070fcba0d
	router.LoadHTMLGlob("templates/**/*.html")

	// 外部のパッケージの何かを呼び出す時は文頭を大文字にする
	// https://qiita.com/ko-watanabe/items/875085780d2ad72fe6af
	router.GET("/", func(c *gin.Context) { controllers.Top(c, router) })

	router.Run()
}
