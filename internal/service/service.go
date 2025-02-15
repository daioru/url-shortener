package service

import "github.com/jmoiron/sqlx"

type Service struct {
	DB *sqlx.DB
}

func NewService(db *sqlx.DB) *Service {
	return &Service{DB: db}
}
