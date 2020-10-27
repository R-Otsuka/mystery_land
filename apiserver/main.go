package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// デフォルトのミドルウェアでginのルーターを作成
	// Logger と アプリケーションクラッシュをキャッチするRecoveryミドルウェア を保有しています
	//db接続
	//mylib.Createtable()

	//router := gin.Default()
	//
	//// ルーター設定
	//// ブラウザで「/」 にアクセスしたら「Hello World!」と表示される設定です
	//router.GET("/", func(c *gin.Context) {
	//	c.String(200, "Hello World!")
	//})
	//router.Run()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/sample", func(c *gin.Context) {
		c.HTML(http.StatusOK, "sample.html", gin.H{
			"title": "Main website",
		})
	})
	router.Run()
}
