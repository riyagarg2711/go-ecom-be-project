package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/riyagarg2711/ecom-api-course/internal/env"
)

func main() {
	ctx := context.Background()

	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "host=localhost port=5433 user=postgres password=postgres dbname=ecom sslmode=disable"),
		},
	}

	//LOGGER

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	//DATABASE
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)

	}
	defer conn.Close(ctx)

	logger.Info("connected to database", "dsn", cfg.db.dsn)

	api := application{
		config: cfg,
		db: conn,
	}

	if err := api.run(api.mount()); err != nil {
		//log.Printf("server has failed to start, err: %s", err)
		slog.Error("server failed to start", "error", err)
		os.Exit(1)

	}


}
