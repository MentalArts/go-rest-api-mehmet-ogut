package handlers

import (
	"mentalartsapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new book
// @Description Add a new book with title, author, and other details
// @Tags Books
// @Accept json
// @Produce json
// @Param book body models.Book true "Book object"
// @Success 201 {object} models.Book
// @Router /api/v1/books [post]
func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, book)
}

// @Summary Get all books
// @Description Get the list of all books
// @Tags Books
// @Produce json
// @Success 200 {array} models.Book
// @Router /api/v1/books [get]
func GetAllBooks(c *gin.Context) {
	var books []models.Book

	if err := db.Preload("Author").Preload("Reviews").Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

// @Summary Get a book by ID
// @Description Get details of a book by its ID
// @Tags Books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Router /api/v1/books/{id} [get]
func GetBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := db.Preload("Author").Preload("Reviews").First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// @Summary Update a book
// @Description Update a book’s details
// @Tags Books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body models.Book true "Updated book object"
// @Success 200 {object} models.Book
// @Router /api/v1/books/{id} [put]
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := db.First(&book, id).Error; err != nil {
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

// @Summary Delete a book
// @Description Delete a book by ID
// @Tags Books
// @Param id path int true "Book ID"
// @Success 200 {object} gin.H
// @Router /api/v1/books/{id} [delete]
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book

	if err := db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	db.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"msg": "Book deleted"})
}
