package backend

import (
	"fmt"
	"viktor/db"

	"github.com/wailsapp/wails/v3/pkg/application"
)

func (a *App) CreateAnsprechpartner(Name string, Telefon, Mobil, Mail *string, lId string) bool {
	_, err := a.DB.CreateAnsprechpartner(Name, Telefon, Mobil, Mail, lId)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("[CreateAnsprechpartner] Fehler beim Anlegen: %v", err))
		dialog.Show()
		return false
	}
	return true
}

func (a *App) GetAnsprechpartner(id string) *db.Ansprechpartner {
	ap, err := a.DB.GetAnsprechpartner(id)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("[GetAnsprechpartner] Ansprechpartner nicht gefunden: %v", err))
		dialog.Show()
		return nil
	}
	return ap
}

func (a *App) GetAnsprechpartnerFromLieferant(id string) []db.Ansprechpartner {
	aps, err := a.DB.GetAnsprechpartnerFromLieferant(id)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("[GetAnsprechpartnerFromLieferant]: %v", err))
		dialog.Show()
		return nil
	}

	return aps
}

func (a *App) UpdateAnsprechpartner(id, Name string, Telefon, Mobil, Mail *string) bool {
	err := a.DB.UpdateAnsprechpartner(id, Name, Telefon, Mobil, Mail)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[UpdateAnsprechpartner] Fehler beim Speichern", err))
		dialog.Show()
		return false
	}
	return true
}

func (a *App) DeleteAnsprechpartner(id string) bool {
	err := a.DB.DeleteAnsprechpartner(id)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[DeleteAnsprechpartner] Fehler", err))
		dialog.Show()
		return false
	}
	return true
}
