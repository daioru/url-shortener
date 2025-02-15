package service

import (
	"net/http"

	"github.com/Masterminds/squirrel"
	sq "github.com/Masterminds/squirrel"
	"github.com/gin-gonic/gin"
)

func (s *Service) RedirectURL(c *gin.Context) {
	short := c.Param("short")

	query, args, _ := sq.Select("original").
		PlaceholderFormat(sq.Dollar).
		From("urls").
		Where(squirrel.Eq{"short": short}).
		ToSql()

	var originalURL string
	err := s.DB.Get(&originalURL, query, args...)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, originalURL)
}
