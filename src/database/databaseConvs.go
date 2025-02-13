package database

import (
	"fmt"
	"log"

	"github.com/MainFoldFunc/Web-sockets-time/src/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBCONVS *gorm.DB

func ConnDatabaseConvs() {
	var err error
	DBCONVS, err = gorm.Open(sqlite.Open("Convs.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to a database")
	}

	fmt.Println("Succesfully connected to a database")

	DB.AutoMigrate(&models.Conv{})
	fmt.Println("Database migrated")
}
