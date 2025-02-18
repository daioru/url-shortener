package service

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Service) RedirectURL(c *gin.Context) {
	shortURL := c.Param("short")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	originalURL, err := s.repo.GetOriginalURL(ctx, shortURL)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, originalURL)
}
