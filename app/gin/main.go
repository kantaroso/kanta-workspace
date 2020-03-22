package main

import (
	"log"

	// ginを利用する
	// http://sekitaka-1214.hatenablog.com/entry/2016/08/11/153816
	"github.com/gin-gonic/gin"
	// controller の 処理はモジュール化
	// https://qiita.com/tkj06/items/a5f79417935100045650
	"github.com/kantaroso/kanta-workspace/config"
	"github.com/kantaroso/kanta-workspace/controllers"
)

func main() {

	log.Print("settings --------")
	// 構造体をロギングする
	// https://qiita.com/h6591/items/66f4da8cf422087966c5
	log.Printf("%#v", config.GetDatabase())
	log.Print("--------")

	router := gin.Default()
	// gin teemplateの使い方
	// https://qiita.com/lanevok/items/dbf591a3916070fcba0d
	router.LoadHTMLGlob("templates/**/*.html")

	// 外部のパッケージの何かを呼び出す時は文頭を大文字にする
	// https://qiita.com/ko-watanabe/items/875085780d2ad72fe6af
	router.GET("/", func(c *gin.Context) { controllers.Top(c, router) })

	router.Run(":8090")
}
