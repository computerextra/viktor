package handler

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/computerextra/viktor/db"
	"github.com/computerextra/viktor/frontend"

	"github.com/a-h/templ"
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

func Component(comp templ.Component) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		comp.Render(r.Context(), w)
	})
}

func (h *Handler) Versand(w http.ResponseWriter, r *http.Request) {
	frontend.Versand(getPath(r.URL.Path)).Render(r.Context(), w)
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

func getPath(path string) string {
	parts := strings.Split(path, "/")
	var uri string
	switch len(parts) {
	case 0:
	case 1:
		uri = ""
	default:
		uri = parts[1]
	}

	return uri
}
