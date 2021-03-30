package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	godotenvErr := godotenv.Load()
	if godotenvErr != nil {
		log.Panicln("Loading .env file error", godotenvErr)
	}
	dbUser := os.Getenv("DATABASE_USERNAME")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbHost := os.Getenv("DATABASE_HOST")
	dbSchema := os.Getenv("DATABASE_SCHEMA")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True", dbUser, dbPassword, dbHost, dbSchema)
	log.Println("Database Info:", dsn)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("Database connection error")
	}
	sqlDB, dberr := DB.DB()

	if dberr != nil {
		log.Panicln("Invalid database instance")
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)
}
