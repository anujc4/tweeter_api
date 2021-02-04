package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// User type
type User struct {
	ID        int64
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func healthCheck(c *gin.Context) {
	c.String(200, "UP")
}

func readMock(c *gin.Context) {
	file, err := os.Open("mock.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to open file"})
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read contents of tile"})
	}
	var users []User
	json.Unmarshal(fileBytes, &users)
	c.JSON(http.StatusOK, users)
}

func main() {
	r := gin.Default()
	r.GET("/", healthCheck)
	r.GET("/simple_health", healthCheck)

	v1 := r.Group("/v1")
	v1.GET("/users", readMock)

	r.Run(":3000")
}
