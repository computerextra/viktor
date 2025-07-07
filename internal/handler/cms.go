package handler

import (
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
		sendQueryError(w, h.logger, err)
	}
	resAngebote, err := h.db.Angebot.FindMany().Select(db.Abteilung.ID.Field()).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	resJobs, err := h.db.Jobs.FindMany().Select(db.Jobs.ID.Field()).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	resMitarbeiter, err := h.db.Mitarbeiter.FindMany().Select(db.Mitarbeiter.ID.Field()).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	resPartner, err := h.db.Partner.FindMany().Select(db.Partner.ID.Field()).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}

	count := Counts{
		Abteilungen: len(resAbteilung),
		Angebote:    len(resAngebote),
		Jobs:        len(resJobs),
		Mitarbeiter: len(resMitarbeiter),
		Partner:     len(resPartner),
	}

	data := marshalData(count, w, h.logger)
	sendJsonData(data, w)
}
