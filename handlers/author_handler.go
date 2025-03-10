package handlers

import (
	"mentalartsapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new author
// @Description Create a new author with name and email
// @Tags Authors
// @Accept json
// @Produce json
// @Param author body models.Author true "Author object"
// @Success 201 {object} models.Author
// @Router /api/v1/authors [post]
func CreateAuthor(c *gin.Context) {
	var author models.Author

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	if err := db.Create(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, author)
}

// @Summary Get all authors
// @Description Get the list of all authors
// @Tags Authors
// @Produce json
// @Success 200 {array} models.Author
// @Router /api/v1/authors [get]
func GetAllAuthors(c *gin.Context) {
	var authors []models.Author

	if err := db.Find(&authors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, authors)
}

// @Summary Get a single author
// @Description Get an author by ID
// @Tags Authors
// @Produce json
// @Param id path int true "Author ID"
// @Success 200 {object} models.Author
// @Router /api/v1/authors/{id} [get]
func GetAuthor(c *gin.Context) {
	id := c.Param("id")
	var author models.Author

	if err := db.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": "author not found"})
		return
	}

	c.JSON(http.StatusOK, author)
}

// @Summary Update an author
// @Description Update an author's details
// @Tags Authors
// @Accept json
// @Produce json
// @Param id path int true "Author ID"
// @Param author body models.Author true "Updated author object"
// @Success 200 {object} models.Author
// @Router /api/v1/authors/{id} [put]
func UpdateAuthor(c *gin.Context) {
	id := c.Param("id")
	var author models.Author

	if err := db.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": "author not found"})
		return
	}

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
		return
	}

	if err := db.Save(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, author)
}

// @Summary Delete an author
// @Description Delete an author by ID
// @Tags Authors
// @Param id path int true "Author ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/authors/{id} [delete]
func DeleteAuthor(c *gin.Context) {
	id := c.Param("id")
	var author models.Author

	if err := db.First(&author, id).Error; err != nil {
		c.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
		return
	}

	if err := db.Delete(&author).Error; err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"msg": "Deleted successfully"})
}
