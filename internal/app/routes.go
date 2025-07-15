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

	// Status
	router.HandleFunc("GET /Status", h.GetStatus)

	// Einkauf
	router.HandleFunc("GET /Einkauf", h.GetListe)                   // Get Einkaufsliste
	router.HandleFunc("GET /Einkauf/{id}", h.GetEinkauf)            // Get Einkauf von Ma
	router.HandleFunc("POST /Einkauf/{id}", h.UpdateEinkauf)        // Update Einkauf von Ma
	router.HandleFunc("POST /Einkauf/{id}/Skip", h.SkipEinkauf)     // Einkauf auf morgen setzen
	router.HandleFunc("POST /Einkauf/{id}/Delete", h.DeleteEinkauf) // Einkauf löschen

	// CMS
	router.HandleFunc("GET /CMS", h.GetCmsCounts) // CMS Übersicht
	// Abteilungen
	router.HandleFunc("GET /CMS/Abteilungen", h.GetAbteilungen)               // Abteilungsübersicht
	router.HandleFunc("GET /CMS/Abteilungen/Neu", h.NewAbteilung)             // Neue Abteilung Formular
	router.HandleFunc("POST /CMS/Abteilungen/Neu", h.CreateAbteilung)         // Abteilung anlegen
	router.HandleFunc("GET /CMS/Abteilungen/{id}", h.GetAbteilung)            // Abteilung bearbeiten Form
	router.HandleFunc("POST /CMS/Abteilungen/{id}", h.UpdateAbteilung)        // Abteilung bearbeiten
	router.HandleFunc("POST /CMS/Abteilungen/{id}/Delete", h.DeleteAbteilung) // Abteilung löschen
	// Angebote
	router.HandleFunc("GET /CMS/Angebote", h.GetAngebote)                // Angebotsübersicht
	router.HandleFunc("GET /CMS/Angebote/Neu", h.NewAngebot)             // Neue Angebot Formular
	router.HandleFunc("POST /CMS/Angebote/Neu", h.CreateAngebot)         // Angebot anlegen
	router.HandleFunc("GET /CMS/Angebote/{id}", h.GetAngebot)            // Angebot bearbeiten Form
	router.HandleFunc("POST /CMS/Angebote/{id}", h.UpdateAngebot)        // Angebot bearbeiten
	router.HandleFunc("POST /CMS/Angebote/{id}/Toggle", h.ToggleAngebot) // Angebot bearbeiten
	router.HandleFunc("POST /CMS/Angebote/{id}/Delete", h.DeleteAngebot) // Angebot bearbeiten
	// Jobs
	router.HandleFunc("GET /CMS/Jobs", h.GetJobs)                // Jobübersicht
	router.HandleFunc("GET /CMS/Jobs/Neu", h.NewJob)             // Neuer Job Formular
	router.HandleFunc("POST /CMS/Jobs/Neu", h.CreateJob)         // Job anlegen
	router.HandleFunc("GET /CMS/Jobs/{id}", h.GetJob)            // Job bearbeiten Form
	router.HandleFunc("POST /CMS/Jobs/{id}", h.UpdateJob)        // Job bearbeiten
	router.HandleFunc("POST /CMS/Jobs/{id}/Toggle", h.ToggleJob) // Job Toggle
	router.HandleFunc("POST /CMS/Jobs/{id}/Delete", h.DeleteJob) // Job löschen
	// Mitarbeiter
	router.HandleFunc("GET /CMS/Mitarbeiter", h.GetMitarbeitersWithAbteilung)   // Mitarbeiterübersicht
	router.HandleFunc("GET /CMS/Mitarbeiter/Neu", h.NewMitarbeiter)             // Neuer Mitarbeiter Formular
	router.HandleFunc("POST /CMS/Mitarbeiter/Neu", h.CreateMitarbeiter)         // Mitarbeiter anlegen
	router.HandleFunc("GET /CMS/Mitarbeiter/{id}", h.GetMitarbeiter)            // Mitarbeiter bearbeiten Form
	router.HandleFunc("POST /CMS/Mitarbeiter/{id}", h.UpdateMitarbeiter)        // Mitarbeiter bearbeiten
	router.HandleFunc("POST /CMS/Mitarbeiter/{id}/Delete", h.DeleteMitarbeiter) // Mitarbeiter löschen
	// Partner
	router.HandleFunc("GET /CMS/Partner", h.GetPartners)                // Partnerübersicht
	router.HandleFunc("GET /CMS/Partner/Neu", h.NewPartner)             // Neuer Partner Formular
	router.HandleFunc("POST /CMS/Partner/Neu", h.CreatePartner)         // Partner anlegen
	router.HandleFunc("GET /CMS/Partner/{id}", h.GetPartner)            // Partner bearbeiten Form
	router.HandleFunc("POST /CMS/Partner/{id}", h.UpdatePartner)        // Partner bearbeiten
	router.HandleFunc("POST /CMS/Partner/{id}/Delete", h.DeletePartner) // Partner löschen

	// Archive
	// TODO: Implement
	router.HandleFunc("GET /Archiv", h.SearchArchive)
	// TODO: Implement
	router.HandleFunc("POST /Archiv", h.SearchArchive)
	// TODO: Implement
	router.HandleFunc("GET /Archiv/{id}", h.GetArchive)

	// Warenlieferung
	// TODO: Implement
	router.HandleFunc("GET /Warenlieferung", h.GenerateWarenlieferung)
	// TODO: Implement
	router.HandleFunc("POST /Warenlieferung/Generate", h.GenerateWarenlieferung)
	// TODO: Implement
	router.HandleFunc("POST /Warenlieferung/Send", h.SendWarenlieferung)

	// Lieferanten
	// TODO: Implement

	// Kunden
	// TODO: Implement

	router.HandleFunc("POST /api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}
