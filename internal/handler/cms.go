package handler

import (
	"net/http"

	"github.com/computerextra/viktor/db"
	"github.com/computerextra/viktor/frontend"
	"github.com/computerextra/viktor/internal"
	"github.com/computerextra/viktor/internal/util/flash"
)

func (h *Handler) GetCmsCounts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	resAbteilung, err := h.db.Abteilung.FindMany().Select(db.Abteilung.ID.Field()).Exec(ctx)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}
	resAngebote, err := h.db.Angebot.FindMany().Select(db.Abteilung.ID.Field()).Exec(ctx)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}
	resJobs, err := h.db.Jobs.FindMany().Select(db.Jobs.ID.Field()).Exec(ctx)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}
	resMitarbeiter, err := h.db.Mitarbeiter.FindMany().Select(db.Mitarbeiter.ID.Field()).Exec(ctx)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}
	resPartner, err := h.db.Partner.FindMany().Select(db.Partner.ID.Field()).Exec(ctx)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}

	count := internal.Counts{
		Abteilungen: len(resAbteilung),
		Angebote:    len(resAngebote),
		Jobs:        len(resJobs),
		Mitarbeiter: len(resMitarbeiter),
		Partner:     len(resPartner),
	}

	uri := getPath(r.URL.Path)

	frontend.CmsOverview(count, uri).Render(ctx, w)
}
