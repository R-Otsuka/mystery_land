package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamio06/mystery_land/apiserver/mylib"
)
type JsonRequest struct {
	Name  string `json:"name"`
	Time  string    `json:"time"`
	Success bool   `json:"success"`
}

func main() {
	//time,name,時間が欲しい
	// デフォルトのミドルウェアでginのルーターを作成
	// Logger と アプリケーションクラッシュをキャッチするRecoveryミドルウェア を保有しています
	//mylib.Createtable()
	//mylib.Createtable()
	router := gin.Default()

	//getで歴代レコードを受け取る
	//mysqlにアクセスして取得したデータを返す
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})
	router.POST("/record/register", func(c *gin.Context) {
		var json JsonRequest

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//createtableかinserttableに↓↓のhash形式の物を引数として渡してあげる。
		mylib.Createtable(json.Name,json.Time,json.Success)

		c.JSON(http.StatusOK, gin.H{"name": json.Name, "time": json.Time, "success": json.Success})
	})

	router.Run(":8800")

}
