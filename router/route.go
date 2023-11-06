package router

import (
	"github.com/1Nelsonel/Gin-Crud/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine)  {
	r.GET("/", controller.GetBooks )
	r.POST("/CreateBook/", controller.CreateBook)
	r.PUT("/UpdateBook/", controller.UpdateBook)
	r.DELETE("/DeleteBook/", controller.DeleteBook)	
}