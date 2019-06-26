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
	err := godotenv.Load("./api/.env")

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
	db.AutoMigrate(&Task{})

	// Seed database with task
	var task Task
	db.Debug().FirstOrCreate(&task, Task{
		Schedule: "0 30 6 ? * MON,TUE,WED,THU,FRI",
		Type: "fade",
		Data: `{"start_color":{"r":0,"g":-50,"b":-120},"end_color":{"r":255,"g":130,"b":40},"start_intensity":0,"end_intensity":1,"duration":3600000}`,
	})
	db.Debug().FirstOrCreate(&task, Task{
		Schedule: "0 0 8 ? * MON,TUE,WED,THU,FRI",
		Type: "clear",
	})

	fmt.Printf("Initialized database\n")
}

func GetDB() *gorm.DB {
	return db
}