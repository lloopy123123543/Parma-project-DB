package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", os.Getenv("pg_host"), os.Getenv("pg_user"), os.Getenv("pg_password"), os.Getenv("pg_dbname"))
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic(err)
	}
	db.AutoMigrate()
	return db
}
