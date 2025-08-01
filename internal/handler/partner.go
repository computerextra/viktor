package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/computerextra/viktor/db"
	"github.com/computerextra/viktor/frontend"
	"github.com/computerextra/viktor/internal/util/flash"
)

type PartnerProps struct {
	Name  string `schema:"name,required"`
	Image string `schema:"image,required"`
	Link  string `schema:"link,required"`
}

func (h *Handler) NewPartner(w http.ResponseWriter, r *http.Request) {
	uri := getPath(r.URL.Path)
	frontend.NeuerPartner(uri).Render(r.Context(), w)
}

func (h *Handler) GetPartners(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	res, err := h.db.Partner.FindMany().OrderBy(db.Partner.Name.Order(db.SortOrderAsc)).Exec(ctx)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}
	uri := getPath(r.URL.Path)
	frontend.ParnterOverview(res, uri).Render(ctx, w)
}

func (h *Handler) GetPartner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	res, err := h.db.Partner.FindUnique(db.Partner.ID.Equals(id)).Exec(ctx)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}
	uri := getPath(r.URL.Path)
	frontend.PartnerBearbeiten(res, uri).Render(ctx, w)
}

func (h *Handler) CreatePartner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	r.ParseMultipartForm(10 << 20) // Max Header size (e.g. 10MB)
	var partner PartnerProps
	err := decoder.Decode(&partner, r.PostForm)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		h.logger.Error("failed to parse formdata", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}
	_, err = h.db.Partner.CreateOne(
		db.Partner.Name.Set(partner.Name),
		db.Partner.Link.Set(partner.Link),
		db.Partner.Image.Set(partner.Image),
	).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	host := r.Host
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	uri := fmt.Sprintf("%s://%s/CMS/Partner", scheme, host)
	http.Redirect(w, r, uri, http.StatusFound)

}

func (h *Handler) UpdatePartner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	r.ParseMultipartForm(10 << 20) // Max Header size (e.g. 10MB)
	id := r.PathValue("id")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	r.ParseForm()
	var partner PartnerProps
	err := decoder.Decode(&partner, r.PostForm)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		h.logger.Error("failed to parse formdata", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}
	_, err = h.db.Partner.FindUnique(db.Partner.ID.Equals(id)).Update(
		db.Partner.Name.Set(partner.Name),
		db.Partner.Link.Set(partner.Link),
		db.Partner.Image.Set(partner.Image),
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
	uri := fmt.Sprintf("%s://%s/CMS/Partner", scheme, host)
	http.Redirect(w, r, uri, http.StatusFound)
}

func (h *Handler) DeletePartner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	_, err := h.db.Partner.FindUnique(db.Partner.ID.Equals(id)).Delete().Exec(ctx)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}
	host := r.Host
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	uri := fmt.Sprintf("%s://%s/CMS/Partner", scheme, host)
	http.Redirect(w, r, uri, http.StatusFound)
}
