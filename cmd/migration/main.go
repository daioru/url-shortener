package main

import (
	"context"
	"flag"
	"time"

	"github.com/daioru/url-shortener/internal/config"
	"github.com/daioru/url-shortener/internal/migrations"
	"github.com/daioru/url-shortener/internal/pkg/db"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
)

func main() {
	var standalone bool
	flag.BoolVar(&standalone, "standalone", false, "Used to connect to postgres when running outside of a container")

	// After declaring all the flags, enable command-line flag parsing:
	flag.Parse()

	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()

	conn, err := db.ConnectDB(&cfg.DB, standalone)
	if err != nil {
		log.Fatal().Err(err).Msg("sql.Open() error")
	}
	defer conn.Close()

	goose.SetBaseFS(migrations.EmbedFS)

	const cmd = "up"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err = goose.RunContext(ctx, cmd, conn.DB, ".")
	if err != nil {
		log.Fatal().Err(err).Msg("goose.Status() error")
	}
}
