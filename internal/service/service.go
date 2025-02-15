package service

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type IRepository interface {
	GetOriginalURL(ctx context.Context, shortURL string) (string, error)
	SaveShortURL(ctx context.Context, originalURL, shortURL string) (int64, error)
}

type Service struct {
	repo IRepository
}

func NewService(db *sqlx.DB) *Service {
	return &Service{repo: NewRepository(db)}
}
