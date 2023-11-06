package controller

import (
	"net/http"
	"github.com/1Nelsonel/Gin-Crud/database"
	"github.com/1Nelsonel/Gin-Crud/model"
	"github.com/gin-gonic/gin"
)

// func GetBooks(c *gin.Context)  {
// 	c.HTML(http.StatusOK, "index.html", c)
// }

// GetBooks lists all books.
func GetBooks(c *gin.Context) {
	db := database.DBConn
    var books []model.Book

    if err := db.Find(&books).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
        return
    }

	// Create a Gin context
	context := make(map[string]interface{})
	
	context["books"] = books
    // c.JSON(http.StatusOK, books)
	c.HTML(http.StatusOK, "index.html", context)
}

// CreateBook creates a new Book.
func CreateBook(c *gin.Context) {
	db := database.DBConn
    var book model.Book

    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := db.Create(&book).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Book"})
        return
    }

    c.JSON(http.StatusCreated, book)
}

// UpdateBook updates an existing Book.
func UpdateBook(c *gin.Context) {
    var book model.Book
    bookID := c.Param("id")

	db := database.DBConn

    if err := db.Where("id = ?", bookID).First(&book).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }

    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db.Save(&book)
    c.JSON(http.StatusOK, book)
}

// DeleteBook deletes an existing Book.
func DeleteBook(c *gin.Context) {
	db := database.DBConn

    bookID := c.Param("id")

    if err := db.Where("id = ?", bookID).Delete(&model.Book{}).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}