package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

// 設定ファイル読み込み環境変数
const targetEnvName = "GO_ENV"

// main entry point
func main() {

	// 設定ファイルの読み込み
	envLoad()

	// ルーティング
	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		// dbからユーザー全件取得
		db := dBConnect()
		defer db.Close()
		users := []user{}
		db.Find(&users)

		c.JSON(200, gin.H{
			"data": users,
		})
	})
	r.Run(":8080")
}

// user model
type user struct {
	Id   int    `json:id`
	Name string `json:"name"`
}

// dBConnect コネクションの取得
func dBConnect() *gorm.DB {

	// 接続文字列
	connect := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PROTOCOL"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB"))

	db, err := gorm.Open(os.Getenv("DBMS"), connect)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// envLoad env load
func envLoad() {
	// 環境未設定の場合はローカルを設定
	if "" == os.Getenv(targetEnvName) {
		_ = os.Setenv(targetEnvName, "local")
	}
	err := godotenv.Load(fmt.Sprintf("../envfiles/%s.env", os.Getenv(targetEnvName)))
	if err != nil {
		log.Fatalf("Error loading env target env is %s", os.Getenv(targetEnvName))
	}
}
