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
	sqldb, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", "username", "password", "database_name"))
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
<<<<<<< HEAD
}
=======
}
>>>>>>> 0efcaa096d1f19ab9d8fdf4e9777c2f59e7bdc78
