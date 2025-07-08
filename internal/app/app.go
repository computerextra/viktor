package app

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/computerextra/viktor/db"
	"github.com/computerextra/viktor/internal/middleware"
	"github.com/computerextra/viktor/internal/util/flash"
	"github.com/rs/cors"
)

type App struct {
	config   Config
	frontend fs.FS
	logger   *slog.Logger
	db       *db.PrismaClient
}

func New(logger *slog.Logger, config Config, frontend fs.FS) (*App, error) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &App{
		config:   config,
		logger:   logger,
		frontend: frontend,
		db:       client,
	}, nil
}

func (a *App) Start(ctx context.Context) error {
	router, err := a.loadRoutes()
	if err != nil {
		return fmt.Errorf("failed when loading routes: %w", err)
	}

	middlewares := middleware.Chain(
		middleware.Logging(a.logger),
		flash.Middleware,
	)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete},
	})

	port := getPort(3000)
	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        c.Handler(middlewares(router)),
		MaxHeaderBytes: 1 << 20, // Max Header size (e.g. 1MB)
	}

	errCh := make(chan error, 1)

	go func() {
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- fmt.Errorf("failed to listen and serve: %w", err)
		}
		close(errCh)
	}()

	a.logger.Info("server running", slog.Int("port", port))

	select {
	case <-ctx.Done():
		break
	case err := <-errCh:
		return err
	}

	sCtx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	srv.Shutdown(sCtx)
	if err := a.db.Disconnect(); err != nil {
		panic(err)
	}

	return nil
}

func getPort(defaultPort int) int {
	portStr, ok := os.LookupEnv("PORT")
	if !ok {
		return defaultPort
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return defaultPort
	}

	return port
}
