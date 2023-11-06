package database

import (
	"github.com/1Nelsonel/Gin-Crud/model"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


var DBConn *gorm.DB

func ConnectDB() {
		
	var err error	

	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	log.Info("Connection success")

	db.AutoMigrate(&model.Book{})

	DBConn = db
}