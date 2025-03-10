package handlers

import (
	"mentalartsapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new review
// @Description Add a review to a book
// @Tags Reviews
// @Accept json
// @Produce json
// @Param review body models.Review true "Review object"
// @Success 201 {object} models.Review
// @Router /api/v1/books/{id}/reviews [post]
func CreateReview(c *gin.Context) {
	var review models.Review

	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, review)
}

// @Summary Get all reviews for a book
// @Description Get all reviews associated with a specific book
// @Tags Reviews
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {array} models.Review
// @Router /api/v1/books/{id}/reviews [get]
func GetReviews(c *gin.Context) {
	bookID := c.Param("id")
	var reviews []models.Review

	if err := db.Where("book_id = ?", bookID).Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reviews)
}
