package DB

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// func ReturnResponse(mes string, c echo.Context) error {
// 	message := Models.Message{}
// 	json.Unmarshal([]byte(`{"Mess": "`+mes+`"}`), &message) //json文字列→構造体へ  設定したkey沿う(今回の場合はmessに対応)
// 	return c.JSON(http.StatusOK, message)                   //json.Marshall(構造体指定のkeyで値を入れる) 構造体→json文字列の作成
// }

func GormConnect() *gorm.DB {
	var isDev bool = os.Getenv("GO_ENV") == "development" //ここら辺は環境変数からmysqlの接続データを取得してくる
	var user = os.Getenv("DB_USER")
	var pwd = os.Getenv("DB_PWD")
	var host = os.Getenv("DB_HOST")
	var port = os.Getenv("DB_PORT")
	var database = os.Getenv("DB_DATABASE")

	connString := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pwd, host, port, database)
	db, err := gorm.Open("mysql", connString)

	if err != nil {
		panic(err.Error())
	}
	db.LogMode(isDev) //sqlのログを出力してくれる、便利
	return db

	// defer db.Close()
}
