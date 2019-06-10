package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *gorm.DB

func init() {
	err := godotenv.Load("./server/.env")

	var username string
	var password string
	var dbname string
	var dbhost string

	if err != nil {
		username = os.Getenv("PGUSER")
		password = os.Getenv("PGPASSWORD")
		dbname = os.Getenv("PGDATABASE")
		dbhost = os.Getenv("PGHOST")
	} else {
		username = os.Getenv("db_user")
		password = os.Getenv("db_password")
		dbname = os.Getenv("db_name")
		dbhost = os.Getenv("db_host")
	}

	uri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbhost, username, dbname, password)
	fmt.Println(uri)

	conn, err := gorm.Open("postgres", uri)
	if err != nil {
		log.Fatal(err)
	}

	db = conn
	db.AutoMigrate(&Account{}, &Task{})

	// Seed database with admin account
	adminUsername := os.Getenv("accounts_admin_username")
	adminPassword := os.Getenv("accounts_admin_password")
	var admin Account
	db.FirstOrInit(&admin, map[string]interface{}{
		"username": adminUsername,
		"password": adminPassword,
		"token": "",
	})
}

func GetDB() *gorm.DB {
	return db
}