package handler

import (
	"encoding/json"
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

func (h *Handler) GetSomething(w http.ResponseWriter, r *http.Request) {
	// Limit the size of the request Body
	// r.Body = http.MaxBytesReader(w, r.Body, maxBodySize)

	ctx := r.Context()
	// shortID := r.PathValue("id")

	// id, err := shortid.GetLongID(shortID)
	// if err != nil {
	// 	h.logger.Error("failed to get long id", slog.Any("error", err))
	// 	w.WriteHeader(http.StatusBadRequest)
	//	return
	// }

	// content, err := h.rdb.HGet(ctx, id.String(), "content").Result()
	// if errors.Is(err, redis.Nil) {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	component.NotFound().Render(ctx, w)
	// 	return
	// } else if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	//	return
	// }

	res, err := h.db.Mitarbeiter.FindMany().Exec(ctx)
	if err != nil {
		h.logger.Error("failed to query database", slog.Any("error", err))
	}

	data, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		h.logger.Error("failed to marshal result", slog.Any("error", err))
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
