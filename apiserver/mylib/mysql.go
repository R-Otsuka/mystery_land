package mylib

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB上のテーブル、カラムと構造体との関連付けが自動的に行われる
type JsonRequest struct {
	Name  string `json:"name"`
	Time  int    `json:"time"`
	Success bool   `json:"success"`
}
type TouchNumbersRecord struct {
	ID     int    `gorm:"primary_key;not null"`
	Name   string `gorm:"type:varchar(200);not null"`
	Time   int `gorm:"type:int;not null"`
	Success bool `gorm:"type:bool;not null"`
}

func getGormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(mysql:3306)"//コンテナtoコンテナの通信はlocalhostではなくmysql(コンテナ名とするひつようあり)
	DBNAME := "mystery_land"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	// DBエンジンを「InnoDB」に設定
	db.Set("gorm:table_options", "ENGINE=InnoDB")

	// 詳細なログを表示
	db.LogMode(true)

	// 登録するテーブル名を単数形にする（デフォルトは複数形）
	db.SingularTable(true)

	// マイグレーション（テーブルが無い時は自動生成）
	db.AutoMigrate(&TouchNumbersRecord{})

	fmt.Println("db connected: ", &db)
	return db
}

// レコードを追加
func Insert(registerRecord *TouchNumbersRecord) {
	db := getGormConnect()

	// insert文
	db.Create(&registerRecord)
	defer db.Close()
}

// 商品テーブルのレコードを全件取得
func FindAllRecord() []TouchNumbersRecord {
	db := getGormConnect()
	var Records []TouchNumbersRecord

	// select文
	db.Order("ID asc").Find(&Records)
	defer db.Close()
	return Records
}

//送信されて来た記録をdbに保存する。
func DataInsert(name string,time int, success bool) {
	//fmt.Println("======")
	//fmt.Println("名前:",name)
	//fmt.Println("時間:",time)
	//fmt.Println("クリア:",success)
	var Record = TouchNumbersRecord{
		Name:   name,
		Time:   time,
		Success: success,
	}
	// 構造体のポインタを渡す
	Insert(&Record)

	//テーブルのレコードを全件取得する

	// Testテーブルのレコードを全件表示する
	//for i := range resultTests {
	//	fmt.Printf("index: %d, 商品ID: %d, 商品名: %s, メモ: %s, ステータス: %s\n",
	//		i, resultTests[i].ID, resultTests[i].Name, resultTests[i].Memo, resultTests[i].Status)
	//}
}
