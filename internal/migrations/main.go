package main

import (
	"context"
	"embed"
	"time"

	"github.com/daioru/url-shortener/internal/config"
	"github.com/daioru/url-shortener/internal/pkg/db"
	
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	conn, err := db.ConnectDB(&cfg.DB)
	if err != nil {
		log.Fatal().Err(err).Msg("sql.Open() error")
	}
	defer conn.Close()

	goose.SetBaseFS(embedMigrations)

	const cmd = "up"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err = goose.RunContext(ctx, cmd, conn.DB, "migrations")
	if err != nil {
		log.Fatal().Err(err).Msg("goose.Status() error")
	}
}
