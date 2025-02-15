package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/shorten", shortenURL)
	r.GET("/:short", redirectURL)

	fmt.Println("Server running on: 8080")
	log.Fatal(r.Run(":8080"))
}

func shortenURL(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Shorten URL endpoint"})
}

func redirectURL(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Redirect URL endpoint"})
}
