package backend

import (
	"fmt"
	"viktor/db"

	"github.com/wailsapp/wails/v3/pkg/application"
)

func (a *App) CreateLieferant(Firma string, Kundennummer, Webseite *string) bool {
	_, err := a.DB.CreateLieferant(Firma, Kundennummer, Webseite)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[CreateLieferant] Fehler beim Anlegen", err))
		dialog.Show()
		return false
	}
	return true
}

func (a *App) GetLieferant(id string) *db.Lieferant {
	lieferant, err := a.DB.GetLieferant(id)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[GetLieferant] Fehler", err))
		dialog.Show()
		return nil
	}

	return lieferant
}

func (a *App) GetLieferanten() []db.Lieferant {
	aps, err := a.DB.GetLieferanten()
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[GetAllAnsprechpartner] Fehler", err))
		dialog.Show()
		return nil
	}

	return aps
}

func (a *App) UpdateLieferant(id, Firma string, Kundennummer, Webseite *string) bool {
	err := a.DB.UpdateLieferant(id, Firma, Kundennummer, Webseite)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[UpdateAnsprechpartner] Fehler beim Speichern", err))
		dialog.Show()
		return false
	}
	return true
}

func (a *App) DeleteLieferant(id string) bool {
	err := a.DB.DeleteLieferant(id)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[DeleteAnsprechpartner] Fehler", err))
		dialog.Show()
		return false
	}
	return true
}
