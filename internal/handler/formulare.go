package handler

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/computerextra/viktor/frontend"
	"github.com/computerextra/viktor/internal"
	"github.com/computerextra/viktor/internal/util/flash"

	_ "github.com/denisenkom/go-mssqldb"
)

type FormularAuswahlProps struct {
	Auswahl string `schema:"auswahl,required"`
}

type FormularKundensucheProps struct {
	Auswahl      string `schema:"auswahl,required"`
	Kundennummer string `schema:"kundennummer,required"`
}

func (h *Handler) FormularOverview(w http.ResponseWriter, r *http.Request) {
	frontend.Formulare(getPath(r.URL.Path), "", nil).Render(r.Context(), w)
}

func (h *Handler) FormularShow(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // Max Header size (e.g. 10MB)
	var props FormularAuswahlProps
	err := decoder.Decode(&props, r.PostForm)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		h.logger.Error("failed to parse formdata", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}
	frontend.Formulare(getPath(r.URL.Path), props.Auswahl, nil).Render(r.Context(), w)
}

func (h *Handler) FormularKundenSuche(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // Max Header size (e.g. 10MB)
	var props FormularKundensucheProps
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

	rows, err := conn.Query(fmt.Sprintf("SELECT Name, Vorname FROM sg_adressen WHERE KundNr LIKE '%s';", props.Kundennummer))
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		h.logger.Error("failed to query sage", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}
	defer rows.Close()

	var result internal.User

	for rows.Next() {
		var name sql.NullString
		var vorname sql.NullString
		if err := rows.Scan(&name, &vorname); err != nil {
			h.logger.Error("failed to query sage", slog.Any("error", err))
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if name.Valid {
			result.Name = name.String
		}
		if vorname.Valid {
			result.Vorname = vorname.String
		}
		if err := rows.Err(); err != nil {
			h.logger.Error("failed to query sage", slog.Any("error", err))
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	result.Kundennummer = props.Kundennummer

	frontend.Formulare(getPath(r.URL.Path), props.Auswahl, &result).Render(r.Context(), w)
}
