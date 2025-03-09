package routes

import (
	"mentalartsapi/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	router.GET("/ping", handlers.HandlePing)
	router.GET("/hello", handlers.HandleHello)
	router.GET("/helloWithPayload", handlers.HandleHelloWithPayload)

	router.POST("/author", handlers.CreateAuthor)
	router.GET("/author", handlers.GetAllAuthors)
	router.GET("/author/:id", handlers.GetAuthor)
	router.PUT("/author/:id", handlers.UpdateAuthor)
	router.DELETE("/author/:id", handlers.DeleteAuthor)
}
