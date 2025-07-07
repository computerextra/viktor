package handler

import (
	"net/http"
	"time"

	"github.com/computerextra/viktor/db"
	"github.com/computerextra/viktor/internal/util/flash"
)

func (h *Handler) GetAngebote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	res, err := h.db.Angebot.FindMany().OrderBy(
		db.Angebot.Title.Order(db.SortOrderAsc),
	).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	sendJsonData(marshalData(res, w, h.logger), w)
}

func (h *Handler) GetAngebot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	res, err := h.db.Angebot.FindUnique(db.Angebot.ID.Equals(id)).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	sendJsonData(marshalData(res, w, h.logger), w)
}

func (h *Handler) CreateAngebot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	Title := r.FormValue("title")
	SubTitle := r.FormValue("subtitle")
	Image := r.FormValue("image")
	Link := r.FormValue("link")
	date_start_string := r.FormValue("date_start")
	date_stop_string := r.FormValue("date_stop")
	Anzeigen_string := r.FormValue("anzeigen")

	Anzeigen := false
	var sub *string

	if len(SubTitle) > 0 {
		sub = &SubTitle
	}

	if Anzeigen_string == "true" {
		Anzeigen = true
	}

	// Check requiered fields.
	if len(Title) < 1 || len(Image) < 1 || len(Link) < 1 || len(date_start_string) < 1 || len(date_stop_string) < 1 {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	Date_Start, err := time.Parse("02.01.2006", date_start_string)
	if err != nil {
		sendError(w, h.logger, "failed to parse date", err)
	}
	Date_Stop, err := time.Parse("02.01.2006", date_stop_string)
	if err != nil {
		sendError(w, h.logger, "failed to parse date", err)
	}

	res, err := h.db.Angebot.CreateOne(
		db.Angebot.Title.Set(Title),
		db.Angebot.DateStart.Set(Date_Start),
		db.Angebot.DateStop.Set(Date_Stop),
		db.Angebot.Link.Set(Link),
		db.Angebot.Image.Set(Image),
		db.Angebot.Subtitle.SetIfPresent(sub),
		db.Angebot.Anzeigen.Set(Anzeigen),
	).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	sendJsonData(marshalData(res, w, h.logger), w)
}

func (h *Handler) ToggleAngebot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")

	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	status, err := h.db.Angebot.FindUnique(db.Angebot.ID.Equals(id)).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}

	res, err := h.db.Angebot.FindUnique(db.Angebot.ID.Equals(id)).Update(
		db.Angebot.Anzeigen.Set(!status.Anzeigen),
	).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	sendJsonData(marshalData(res, w, h.logger), w)
}

func (h *Handler) UpdateAngebot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	Title := r.FormValue("title")
	SubTitle := r.FormValue("subtitle")
	Image := r.FormValue("image")
	Link := r.FormValue("link")
	date_start_string := r.FormValue("date_start")
	date_stop_string := r.FormValue("date_stop")
	Anzeigen_string := r.FormValue("anzeigen")

	Anzeigen := false
	var sub *string

	if len(SubTitle) > 0 {
		sub = &SubTitle
	}

	if Anzeigen_string == "true" {
		Anzeigen = true
	}

	// Check requiered fields.
	if len(Title) < 1 || len(Image) < 1 || len(Link) < 1 || len(date_start_string) < 1 || len(date_stop_string) < 1 {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	Date_Start, err := time.Parse("02.01.2006", date_start_string)
	if err != nil {
		sendError(w, h.logger, "failed to parse date", err)
	}
	Date_Stop, err := time.Parse("02.01.2006", date_stop_string)
	if err != nil {
		sendError(w, h.logger, "failed to parse date", err)
	}

	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	res, err := h.db.Angebot.FindUnique(db.Angebot.ID.Equals(id)).Update(
		db.Angebot.Title.Set(Title),
		db.Angebot.DateStart.Set(Date_Start),
		db.Angebot.DateStop.Set(Date_Stop),
		db.Angebot.Link.Set(Link),
		db.Angebot.Image.Set(Image),
		db.Angebot.Subtitle.SetIfPresent(sub),
		db.Angebot.Anzeigen.Set(Anzeigen),
	).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	sendJsonData(marshalData(res, w, h.logger), w)
}

func (h *Handler) DeleteAngebot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	res, err := h.db.Angebot.FindUnique(db.Angebot.ID.Equals(id)).Delete().Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	sendJsonData(marshalData(res, w, h.logger), w)
}
