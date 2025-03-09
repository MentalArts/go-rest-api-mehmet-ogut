package main

import (
	"fmt"
	"log"
	"mentalartsapi/handlers"
	"mentalartsapi/models"

	//	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserNaameChangeDTO struct {
	UserID   uint   `json:"userID"`
	UserName string `json:"userName"`
}

func main() {
	//	pgPassword := os.Getenv("PG_PASSWORD")
	dsn := fmt.Sprintf("host=localhost user=postgres password=123abc dbname=postgres port=5432 sslmode=disable ")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect database: %s", err)
	}
	db.AutoMigrate(&models.Author{})

	router := gin.Default()
	router.GET("/ping", handlers.HandlePing)
	router.GET("/hello/:name", handlers.HandleHello)
	router.GET("/helloWithPayload", handlers.HandlehelloWithPayload)
	router.Run(":8000")

}

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
//)

// func main() {

// 	http.HandleFunc("GET /ping", handelePing)

// 	log.Println("Server listining...")
// 	log.Fatal(http.ListenAndServe(":8000", nil))

// }

// func handelePing(w http.ResponseWriter, r *http.Request) {

// 	res := response{Msg: "pong"}

// 	json.NewEncoder(w).Encode(res)

// 	w.WriteHeader(http.StatusOK)

// 	log.Println("Request rescives")
// }
