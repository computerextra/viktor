//go:generate bun install
//go:generate bunx @tailwindcss/cli -i ./input.css -o ./static/css/style.css --minify
//go:generate templ generate
//go:generate go run github.com/steebchen/prisma-client-go generate
package main

import (
	"context"
	"embed"
	"log/slog"
	"os"
	"os/signal"

	"github.com/computerextra/viktor/internal/app"
	"github.com/joho/godotenv"
)

//go:embed static
var files embed.FS

func main() {
	godotenv.Load()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	app, err := app.New(logger, app.Config{}, files)

	if err != nil {
		logger.Error("failed to create app", slog.Any("error", err))
	}

	if err := app.Start(ctx); err != nil {
		logger.Error("failed to start app", slog.Any("error", err))
	}
}
