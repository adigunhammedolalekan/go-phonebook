package models

import (
	"github.com/joho/godotenv"
	"os"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var conn *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dbName := os.Getenv("db_name");
	dbPassword := os.Getenv("db_pass")
	dbUsername := os.Getenv("db_user")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	conString := fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", dbUsername,
		":" , dbPassword, "@tcp(", dbHost, ":", dbPort, ")/", dbName, "?parseTime=true")

	fmt.Println(conString)
	var db, errr = gorm.Open("mysql", conString)
	if errr != nil {
		panic(errr)
	}

	db.Debug().AutoMigrate(&User{}, &Contact{})
	conn = db
}

func GetConn() (*gorm.DB) {
	return conn
}