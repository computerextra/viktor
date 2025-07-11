package app

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"

	"github.com/computerextra/viktor/internal/handler"
)

func (a *App) loadRoutes() (http.Handler, error) {
	static, err := a.loadStaticFiles()
	if err != nil {
		return nil, fmt.Errorf("failed to load static files: %w", err)
	}
	router := http.NewServeMux()

	router.Handle("GET /static/", http.StripPrefix("/static", static))

	a.loadPages(router)

	return router, nil
}

func (a *App) loadStaticFiles() (http.Handler, error) {
	if os.Getenv("BUILD_MODE") == "develop" {
		a.logger.Info("Running in develop Mode")
		return http.FileServer(http.Dir("./static")), nil
	}
	mode, ok := os.LookupEnv("MODE")
	if ok && mode == "dev" {
		a.logger.Info("Running in dev Mode from .env")
		return http.FileServer(http.Dir("./static")), nil
	}

	static, err := fs.Sub(a.files, "static")
	if err != nil {
		return nil, fmt.Errorf("failed to subdir static: %w", err)
	}

	return http.FileServerFS(static), nil
}

func (a *App) loadPages(router *http.ServeMux) {
	h := handler.New(a.logger, a.db)

	router.HandleFunc("GET /{$}", h.GetIndex)

	// Einkauf
	router.HandleFunc("GET /Einkauf", h.GetListe)                   // Get Einkaufsliste
	router.HandleFunc("GET /Einkauf/{id}", h.GetEinkauf)            // Get Einkauf von Ma
	router.HandleFunc("POST /Einkauf/{id}", h.UpdateEinkauf)        // Update Einkauf von Ma
	router.HandleFunc("POST /Einkauf/{id}/Skip", h.SkipEinkauf)     // Einkauf auf morgen setzen
	router.HandleFunc("POST /Einkauf/{id}/Delete", h.DeleteEinkauf) // Einkauf löschen

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

	// Warenlieferung
	router.HandleFunc("POST /api/Warenlieferung/Generate", h.GenerateWarenlieferung)
	router.HandleFunc("POST /api/Warenlieferung/Send", h.SendWarenlieferung)

	router.HandleFunc("POST /api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}
