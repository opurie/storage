package main

import (
	"github.com/gin-gonic/gin"
)


var users = []User{
	{ID: 1, Username: "john_doe", Email: "john@example.com"},
	{ID: 2, Username: "jane_doe", Email: "jane@example.com"},
	{ID: 3, Username: "alice", Email: "alice@example.com"},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(200, users)
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.Run("localhost:8080")
}