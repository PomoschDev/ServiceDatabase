package main

import (
	"DababaseService/iternal/app"
	"DababaseService/iternal/lib/logger/handlers/slogpretty"
	"DababaseService/iternal/lib/logger/sl"
	"DababaseService/iternal/migrator"
	"DababaseService/pkg/config"
	"DababaseService/pkg/database"
	"DababaseService/pkg/logger"
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}

func setupLogger(env string) *slog.Logger {
	log := new(slog.Logger)

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelInfo,
			}),
		)
	}

	return log
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//Инициализация конфигурации
	cfg := config.MustLoad()

	//Инициализация логгера
	log := setupLogger(cfg.Env)

	logger.Log = log

	log.Info("Starting app", slog.Any("cfg", cfg))

	//Инициализация БД
	db, err := database.Init(cfg)
	if err != nil {
		panic(any(fmt.Errorf("%w", err)))
	}

	//Миграции
	{
		if cfg.Database.Migrations {
			err := migrator.Migrations(db)
			if err != nil {
				log.Error("%w", sl.Err(err))
			}
		}
	}

	//Инициализация приложения
	application := app.New(log, cfg.GRPC.Port)

	//Правильное завершение сервиса
	{
		wait := time.Second * 15

		// Запуск сервера в отдельном потоке
		go application.GRPCSrv.MustRun()

		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

		<-c

		_, cancel := context.WithTimeout(context.Background(), wait)
		defer cancel()
		application.GRPCSrv.Stop()
		database.CloseConnection()
		os.Exit(0)
	}
}
