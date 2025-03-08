package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserNaameChangeDTO struct {
	UserID   uint   `json:"userID"`
	UserName string `json:"userName"`
}

type Response struct {
	Msg string `json:"message"`
}

func main() {

	router := gin.Default()
	router.GET("/ping", handlePing)
	router.GET("/hello/:name", handleHello)
	router.GET("/helloWithPayload", handlehelloWithPayload)
	router.Run(":8000")

}

func handleHello(c *gin.Context) {

	name := c.Query("name")

	var msg string
	if name != "" {
		msg = fmt.Sprintf("Welcome, %s", name)
	} else {
		msg = "Welcome, user"
	}

	c.String(http.StatusOK, msg)
}

func handlePing(c *gin.Context) {

	res := Response{Msg: "/pong"}
	c.JSON(http.StatusOK, res)
}

type DTO struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func handlehelloWithPayload(c *gin.Context) {
	var dto DTO

	err := c.BindJSON(&dto)
	if err != nil {
		c.String(http.StatusBadRequest, "bad Reqquest")

		return
	}

	// validation
	if dto.Name == "" {

		c.String(http.StatusBadRequest, " empty is not accepted ")

		return

	}

	msg := fmt.Sprintf("Welcome, %s %s", dto.Name, dto.Surname)
	c.String(http.StatusOK, msg)

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
