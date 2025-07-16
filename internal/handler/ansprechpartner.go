package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/computerextra/viktor/db"
	"github.com/computerextra/viktor/frontend"
	"github.com/computerextra/viktor/internal/util/flash"
)

type AnsprechpartnerProps struct {
	Name    string  `schema:"name,required"`
	Telefon *string `schema:"telefon"`
	Mobil   *string `schema:"mobil"`
	Mail    *string `schema:"mail"`
}

func (h *Handler) NewAnsprechpartner(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	frontend.NeuerAnsprechpartner(getPath(r.URL.Path), id).Render(r.Context(), w)
}

func (h *Handler) CreateAnsprechpartner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	r.ParseMultipartForm(10 << 20)
	var props AnsprechpartnerProps
	err := decoder.Decode(&props, r.PostForm)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		h.logger.Error("failed to parse formdata", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}

	_, err = h.db.Ansprechpartner.CreateOne(
		db.Ansprechpartner.Name.Set(props.Name),
		db.Ansprechpartner.Lieferant.Link(
			db.Lieferant.ID.Equals(id),
		),
		db.Ansprechpartner.Mail.SetIfPresent(props.Mail),
		db.Ansprechpartner.Mobil.SetIfPresent(props.Mobil),
		db.Ansprechpartner.Telefon.SetIfPresent(props.Telefon),
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

func (h *Handler) GetAnsprechpartner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	LieferantenId := r.PathValue("id")
	id := r.PathValue("aid")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	res, err := h.db.Ansprechpartner.FindUnique(db.Ansprechpartner.ID.Equals(id)).Exec(ctx)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}
	uri := getPath(r.URL.Path)
	frontend.AnsprechpartnerEdit(res, uri, LieferantenId).Render(ctx, w)
}

func (h *Handler) UpdateAnsprechpartner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	r.ParseMultipartForm(10 << 20)
	LieferantenId := r.PathValue("id")
	AnsprechpartnerId := r.PathValue("aid")
	if LieferantenId == "" || AnsprechpartnerId == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	var props AnsprechpartnerProps
	err := decoder.Decode(&props, r.PostForm)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		h.logger.Error("failed to parse formdata", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}
	_, err = h.db.Ansprechpartner.FindUnique(
		db.Ansprechpartner.ID.Equals(AnsprechpartnerId),
	).Update(
		db.Ansprechpartner.Name.Set(props.Name),
		db.Ansprechpartner.Lieferant.Link(
			db.Lieferant.ID.Equals(LieferantenId),
		),
		db.Ansprechpartner.Mail.SetIfPresent(props.Mail),
		db.Ansprechpartner.Mobil.SetIfPresent(props.Mobil),
		db.Ansprechpartner.Telefon.SetIfPresent(props.Telefon),
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

func (h *Handler) DeleteAnsprechpartner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.PathValue("aid")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	_, err := h.db.Ansprechpartner.FindUnique(
		db.Ansprechpartner.ID.Equals(id),
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
