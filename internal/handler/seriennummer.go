package handler

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/computerextra/viktor/frontend"
	"github.com/computerextra/viktor/internal/util/flash"

	_ "github.com/denisenkom/go-mssqldb"
)

type SNProps struct {
	Search string `schema:"search,required"`
}

func (h *Handler) Seriennummern(w http.ResponseWriter, r *http.Request) {
	frontend.Seriennummer(getPath(r.URL.Path), "", "").Render(r.Context(), w)
}

func (h *Handler) SearchSeriennummer(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxBodySize)
	r.ParseMultipartForm(10 << 20) // Max Header size (e.g. 10MB)
	var props SNProps
	err := decoder.Decode(&props, r.PostForm)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		h.logger.Error("failed to parse formdata", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}

	connString := getSageConnectionString()
	conn, err := sql.Open("sqlserver", connString)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		h.logger.Error("failed to connect to sage", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}
	defer conn.Close()

	var res string

	query := fmt.Sprintf(`
	SELECT SUCHBEGRIFF FROM sg_auf_artikel WHERE ARTNR LIKE '%s';
	`, props.Search)

	rows, err := conn.Query(query)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		h.logger.Error("failed to connect to sage", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&res,
		); err != nil {
			flash.SetFlashMessage(w, "error", err.Error())
			h.logger.Error("failed to connect to sage", slog.Any("error", err))
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	frontend.Seriennummer(getPath(r.URL.Path), res, props.Search).Render(r.Context(), w)
}
