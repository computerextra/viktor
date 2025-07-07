package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/computerextra/viktor/db"

	"github.com/gorilla/schema"
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

var decoder = schema.NewDecoder()

const maxBodySize = 1 << 20 // 1 MB

func sendError(w http.ResponseWriter, logger *slog.Logger, msg string, err error) {
	logger.Error(msg, slog.Any("error", err))
	w.WriteHeader(http.StatusInternalServerError)
}

func sendQueryError(w http.ResponseWriter, logger *slog.Logger, err error) {
	sendError(w, logger, "failed to query db", err)
}

func sendMarshalError(w http.ResponseWriter, logger *slog.Logger, err error) {
	sendError(w, logger, "failed to marshal data", err)
}

func marshalData(data any, w http.ResponseWriter, l *slog.Logger) []byte {
	d, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		sendMarshalError(w, l, err)
		return nil
	}
	return d
}

func sendJsonData(data []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/jason")
	w.Write(data)
}
