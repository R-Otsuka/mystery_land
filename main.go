package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// デフォルトのミドルウェアでginのルーターを作成
	// Logger と アプリケーションクラッシュをキャッチするRecoveryミドルウェア を保有しています
	//mylib.Createtable()
	router := gin.Default()

	// ルーター設定
	// ブラウザで「/」 にアクセスしたら「Hello World!」と表示される設定です
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})
	router.Run()
}