package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetBlogs(c *gin.Context)  {
	c.JSON(http.StatusOK, c)
}