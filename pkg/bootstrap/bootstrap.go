package bootstrap

import (
	"fmt"
	"github.com/joho/godotenv"
	"gocourse/internal/courses"
	"gocourse/internal/users"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func InitLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}
func DBConnection() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	mysqlPort := os.Getenv("MYSQL_PORT")
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlUser,
		mysqlPassword,
		"127.0.0.1",
		mysqlPort,
		mysqlDatabase,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if os.Getenv("MYSQL_DEBUG") == "True" {
		db = db.Debug()
	}
	if os.Getenv("DATABASE_MIGRATE") == "True" {
		if err := db.AutoMigrate(&users.User{}); err != nil {
			return nil, err
		}
		if err := db.AutoMigrate(&courses.Course{}); err != nil {
			return nil, err
		}
	}
	return db, nil
}
