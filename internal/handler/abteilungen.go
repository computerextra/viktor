package handler

import (
	"encoding/json"
	"net/http"

	"github.com/computerextra/viktor/db"
	"github.com/computerextra/viktor/internal/util/flash"
)

func (h *Handler) GetAbteilungen(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	res, err := h.db.Abteilung.FindMany().Exec(ctx)
	if err != nil {
		sendError(w, h.logger, "failed to query db", err)
	}

	data, err := json.MarshalIndent(res, "", " ")
	if err != nil {
		sendError(w, h.logger, "failed to marshal results", err)
	}
	w.Header().Set("Content-Type", "application/jason")
	w.Write(data)
}

func (h *Handler) GetAbteilung(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := r.FormValue("id")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	res, err := h.db.Abteilung.FindUnique(db.Abteilung.ID.Equals(id)).Exec(ctx)
	if err != nil {
		sendError(w, h.logger, "failed to query db", err)
	}

	data, err := json.MarshalIndent(res, "", " ")
	if err != nil {
		sendError(w, h.logger, "failed to marshal data", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (h *Handler) CreateAbteilung(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	name := r.FormValue("name")
	if name == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	res, err := h.db.Abteilung.CreateOne(db.Abteilung.Name.Set(name)).Exec(ctx)
	if err != nil {
		sendError(w, h.logger, "failed to create entity", err)
	}

	data, err := json.MarshalIndent(res, "", " ")
	if err != nil {
		sendError(w, h.logger, "failed to marshal data", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (h *Handler) UpdateAbteilung(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

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
		sendError(w, h.logger, "failed to update data", err)
	}

	data, err := json.MarshalIndent(res, "", " ")
	if err != nil {
		sendError(w, h.logger, "failed to marshal data", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
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
		sendError(w, h.logger, "failed to delete data", err)
	}
	data, err := json.MarshalIndent(res, "", " ")
	if err != nil {
		sendError(w, h.logger, "failed to marshal data", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
