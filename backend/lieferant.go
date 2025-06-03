package backend

import (
	"fmt"
	"sort"
	"strings"
	"viktor/db"

	"github.com/lucsky/cuid"
	"github.com/wailsapp/wails/v3/pkg/application"
)

func (a *App) CreateLieferant(Firma string, Kundennummer, Webseite *string) bool {
	ap := db.Lieferant{
		Id:              cuid.New(),
		Firma:           Firma,
		Kundennummer:    Kundennummer,
		Webseite:        Webseite,
		Ansprechpartner: []db.Ansprechpartner{},
	}
	if !a.DB.HasKey("Lieferant") {
		return a.DB.Set("Lieferant", ap) == nil
	}
	aps, err := a.DB.Get("Lieferant")
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[CreateLieferant] [GET(Lieferant)] Fehler", err))
		dialog.Show()
		return false
	}

	for _, x := range aps.([]db.Lieferant) {

		if x.Firma == Firma {
			dialog := application.ErrorDialog()
			dialog.SetTitle("FEHLER!")
			dialog.SetMessage(fmt.Sprintf("%s: %s", "[CreateLieferant] Lieferant bereits vorhanden", err))
			dialog.Show()
			return false
		}
	}
	err = a.DB.Update("Lieferant", ap)
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
	lieferanten := a.GetLieferanten()

	for _, x := range lieferanten {
		if x.Id == id {
			return &x
		}
	}
	dialog := application.ErrorDialog()
	dialog.SetTitle("FEHLER!")
	dialog.SetMessage("[GetAnsprechpartner] Ansprechpartner nicht gefunden")
	dialog.Show()
	return nil
}

func (a *App) GetLieferanten() []db.Lieferant {
	aps, err := a.DB.Get("Lieferant")
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[GetAllAnsprechpartner] Fehler", err))
		dialog.Show()
		return nil
	}
	var res []db.Lieferant
	for _, x := range aps.([]db.Lieferant) {
		res = append(res, db.Lieferant{
			Id:              x.Id,
			Firma:           x.Firma,
			Kundennummer:    x.Kundennummer,
			Webseite:        x.Webseite,
			Ansprechpartner: a.GetAnsprechpartnerFromLieferant(x.Id),
		})
	}

	sort.Slice(res, func(i, j int) bool {
		return strings.ToLower(res[i].Firma) < strings.ToLower(res[j].Firma)
	})
	return res
}

func (a *App) UpdateLieferant(id, Firma string, Kundennummer, Webseite *string) bool {
	lieferant := a.GetLieferant(id)

	lieferant.Firma = Firma
	lieferant.Kundennummer = Kundennummer
	lieferant.Webseite = Webseite

	err := a.DB.Update("Lieferant", lieferant)
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
	err := a.DB.Delete("Lieferant", id)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[DeleteAnsprechpartner] Fehler", err))
		dialog.Show()
		return false
	}
	return true
}
