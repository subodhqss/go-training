package repository

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

func DBInit() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//
	dsn := "root:qss@2021@/studentmodels"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Unable to connect to Gorm, exiting: %v", err)
	}
	fmt.Println("DB connected")
	gormDB = db
}
