package main

import (
	"log"
	"mentalartsapi/handlers"
	"mentalartsapi/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "mentalartsapi/docs"
)

// @title Bookstore API
// @version 1.0
// @description API for managing books, authors, and reviews
// @host localhost:8000
// @BasePath /api/v1
func main() {
	dsn := "host=localhost user=postgres password=123abc dbname=postgres port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect database: %v", err)
	}

	db.AutoMigrate(&models.Author{}, &models.Book{}, &models.Review{})

	handlers.InitDB(db)

	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	router.POST("/api/v1/authors", handlers.CreateAuthor)
	router.GET("/api/v1/authors", handlers.GetAllAuthors)
	router.GET("/api/v1/authors/:id", handlers.GetAuthor)
	router.PUT("/api/v1/authors/:id", handlers.UpdateAuthor)
	router.DELETE("/api/v1/authors/:id", handlers.DeleteAuthor)

	router.POST("/api/v1/books", handlers.CreateBook)
	router.GET("/api/v1/books", handlers.GetAllBooks)
	router.GET("/api/v1/books/:id", handlers.GetBook)
	router.PUT("/api/v1/books/:id", handlers.UpdateBook)
	router.DELETE("/api/v1/books/:id", handlers.DeleteBook)

	router.GET("/api/v1/books/:id/reviews", handlers.GetReviews)
	router.POST("/api/v1/books/:id/reviews", handlers.CreateReview)

	log.Println("🚀 Server listening on port 8000...")
	router.Run(":8000")
}
