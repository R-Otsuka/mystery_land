package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kamio06/mystery_land/apiserver/mylib"
)

type JsonRequest struct {
	Name  string `json:name`
	Time  int    `json:time`
	Success bool   `json:success`
}

func main() {
	r := gin.Default()

	//ここからCorsの設定
	r.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"http://localhost:8080",
		},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"POST",
			"GET",
			"FETCH",
			"OPTIONS",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: false,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

	r.GET("/record", func(c *gin.Context) {
		records := mylib.FindAllRecord()
		c.JSON(http.StatusOK, gin.H{"records" : records})
	})
	r.POST("/record/register", func(c *gin.Context) {
		var json JsonRequest
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		mylib.DataInsert(json.Name,json.Time,json.Success)
		c.JSON(http.StatusOK, gin.H{"name": json.Name, "time": json.Time, "success": json.Success})
	})

	r.Run(":8800")

}
