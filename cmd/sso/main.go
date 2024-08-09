package main

import (
	"fmt"
	"log/slog"
	"os"
	"sso/internal/app"
	"sso/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	log := setupLogger(cfg.Env)

	log.Info("Starting application",
		slog.String("env", cfg.Env),
		slog.Any("env", cfg),
		slog.Int("port", cfg.GRPC.Port),
	)

	application := app.New(log, cfg.GRPC.port, cfg.StoragePath, cfg.TokenTTL)

	application.GRPCSrv.MustRun()

	log.Debug("Debug Message")
	log.Error("error Message")
	log.Warn("warn Message")
}
func setupLogger(env string) *slog.Logger { //сложно чот
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
