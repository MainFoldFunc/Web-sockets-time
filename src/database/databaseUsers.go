package database

import (
	"fmt"
	"log"

	"github.com/MainFoldFunc/Web-sockets-time/src/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("Users.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to a database")
	}

	fmt.Println("Succesfully connected to a database")

	DB.AutoMigrate(&models.Users{})
	DB.AutoMigrate(&models.ChatReqest{})
	fmt.Println("Database migrated")
}
