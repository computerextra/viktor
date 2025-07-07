package handler

import (
	"net/http"

	"github.com/computerextra/viktor/db"
	"github.com/computerextra/viktor/internal/util/flash"
)

func (h *Handler) GetAbteilungen(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	res, err := h.db.Abteilung.FindMany().OrderBy(
		db.Abteilung.Name.Order(db.SortOrderAsc),
	).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}

	data := marshalData(res, w, h.logger)
	sendJsonData(data, w)
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

	data := marshalData(res, w, h.logger)
	sendJsonData(data, w)
}

func (h *Handler) CreateAbteilung(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	r.ParseForm()
	name := r.FormValue("name")
	if name == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	res, err := h.db.Abteilung.CreateOne(db.Abteilung.Name.Set(name)).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}

	data := marshalData(res, w, h.logger)
	sendJsonData(data, w)
}

func (h *Handler) UpdateAbteilung(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	r.ParseForm()
	id := r.PathValue("id")
	name := r.FormValue("name")

	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if name == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	res, err := h.db.Abteilung.FindUnique(db.Abteilung.ID.Equals(id)).Update(
		db.Abteilung.Name.Set(name),
	).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}

	data := marshalData(res, w, h.logger)
	sendJsonData(data, w)
}

func (h *Handler) DeleteAbteilung(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.PathValue("id")

	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	res, err := h.db.Abteilung.FindUnique(db.Abteilung.ID.Equals(id)).Delete().Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	data := marshalData(res, w, h.logger)
	sendJsonData(data, w)
}
