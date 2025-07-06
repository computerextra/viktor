package handler

import (
	"log/slog"
	"net/http"

	"github.com/computerextra/viktor/db"
)

type Handler struct {
	logger *slog.Logger
	db     *db.PrismaClient
}

func New(
	logger *slog.Logger,
	db *db.PrismaClient,
) *Handler {
	return &Handler{
		logger: logger,
		db:     db,
	}
}

const maxBodySize = 1 << 20 // 1 MB

func sendError(w http.ResponseWriter, logger *slog.Logger, msg string, err error) {
	logger.Error(msg, slog.Any("error", err))
	w.WriteHeader(http.StatusInternalServerError)
}
