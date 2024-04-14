package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"strconv"
	"url-short/pkg/logger"
	"url-short/pkg/storage"

	"github.com/joho/godotenv"
)

const _configPath = "configs/config.env"

func must[A any](a A, err error) A {
	if err != nil {
		panic(err)
	}
	return a
}
func log_init() *slog.Logger {
	cfg := logger.NewConfig()
	cfg.Env = os.Getenv("LOGGER")
	return logger.NewLogger(cfg)
}
func main() {
	if err := godotenv.Load(_configPath); err != nil {
		panic(err)
	}
	run()
}
func run() {
	log := log_init()
	log.Debug("logger started")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	dbconfig := storage.PostgresConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     must(strconv.Atoi(os.Getenv("DB_PORT"))),
		DBname:   os.Getenv("DB_NAME"),
	}

	db, err := storage.NewCockroachDB(ctx, dbconfig)
	if err != nil {
		log.Error("SQL database:", "error", err)
		return
	}
	log.Debug("SQL database upted")
	_ = db

	rediscfg := storage.RedisConfig{
		URL:  os.Getenv("REDIS_HOST"),
		Port: must(strconv.Atoi(os.Getenv("REDIS_PORT"))),
	}

	rs, err := storage.NewRedisCli(ctx, rediscfg)
	if err != nil {
		log.Error("Redis: ", "error", err)
		return
	}
	log.Debug("Redis database upted")
	_ = rs

	ch := make(chan struct{})

	go func() {
		defer done(ch)

	}()

	<-ch
	log.Debug("server stopped")
}
func done(ch chan<- struct{}) {
	ch <- struct{}{}
}
