package handler

import (
	"encoding/json"
	"net/http"

	"github.com/computerextra/viktor/db"
)

func (h *Handler) GetCmsCounts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	type Counts struct {
		Abteilungen int
		Angebote    int
		Jobs        int
		Mitarbeiter int
		Partner     int
	}

	resAbteilung, err := h.db.Abteilung.FindMany().Select(db.Abteilung.ID.Field()).Exec(ctx)
	if err != nil {
		sendError(w, h.logger, "failed to read Abteilungen", err)
	}
	resAngebote, err := h.db.Angebot.FindMany().Select(db.Abteilung.ID.Field()).Exec(ctx)
	if err != nil {
		sendError(w, h.logger, "failed to read Angebote", err)
	}
	resJobs, err := h.db.Jobs.FindMany().Select(db.Jobs.ID.Field()).Exec(ctx)
	if err != nil {
		sendError(w, h.logger, "failed to read Jobs", err)
	}
	resMitarbeiter, err := h.db.Mitarbeiter.FindMany().Select(db.Mitarbeiter.ID.Field()).Exec(ctx)
	if err != nil {
		sendError(w, h.logger, "failed to read Mitarbeiter", err)
	}
	resPartner, err := h.db.Partner.FindMany().Select(db.Partner.ID.Field()).Exec(ctx)
	if err != nil {
		sendError(w, h.logger, "failed to read Partner", err)
	}

	count := Counts{
		Abteilungen: len(resAbteilung),
		Angebote:    len(resAngebote),
		Jobs:        len(resJobs),
		Mitarbeiter: len(resMitarbeiter),
		Partner:     len(resPartner),
	}

	data, err := json.MarshalIndent(count, "", " ")
	if err != nil {
		sendError(w, h.logger, "failed to marshal results", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
