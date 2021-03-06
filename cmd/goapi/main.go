package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/sotanodroid/GO_API/configs"
	"github.com/sotanodroid/GO_API/pkg/api"
	"github.com/sotanodroid/GO_API/pkg/models"

	"github.com/go-kit/kit/log/level"
	"github.com/jackc/pgx/v4"

	"github.com/go-kit/kit/log"
	"github.com/joeshaw/envdecode"
)

func init() {
	var logger log.Logger
	if err := godotenv.Load(); err != nil {
		level.Info(logger).Log(err)
	}
}

func main() {
	var logger log.Logger

	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(
			logger, "service", "bookStore",
			"time", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "Starting App...")
	defer level.Info(logger).Log("msg", "App stopped")

	var cfg configs.Config
	if err := envdecode.Decode(&cfg); err != nil {
		level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}

	var db *pgx.Conn
	ctx := context.Background()

	{
		var err error

		db, err = pgx.Connect(ctx, cfg.DbSource)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
	}

	var srv api.Service

	{
		repository := models.NewRepo(db, logger)
		srv = api.NewService(repository, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := api.MakeEndpoints(srv)

	go func() {
		handler := api.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(":"+cfg.Port, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
