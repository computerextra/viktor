package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/computerextra/viktor/db"
	"github.com/computerextra/viktor/frontend"
	"github.com/computerextra/viktor/internal/util/flash"
)

type AbteilungProps struct {
	Name string `schema:"name,required"`
}

func (h *Handler) GetAbteilungen(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	res, err := h.db.Abteilung.FindMany().OrderBy(
		db.Abteilung.Name.Order(db.SortOrderAsc),
	).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}

	uri := getPath(r.URL.Path)

	frontend.AbteilungsOverview(res, uri).Render(ctx, w)
}

func (h *Handler) GetAbteilung(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.PathValue("id")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	res, err := h.db.Abteilung.FindUnique(db.Abteilung.ID.Equals(id)).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	uri := getPath(r.URL.Path)
	frontend.AbteilungBearbeiten(*res, uri).Render(ctx, w)
}

func (h *Handler) NewAbteilung(w http.ResponseWriter, r *http.Request) {
	uri := getPath(r.URL.Path)

	frontend.NeueAbteilung(uri).Render(r.Context(), w)
}

func (h *Handler) CreateAbteilung(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	r.ParseMultipartForm(10 << 20) // Max Header size (e.g. 10MB)
	var props AbteilungProps
	err := decoder.Decode(&props, r.PostForm)
	if err != nil {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		h.logger.Error("failed to parse formdata", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}

	_, err = h.db.Abteilung.CreateOne(db.Abteilung.Name.Set(props.Name)).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}

	host := r.Host
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	uri := fmt.Sprintf("%s://%s/CMS/Abteilungen", scheme, host)
	http.Redirect(w, r, uri, http.StatusFound)
}

func (h *Handler) UpdateAbteilung(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	r.ParseMultipartForm(10 << 20) // Max Header size (e.g. 10MB)
	id := r.PathValue("id")
	var props AbteilungProps
	err := decoder.Decode(&props, r.PostForm)
	if err != nil {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		h.logger.Error("failed to parse formdata", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	_, err = h.db.Abteilung.FindUnique(db.Abteilung.ID.Equals(id)).Update(
		db.Abteilung.Name.Set(props.Name),
	).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	host := r.Host
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	uri := fmt.Sprintf("%s://%s/CMS/Abteilungen", scheme, host)
	http.Redirect(w, r, uri, http.StatusFound)
}

func (h *Handler) DeleteAbteilung(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")

	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	_, err := h.db.Abteilung.FindUnique(db.Abteilung.ID.Equals(id)).Delete().Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}

	host := r.Host
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	uri := fmt.Sprintf("%s://%s/CMS/Abteilungen", scheme, host)
	http.Redirect(w, r, uri, http.StatusFound)
}
