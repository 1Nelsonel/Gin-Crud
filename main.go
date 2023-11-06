package main

import (
	"github.com/1Nelsonel/Gin-Crud/database"
	"github.com/1Nelsonel/Gin-Crud/router"
	"github.com/gin-gonic/gin"
)

func init() {
	database.ConnectDB()
}

func main() {
	// Database connect
	sqlDb, err := database.DBConn.DB()

	if err != nil {
		panic("Error in connection")
	}

	defer sqlDb.Close()

	// Gin 
	r := gin.Default()
	
	// Templates
	r.LoadHTMLGlob("templates/*")
	
	// Initialize routes by calling SetupRoutes
	router.SetupRoutes(r)

	// Server
	r.Run(":3000")
}