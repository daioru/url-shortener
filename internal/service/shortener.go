package service

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ShortenRequest описывает тело запроса для создания короткого URL
type ShortenRequest struct {
	URL string `json:"url" binding:"required"`
}

// ShortenResponse описывает успешный ответ
type ShortenResponse struct {
	ShortURL string `json:"short_url"`
	ID       int    `json:"id"`
}

// ErrorResponse описывает ошибочный ответ
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"` // Оставляем пустым, если нет доп. сообщения
}

// @Summary Создать короткий URL
// @Description Принимает оригинальный URL и возвращает сокращённый
// @Accept  json
// @Produce  json
// @Param request body ShortenRequest true "URL для сокращения"
// @Success 200 {object} ShortenResponse "Успешный ответ с сокращённым URL"
// @Failure 400 {object} ErrorResponse "Ошибка валидации"
// @Failure 500 {object} ErrorResponse "Ошибка сервера"
// @Router /shorten [post]
func (s *Service) ShortenURL(c *gin.Context) {
	var req struct {
		URL string `json:"url" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request parameters", "message": err.Error()})
		return
	}

	shortURL := generateShortURL()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	id, err := s.repo.SaveShortURL(ctx, req.URL, shortURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save URL", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"short_url": "http://localhost:8080/" + shortURL, "id": id})
}
