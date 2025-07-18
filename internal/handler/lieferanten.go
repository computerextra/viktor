package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/computerextra/viktor/db"
	"github.com/computerextra/viktor/frontend"
	"github.com/computerextra/viktor/internal/util/flash"
)

type LieferantProps struct {
	Firma        string  `schema:"firma,required"`
	Kundennummer *string `schema:"kundennummer"`
	Webseite     *string `schema:"webseite"`
}

func (h *Handler) GetLieferanten(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	res, err := h.db.Lieferant.FindMany().With(
		db.Lieferant.Ansprechpartner.Fetch(),
	).OrderBy(db.Lieferant.Firma.Order(db.SortOrderAsc)).Exec(ctx)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}

	uri := getPath(r.URL.Path)
	frontend.LieferantenOverview(res, uri).Render(ctx, w)
}

func (h *Handler) NewLieferant(w http.ResponseWriter, r *http.Request) {
	uri := getPath(r.URL.Path)
	frontend.NeuerLieferant(uri).Render(r.Context(), w)
}

func (h *Handler) CreateLieferant(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	r.ParseMultipartForm(10 << 20)
	var props LieferantProps
	err := decoder.Decode(&props, r.PostForm)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		h.logger.Error("failed to parse formdata", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}

	_, err = h.db.Lieferant.CreateOne(
		db.Lieferant.Firma.Set(props.Firma),
		db.Lieferant.Kundennummer.SetIfPresent(props.Kundennummer),
		db.Lieferant.Webseite.SetIfPresent(props.Webseite),
	).Exec(ctx)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}

	host := r.Host
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	uri := fmt.Sprintf("%s://%s/Lieferanten", scheme, host)
	http.Redirect(w, r, uri, http.StatusFound)
}

func (h *Handler) GetLieferant(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.PathValue("id")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	res, err := h.db.Lieferant.FindUnique(db.Lieferant.ID.Equals(id)).Exec(ctx)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}
	uri := getPath(r.URL.Path)
	frontend.LieferantEdit(res, uri).Render(ctx, w)
}

func (h *Handler) UpdateLieferant(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	r.ParseMultipartForm(10 << 20)
	id := r.PathValue("id")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	var props LieferantProps
	err := decoder.Decode(&props, r.PostForm)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		h.logger.Error("failed to parse formdata", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}
	_, err = h.db.Lieferant.FindUnique(
		db.Lieferant.ID.Equals(id),
	).Update(
		db.Lieferant.Firma.Set(props.Firma),
		db.Lieferant.Kundennummer.SetIfPresent(props.Kundennummer),
		db.Lieferant.Webseite.SetIfPresent(props.Webseite),
	).Exec(ctx)

	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}
	host := r.Host
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	uri := fmt.Sprintf("%s://%s/Lieferanten", scheme, host)
	http.Redirect(w, r, uri, http.StatusFound)
}

func (h *Handler) DeleteLieferant(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.PathValue("id")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	_, err := h.db.Lieferant.FindUnique(
		db.Lieferant.ID.Equals(id),
	).Delete().Exec(ctx)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}

	host := r.Host
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	uri := fmt.Sprintf("%s://%s/Lieferanten", scheme, host)
	http.Redirect(w, r, uri, http.StatusFound)
}
