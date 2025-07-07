package handler

import (
	"log/slog"
	"net/http"

	"github.com/computerextra/viktor/db"
	"github.com/computerextra/viktor/internal/util/flash"
)

func (h *Handler) GetPartners(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	res, err := h.db.Partner.FindMany().OrderBy(db.Partner.Name.Order(db.SortOrderAsc)).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	sendJsonData(marshalData(res, w, h.logger), w)
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
		sendQueryError(w, h.logger, err)
	}
	sendJsonData(marshalData(res, w, h.logger), w)
}

type PartnerProps struct {
	Name  string `schema:"name,required"`
	Image string `schema:"image,required"`
	Link  string `schema:"link,required"`
}

func (h *Handler) CreatePartner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	r.ParseForm()
	var partner PartnerProps
	err := decoder.Decode(&partner, r.PostForm)
	if err != nil {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		h.logger.Error("failed to parse formdata", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}
	res, err := h.db.Partner.CreateOne(
		db.Partner.Name.Set(partner.Name),
		db.Partner.Link.Set(partner.Link),
		db.Partner.Image.Set(partner.Image),
	).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	sendJsonData(marshalData(res, w, h.logger), w)
}

func (h *Handler) UpdatePartner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	r.ParseForm()
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
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		h.logger.Error("failed to parse formdata", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}
	res, err := h.db.Partner.FindUnique(db.Partner.ID.Equals(id)).Update(
		db.Partner.Name.Set(partner.Name),
		db.Partner.Link.Set(partner.Link),
		db.Partner.Image.Set(partner.Image),
	).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	sendJsonData(marshalData(res, w, h.logger), w)
}

func (h *Handler) DeletePartner(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	res, err := h.db.Partner.FindUnique(db.Partner.ID.Equals(id)).Delete().Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	sendJsonData(marshalData(res, w, h.logger), w)
}
