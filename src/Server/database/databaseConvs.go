package database

import (
	"fmt"
	"log"

	"github.com/MainFoldFunc/Web-sockets-time/src/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBC *gorm.DB

func ConnDatabaseConv() {
	var err error
	DBC, err = gorm.Open(sqlite.Open("Convs.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to a database")
	}
	DBC.AutoMigrate(&models.ChatReqest{})
	fmt.Println("Sucesfully conected to a database")
}
