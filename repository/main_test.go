package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestMain(t *testing.M) {
	sqldb, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", "root", "8520", "classicmodels"))
	if err != nil {
		log.Fatalf("Unable to connect to test database, exiting: %v", err)
	}
	sqlDB = sqldb
	//Configure Database connection pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	gormDatabase, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Unable to connect to test connection with DB using Gorm, exiting: %v", err)
	}
	gormDB = gormDatabase
	os.Exit(t.Run())
}