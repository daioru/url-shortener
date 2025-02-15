package service

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	DB *sqlx.DB
}

func NewRepository(db *sqlx.DB) repository {
	return repository{DB: db}
}

func (r repository) GetOriginalURL(ctx context.Context, shortURL string) (string, error) {
	query, args, _ := sq.Select("original").
		PlaceholderFormat(sq.Dollar).
		From("urls").
		Where(sq.Eq{"short": shortURL}).
		ToSql()

	var originalURL string
	err := r.DB.GetContext(ctx, &originalURL, query, args...)
	return originalURL, err
}

func (r repository) SaveShortURL(ctx context.Context, originalURL, shortURL string) (int64, error) {
	query := sq.Insert("urls").
		PlaceholderFormat(sq.Dollar).
		Columns("short", "original").
		Values(shortURL, originalURL).Suffix("RETURNING id").RunWith(r.DB)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return 0, err
	}

	var id int64
	if rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return 0, err
		}

		return id, nil
	} else {
		return 0, sql.ErrNoRows
	}
}
