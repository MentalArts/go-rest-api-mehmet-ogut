package handlers

import (
	"fmt"
	"mentalartsapi/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleHello(c *gin.Context) {

	name := c.Query("name")

	var msg string
	if name != "" {
		msg = fmt.Sprintf("Welcome, %s", name)
	} else {
		msg = "Welcome, user"
	}

	c.String(http.StatusOK, msg)
}

func HandlePing(c *gin.Context) {

	res := dto.Response{Msg: "/pong"}
	c.JSON(http.StatusOK, res)
}

func HandlehelloWithPayload(c *gin.Context) {
	var user dto.User

	err := c.BindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "bad Reqquest")

		return
	}

	// validation
	if user.Name == "" {

		c.String(http.StatusBadRequest, " empty is not accepted ")

		return

	}

	msg := fmt.Sprintf("Welcome, %s %s", user.Name, user.Surname)
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
