package service

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Перенаправление по короткому URL
// @Description Перенаправляет пользователя на оригинальный URL
// @Produce  json
// @Param short path string true "Короткий URL"
// @Success 301
// @Failure 404 {object} map[string]string
// @Router /{short} [get]
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
