package mylib

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB上のテーブル、カラムと構造体との関連付けが自動的に行われる
type Test struct {
	ID     int    `gorm:"primary_key;not null"`
	Name   string `gorm:"type:varchar(200);not null"`
	Memo   string `gorm:"type:varchar(400)"`
	Status string `gorm:"type:char(2);not null"`
}

type TouchNumbersRecord struct {
	ID     int    `gorm:"primary_key;not null"`
	Name   string `gorm:"type:varchar(200);not null"`
	Time   string `gorm:"type:varchar(200);not null"`
	Success bool `gorm:"type:bool;not null"`
}

func getGormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "root"
	PROTOCOL := "tcp(localhost:3306)"
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
func insertTest(registerRecord *TouchNumbersRecord) {
	db := getGormConnect()

	// insert文
	db.Create(&registerRecord)
	defer db.Close()
}

// 商品テーブルのレコードを全件取得
func findAllTest() []Test {
	db := getGormConnect()
	var Tests []Test

	// select文
	db.Order("ID asc").Find(&Tests)
	defer db.Close()
	return Tests
}

func Createtable() {
	// Testテーブルにデータを運ぶための構造体を初期化
	var Record = TouchNumbersRecord{
		Name:   "testotsuka",
		Time:   "00:26:10",
		Success: false,
	}

	// 構造体のポインタを渡す
	insertTest(&Record)

	// Testテーブルのレコードを全件取得する
	resultTests := findAllTest()

	// Testテーブルのレコードを全件表示する
	for i := range resultTests {
		fmt.Printf("index: %d, 商品ID: %d, 商品名: %s, メモ: %s, ステータス: %s\n",
			i, resultTests[i].ID, resultTests[i].Name, resultTests[i].Memo, resultTests[i].Status)
	}
}
