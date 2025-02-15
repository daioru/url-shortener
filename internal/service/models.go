package service

import "time"

type ShortURLObject struct {
	ID          int64     `db:"id"`
	ShortURL    string    `db:"short"`
	OriginalURL string    `db:"original"`
	Created     time.Time `db:"created_at"`
}
