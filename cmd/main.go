package main

import (
	"fmt"
	"log"

	"github.com/daioru/url-shortener/internal/service"

	"github.com/daioru/url-shortener/internal/config"
	"github.com/daioru/url-shortener/internal/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {
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

	service := service.NewService(db)

	r := gin.Default()

	r.POST("/shorten", service.ShortenURL)
	r.GET("/:short", service.RedirectURL)

	fmt.Println("Server running on: 8080")
	log.Fatal(r.Run(":8080"))
}
