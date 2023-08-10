package main

import (
	"context"
	"fmt"
	"github.com/maragudk/env"
	"github.com/shawkyelshalawy/Daily_Brief/server"
	"github.com/shawkyelshalawy/Daily_Brief/storage"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// release is set through the linker at build time, generally from a git sha.
// Used for logging and error reporting.
var release string

func main() {
	os.Exit(start())
}

func start() int {
	_ = env.Load()

	logEnv := env.GetStringOrDefault("LOG_ENV", "development")
	log, err := createLogger(logEnv)
	if err != nil {
		fmt.Println("Error setting up the logger:", err)
		return 1
	}
	log = log.With(zap.String("release", release))
	defer func() {
		// If we cannot sync, there's probably something wrong with outputting logs,
		// so we probably cannot write using fmt.Println either. So just ignore the error.
		_ = log.Sync()
	}()

	host := env.GetStringOrDefault("HOST", "localhost")
	port := env.GetIntOrDefault("PORT", 8080)

	s := server.New(server.Options{
		Database: createDatabase(log),
		Host:     host,
		Log:      log,
		Port:     port,
	})

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		if err := s.Start(); err != nil {
			log.Info("Error starting server", zap.Error(err))
			return err
		}
		return nil
	})

	<-ctx.Done()

	eg.Go(func() error {
		if err := s.Stop(); err != nil {
			log.Info("Error stopping server", zap.Error(err))
			return err
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		return 1
	}
	return 0
}

func createLogger(env string) (*zap.Logger, error) {
	switch env {
	case "production":
		return zap.NewProduction()
	case "development":
		return zap.NewDevelopment()
	default:
		return zap.NewNop(), nil
	}
}

func createDatabase(log *zap.Logger) *storage.Database {
	return storage.NewDatabase(storage.NewDatabaseOptions{
		Host:                  env.GetStringOrDefault("DB_HOST", "localhost"),
		Port:                  env.GetIntOrDefault("DB_PORT", 5432),
		User:                  env.GetStringOrDefault("DB_USER", ""),
		Password:              env.GetStringOrDefault("DB_PASSWORD", ""),
		Name:                  env.GetStringOrDefault("DB_NAME", ""),
		MaxOpenConnections:    env.GetIntOrDefault("DB_MAX_OPEN_CONNECTIONS", 10),
		MaxIdleConnections:    env.GetIntOrDefault("DB_MAX_IDLE_CONNECTIONS", 10),
		ConnectionMaxLifetime: env.GetDurationOrDefault("DB_CONNECTION_MAX_LIFETIME", time.Hour),
		Log:                   log,
	})
}
