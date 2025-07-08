package app

import (
	"fmt"
	"io/fs"
	"net/http"

	"github.com/computerextra/viktor/internal/handler"
)

func (a *App) loadRoutes() (http.Handler, error) {
	router := http.NewServeMux()

	reactApp, err := fs.Sub(a.frontend, "frontend/dist")
	if err != nil {
		return nil, fmt.Errorf("error finding dist folder: %w", err)
	}

	router.Handle("GET /", http.FileServerFS(reactApp))

	a.loadPages(router)

	return router, nil
}

func (a *App) loadPages(router *http.ServeMux) {
	h := handler.New(a.logger, a.db)

	// CMS ROUTES BEGIN

	// CMS
	router.HandleFunc("GET /api/cms", h.GetCmsCounts)

	// Abteilungen
	router.HandleFunc("GET /api/Abteilung", h.GetAbteilungen)
	router.HandleFunc("POST /api/Abteilung", h.CreateAbteilung)
	router.HandleFunc("GET /api/Abteilung/{id}", h.GetAbteilung)
	router.HandleFunc("POST /api/Abteilung/{id}", h.UpdateAbteilung)
	router.HandleFunc("DELETE /api/Abteilung/{id}", h.DeleteAbteilung)

	// Angebote
	router.HandleFunc("GET /api/Angebot", h.GetAngebote)
	router.HandleFunc("POST /api/Angebot", h.CreateAngebot)
	router.HandleFunc("GET /api/Angebot/{id}", h.GetAngebot)
	router.HandleFunc("POST /api/Angebot/{id}", h.UpdateAngebot)
	router.HandleFunc("POST /api/Angebot/{id}/toggle", h.ToggleAngebot)
	router.HandleFunc("DELETE /api/Angebot/{id}", h.DeleteAngebot)

	// Jobs
	router.HandleFunc("GET /api/Job", h.GetJobs)
	router.HandleFunc("POST /api/Job", h.CreateJob)
	router.HandleFunc("GET /api/Job/{id}", h.GetJob)
	router.HandleFunc("POST /api/Job/{id}", h.UpdateJob)
	router.HandleFunc("POST /api/Job/{id}/toggle", h.ToggleJob)
	router.HandleFunc("DELETE /api/Job/{id}", h.DeleteJob)

	// Mitarbeiter
	router.HandleFunc("GET /api/Mitarbeiter", h.GetMitarbeiters)
	router.HandleFunc("GET /api/Mitarbeiter/Abteilung", h.GetMitarbeitersWithAbteilung)
	router.HandleFunc("POST /api/Mitarbeiter", h.CreateMitarbeiter)
	router.HandleFunc("GET /api/Mitarbeiter/{id}", h.GetMitarbeiter)
	router.HandleFunc("POST /api/Mitarbeiter/{id}", h.UpdateMitarbeiter)
	router.HandleFunc("DELETE /api/Mitarbeiter/{id}", h.DeleteMitarbeiter)

	// Partner
	router.HandleFunc("GET /api/Partner", h.GetPartners)
	router.HandleFunc("POST /api/Partner", h.CreatePartner)
	router.HandleFunc("GET /api/Partner/{id}", h.GetPartner)
	router.HandleFunc("POST /api/Partner/{id}", h.UpdatePartner)
	router.HandleFunc("DELETE /api/Partner/{id}", h.DeletePartner)

	// CMS ROUTES END

	// Archive
	router.HandleFunc("GET /api/Archiv/{id}", h.GetArchive)
	router.HandleFunc("POST /api/Archiv", h.SearchArchive)

	// Einkauf
	router.HandleFunc("GET /api/Einkauf/{id}/Image", h.GetImage)       // TODO
	router.HandleFunc("POST /api/Einkauf/{id}/Image", h.UploadImage)   // TODO
	router.HandleFunc("DELETE /api/Einkauf/{id}/Image", h.DeleteImage) // TODO
	router.HandleFunc("GET /api/Einkauf/{id}", h.GetEinkauf)
	router.HandleFunc("GET /api/Einkauf", h.GetListe)
	router.HandleFunc("POST /api/Einkauf/{id}/Skip", h.SkipEinkauf)
	router.HandleFunc("DELETE /api/Einkauf/{id}", h.DeleteEinkauf)
	router.HandleFunc("POST /api/Einkauf/{id}", h.UpdateEinkauf)

	// Warenlieferung
	router.HandleFunc("POST /api/Warenlieferung/Generate", h.GenerateWarenlieferung)
	router.HandleFunc("POST /api/Warenlieferung/Send", h.SendWarenlieferung)

	router.HandleFunc("POST /api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}
