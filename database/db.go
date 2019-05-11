package database

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/joho/godotenv/autoload"
)

// var DB = make(chan *gorm.DB, 10)
var err error
var DB *gorm.DB

func Connect() {
	// runtime.GOMAXPROCS(5)
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		DbChannel, _ := gorm.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME")+"?charset=utf8&parseTime=True&loc=Local")
	// 		DB <- DbChannel
	// 	}
	// }()
	DB, err = gorm.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME")+"?charset=utf8&parseTime=True&loc=Local")

	// DB.SingularTable(true)
	// defer db.Close()
}
