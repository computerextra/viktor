package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/computerextra/viktor/db"
	"github.com/computerextra/viktor/internal/util/flash"
)

func (h *Handler) SearchArchive(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxBodySize)

	search := r.FormValue("search")
	if search == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	ctx := r.Context()

	searchResults, err := h.db.Pdfs.FindMany(db.Pdfs.Or(
		db.Pdfs.Title.Contains(search),
		db.Pdfs.Body.Contains(search),
	),
	).Select(db.Pdfs.Title.Field(), db.Pdfs.ID.Field()).Exec(ctx)
	if err != nil {
		sendError(w, h.logger, "failed to query database", err)
	}

	data, err := json.MarshalIndent(searchResults, "", " ")
	if err != nil {
		sendError(w, h.logger, "failed to marshal data", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (h *Handler) GetArchive(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sendError(w, h.logger, "failed to convert string to int", err)
	}

	// Find File
	res, err := h.db.Pdfs.FindUnique(db.Pdfs.ID.Equals(id)).Exec(ctx)
	if err != nil {
		sendError(w, h.logger, "failed to query database", err)
	}

	// Try to read file
	archivePath, ok := os.LookupEnv("ARCHIVE_PATH")
	if !ok {
		sendError(w, h.logger, "failed to find archive path", err)
	}

	body, err := os.ReadFile(path.Join(archivePath, res.Title))
	if err != nil {
		sendError(w, h.logger, "failed read file", err)
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("Attachment; filename=%s", res.Title))
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", r.Header.Get("Conten-Length"))
	w.Write(body)
}
