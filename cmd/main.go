package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/daioru/url-shortener/internal/config"
	"github.com/daioru/url-shortener/internal/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/shorten", shortenURL)
	r.GET("/:short", redirectURL)

	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	db, err := db.ConnectDB(&cfg.DB)
	if err != nil {
		log.Fatalf("sqlx_Open error: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error testing db connection: %v", err)
	}

    fmt.Println("Server running on: 8080")
	log.Fatal(r.Run(":8080"))
}

func shortenURL(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Shorten URL endpoint"})
}

func redirectURL(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Redirect URL endpoint"})
}
