package main

import (
	"log"
	"mentalartsapi/handlers"
	"mentalartsapi/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

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

	// ✅ **Swagger UI'yi aktif et**
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// ✅ **TÜM ROUTE'LARI `/api/v1` ALTINA GRUPLADIK**
	api := router.Group("/api/v1")
	{
		// **Authors**
		api.POST("/authors", handlers.CreateAuthor)
		api.GET("/authors", handlers.GetAllAuthors)
		api.GET("/authors/:id", handlers.GetAuthor)
		api.PUT("/authors/:id", handlers.UpdateAuthor)
		api.DELETE("/authors/:id", handlers.DeleteAuthor)

		// **Books**
		api.POST("/books", handlers.CreateBook)
		api.GET("/books", handlers.GetAllBooks)
		api.GET("/books/:id", handlers.GetBook)
		api.PUT("/books/:id", handlers.UpdateBook)
		api.DELETE("/books/:id", handlers.DeleteBook)

		// **Reviews**
		api.GET("/books/:id/reviews", handlers.GetReviews)
		api.POST("/books/:id/reviews", handlers.CreateReview)
	}

	log.Println("🚀 Server listening on port 8000...")
	router.Run(":8000")
}
