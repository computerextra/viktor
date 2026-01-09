package app

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"

	"github.com/computerextra/viktor/internal/handler"
	"github.com/computerextra/viktor/internal/middleware"
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
	router.HandleFunc("GET /Mitarbeiter", h.GetMitarbeitersWithAbteilung)

	// Sepa
	router.HandleFunc("GET /Sepa", h.GetMandate)
	router.HandleFunc("POST /Sepa", h.SetOfflineMandat)

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
	router.HandleFunc("GET /CMS/Abteilungen", h.GetAbteilungen)                                // Abteilungsübersicht
	router.HandleFunc("GET /CMS/Abteilungen/Neu", middleware.Auth(h.NewAbteilung))             // Neue Abteilung Formular
	router.HandleFunc("POST /CMS/Abteilungen/Neu", middleware.Auth(h.CreateAbteilung))         // Abteilung anlegen
	router.HandleFunc("GET /CMS/Abteilungen/{id}", middleware.Auth(h.GetAbteilung))            // Abteilung bearbeiten Form
	router.HandleFunc("POST /CMS/Abteilungen/{id}", middleware.Auth(h.UpdateAbteilung))        // Abteilung bearbeiten
	router.HandleFunc("POST /CMS/Abteilungen/{id}/Delete", middleware.Auth(h.DeleteAbteilung)) // Abteilung löschen
	// Angebote
	router.HandleFunc("GET /CMS/Angebote", h.GetAngebote)                                 // Angebotsübersicht
	router.HandleFunc("GET /CMS/Angebote/Neu", middleware.Auth(h.NewAngebot))             // Neue Angebot Formular
	router.HandleFunc("POST /CMS/Angebote/Neu", middleware.Auth(h.CreateAngebot))         // Angebot anlegen
	router.HandleFunc("GET /CMS/Angebote/{id}", middleware.Auth(h.GetAngebot))            // Angebot bearbeiten Form
	router.HandleFunc("POST /CMS/Angebote/{id}", middleware.Auth(h.UpdateAngebot))        // Angebot bearbeiten
	router.HandleFunc("POST /CMS/Angebote/{id}/Toggle", middleware.Auth(h.ToggleAngebot)) // Angebot bearbeiten
	router.HandleFunc("POST /CMS/Angebote/{id}/Delete", middleware.Auth(h.DeleteAngebot)) // Angebot bearbeiten
	// Jobs
	router.HandleFunc("GET /CMS/Jobs", h.GetJobs)                                 // Jobübersicht
	router.HandleFunc("GET /CMS/Jobs/Neu", middleware.Auth(h.NewJob))             // Neuer Job Formular
	router.HandleFunc("POST /CMS/Jobs/Neu", middleware.Auth(h.CreateJob))         // Job anlegen
	router.HandleFunc("GET /CMS/Jobs/{id}", middleware.Auth(h.GetJob))            // Job bearbeiten Form
	router.HandleFunc("POST /CMS/Jobs/{id}", middleware.Auth(h.UpdateJob))        // Job bearbeiten
	router.HandleFunc("POST /CMS/Jobs/{id}/Toggle", middleware.Auth(h.ToggleJob)) // Job Toggle
	router.HandleFunc("POST /CMS/Jobs/{id}/Delete", middleware.Auth(h.DeleteJob)) // Job löschen
	// Mitarbeiter
	router.HandleFunc("GET /CMS/Mitarbeiter", h.GetMitarbeitersWithAbteilung)                    // Mitarbeiterübersicht
	router.HandleFunc("GET /CMS/Mitarbeiter/Neu", middleware.Auth(h.NewMitarbeiter))             // Neuer Mitarbeiter Formular
	router.HandleFunc("POST /CMS/Mitarbeiter/Neu", middleware.Auth(h.CreateMitarbeiter))         // Mitarbeiter anlegen
	router.HandleFunc("GET /CMS/Mitarbeiter/{id}", middleware.Auth(h.GetMitarbeiter))            // Mitarbeiter bearbeiten Form
	router.HandleFunc("POST /CMS/Mitarbeiter/{id}", middleware.Auth(h.UpdateMitarbeiter))        // Mitarbeiter bearbeiten
	router.HandleFunc("POST /CMS/Mitarbeiter/{id}/Delete", middleware.Auth(h.DeleteMitarbeiter)) // Mitarbeiter löschen
	// Partner
	router.HandleFunc("GET /CMS/Partner", h.GetPartners)                                 // Partnerübersicht
	router.HandleFunc("GET /CMS/Partner/Neu", middleware.Auth(h.NewPartner))             // Neuer Partner Formular
	router.HandleFunc("POST /CMS/Partner/Neu", middleware.Auth(h.CreatePartner))         // Partner anlegen
	router.HandleFunc("GET /CMS/Partner/{id}", middleware.Auth(h.GetPartner))            // Partner bearbeiten Form
	router.HandleFunc("POST /CMS/Partner/{id}", middleware.Auth(h.UpdatePartner))        // Partner bearbeiten
	router.HandleFunc("POST /CMS/Partner/{id}/Delete", middleware.Auth(h.DeletePartner)) // Partner löschen

	// Archive
	router.HandleFunc("GET /Archiv", h.Archive)
	router.HandleFunc("POST /Archiv", h.SearchArchive)
	router.HandleFunc("GET /Archiv/{id}", h.GetArchive)

	// Warenlieferung
	router.HandleFunc("GET /Warenlieferung", middleware.Auth(h.Warenlieferung))
	router.HandleFunc("POST /Warenlieferung/Generate", middleware.Auth(h.GenerateWarenlieferung))
	router.HandleFunc("POST /Warenlieferung/Send", middleware.Auth(h.SendWarenlieferung))

	// Kunden
	router.HandleFunc("GET /Kunden", h.Kunden)
	router.HandleFunc("POST /Kunden", h.SucheKunde)

	// Lieferanten
	router.HandleFunc("GET /Lieferanten", h.GetLieferanten)
	router.HandleFunc("GET /Lieferanten/Neu", h.NewLieferant)
	router.HandleFunc("POST /Lieferanten/Neu", middleware.Auth(h.CreateLieferant))
	router.HandleFunc("GET /Lieferanten/{id}", middleware.Auth(h.GetLieferant))
	router.HandleFunc("POST /Lieferanten/{id}", middleware.Auth(h.UpdateLieferant))
	router.HandleFunc("POST /Lieferanten/{id}/Delete", middleware.Auth(h.DeleteLieferant))
	router.HandleFunc("GET /Lieferanten/{id}/Neu", middleware.Auth(h.NewAnsprechpartner))
	router.HandleFunc("POST /Lieferanten/{id}/Neu", middleware.Auth(h.CreateAnsprechpartner))
	router.HandleFunc("GET /Lieferanten/{id}/{aid}", middleware.Auth(h.GetAnsprechpartner))
	router.HandleFunc("POST /Lieferanten/{id}/{aid}", middleware.Auth(h.UpdateAnsprechpartner))
	router.HandleFunc("POST /Lieferanten/{id}/{aid}/Delete", middleware.Auth(h.DeleteAnsprechpartner))

	// Seriennummern
	router.HandleFunc("GET /Seriennummer", middleware.Auth(h.Seriennummern))
	router.HandleFunc("POST /Seriennummer", middleware.Auth(h.SearchSeriennummer))

	// Info an Kunde
	router.HandleFunc("GET /Info", middleware.Auth(h.Info))
	router.HandleFunc("POST /Info", middleware.Auth(h.SendInfo))

	// Label Sync
	router.HandleFunc("GET /Label", middleware.Auth(h.Label))
	router.HandleFunc("POST /Label", middleware.Auth(h.SyncLabel))

	// Aussteller
	router.HandleFunc("GET /Aussteller", middleware.Auth(h.Aussteller))
	router.HandleFunc("POST /Aussteller/Sync", middleware.Auth(h.SyncAussteller))
	router.HandleFunc("POST /Aussteller/Update", middleware.Auth(h.UpdateAussteller))
	router.HandleFunc("POST /Aussteller/Upload", middleware.Auth(h.UploadAussteller))

	// Formulare
	router.HandleFunc("GET /Formular", h.FormularOverview)
	router.HandleFunc("POST /Formular", h.FormularShow)
	router.HandleFunc("POST /Formular/Kunde", h.FormularKundenSuche)

	// Versand
	router.HandleFunc("GET /Versand", h.Versand)

	router.HandleFunc("POST /api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}
