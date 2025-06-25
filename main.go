package main

import (
	"github.com/gin-gonic/gin"
	"main/app/models"
	// . "main/config"
)

var users = []model.User{
	{ID: 1, Username: "alice", Email: "d"},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(200, users)
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.Run("localhost:8080")
}