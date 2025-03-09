package handlers

import (
	"mentalartsapi/database" //
	"mentalartsapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAuthor(c *gin.Context) {
	var author models.Author

	err := c.ShouldBindJSON(&author)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB() // ✅ Artık veritabanı bağlantısını config üzerinden alıyoruz
	result := db.Create(&author)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusCreated, author)
}

func GetAllAuthors(c *gin.Context) {
	var authors []models.Author
	db := database.GetDB() // ✅ Veritabanı bağlantısını al

	if err := db.Find(&authors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, authors)
}
