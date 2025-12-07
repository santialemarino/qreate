package main

import (
	"context"
	"net/http"
	"qreate/config"

	"qreate/internal/core/adapter/db"
	httpdelivery "qreate/internal/delivery/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func registerHooks(lifecycle fx.Lifecycle, router *gin.Engine, logger *zap.Logger, cfg *config.Settings) {
	addr := ":" + strconv.Itoa(cfg.App.Port)
	srv := &http.Server{Addr: addr, Handler: router}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				logger.Info("Starting HTTP server", zap.String("addr", addr))
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					logger.Fatal("Failed to start HTTP server", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping HTTP server")
			return srv.Shutdown(ctx)
		},
	})
}

func main() {
	app := fx.New(
		fx.Provide(
			// Config
			func() (*config.Settings, error) {
				return config.New(context.Background())
			},
			// Logger
			func(cfg *config.Settings) (*zap.Logger, error) {
				if cfg.Environment == config.EnvProduction {
					return zap.NewProduction()
				}
				return zap.NewDevelopment()
			},
			// Database
			func(cfg *config.Settings) (*db.Database, error) {
				return db.New(context.Background(), cfg)
			},
			// HTTP Client
			func() *http.Client {
				return &http.Client{}
			},
			// Router
			func(db *db.Database) *gin.Engine {
				return httpdelivery.NewRouter(db)
			},
		),
		fx.Invoke(registerHooks),
	)

	app.Run()
}
