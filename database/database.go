package database

import (
	"log"

	"github.com/beintil/goregisterjwt/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBCon   *gorm.DB
	DBError error
)

//Connect's database
func DBConnect(connect string) () {
	DBCon, DBError = gorm.Open(mysql.Open(connect), &gorm.Config{})
	if DBError != nil {
		log.Fatal(DBError)
		panic("Connection errors with database")
	}

	log.Print("Connect database")

	DBCon.AutoMigrate(&user.User{})
}

	

