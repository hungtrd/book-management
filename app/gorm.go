package app

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/revel/revel"
)

// DB object
var DB *gorm.DB

// DB initialization
func initDB() {
	dbInfo, _ := revel.Config.String("db.info")
	db, err := gorm.Open("mysql", dbInfo)
	if err != nil {
		log.Panicf("Failed gorm.Open: %v\n", err)
	}

	DB = db
}
