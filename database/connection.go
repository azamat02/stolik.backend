package database

import (
	"fmt"
	"github.com/mborders/logmatic"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"stolik.online/models"
)

var DB *gorm.DB
var l = logmatic.NewLogger()

func Connect() {
	DB_HOST, _ := os.LookupEnv("DB_HOST")
	DB_USER, _ := os.LookupEnv("DB_USER")
	DB_PASS, _ := os.LookupEnv("DB_PASS")
	DB_NAME, _ := os.LookupEnv("DB_NAME")
	DB_PORT, _ := os.LookupEnv("DB_PORT")
	DB_SSL, _ := os.LookupEnv("DB_SSL")

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME, DB_SSL)

	connection, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{PrepareStmt: true})

	if err != nil {
		panic("Could not connect to database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
