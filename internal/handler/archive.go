package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/computerextra/viktor/db"
	"github.com/computerextra/viktor/frontend"
	"github.com/computerextra/viktor/internal/util/flash"
)

type ArchiveProps struct {
	Search string `schema:"search,required"`
}

func (h *Handler) Archive(w http.ResponseWriter, r *http.Request) {
	uri := getPath(r.URL.Path)
	frontend.Archiv(uri, nil, "").Render(r.Context(), w)
}

func (h *Handler) SearchArchive(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxBodySize)
	r.ParseMultipartForm(10 << 20) // Max Header size (e.g. 10MB)
	var props ArchiveProps
	err := decoder.Decode(&props, r.PostForm)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		h.logger.Error("failed to parse formdata", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}

	ctx := r.Context()

	searchResults, err := h.db.Pdfs.FindMany(db.Pdfs.Or(
		db.Pdfs.Title.Contains(props.Search),
		db.Pdfs.Body.Contains(props.Search),
	),
	).Select(db.Pdfs.Title.Field(), db.Pdfs.ID.Field()).Exec(ctx)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}

	uri := getPath(r.URL.Path)
	frontend.Archiv(uri, searchResults, props.Search).Render(r.Context(), w)
}

func (h *Handler) GetArchive(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendError(w, h.logger, "failed to convert string to int", err)
	}

	// Find File
	res, err := h.db.Pdfs.FindUnique(db.Pdfs.ID.Equals(id)).Exec(ctx)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}

	// Try to read file
	archivePath, ok := os.LookupEnv("ARCHIVE_PATH")
	if !ok {
		flash.SetFlashMessage(w, "error", "failed to find archive path")
		sendError(w, h.logger, "failed to find archive path", err)
	}

	body, err := os.ReadFile(path.Join(archivePath, res.Title))
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendError(w, h.logger, "failed read file", err)
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("Attachment; filename=%s", res.Title))
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", r.Header.Get("Conten-Length"))
	w.Write(body)
}
