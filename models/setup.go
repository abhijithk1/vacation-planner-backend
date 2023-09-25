package models

import (
	"fmt"
	"log"

	"github.com/abhijithk1/vacation-planner/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init () {
	config := utils.GetAppConfig()

	DbUrl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)
	fmt.Println(DbUrl)

	DB, err = gorm.Open(postgres.Open(DbUrl), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect to database ", config.DBName)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("We are connected to the database ", config.DBName)
	}
}

func MigrateDataBase() {
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Vacation{})
}