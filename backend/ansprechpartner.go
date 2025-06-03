package backend

import (
	"fmt"
	"sort"
	"strings"
	"viktor/db"

	"github.com/lucsky/cuid"
	"github.com/wailsapp/wails/v3/pkg/application"
)

func (a *App) CreateAnsprechpartner(Name string, Telefon, Mobil, Mail *string, lId string) bool {
	ap := db.Ansprechpartner{
		Id:            cuid.New(),
		Name:          Name,
		Telefon:       Telefon,
		Mobil:         Mobil,
		Mail:          Mail,
		LieferantenId: lId,
	}
	if !a.DB.HasKey("Ansprechpartner") {
		return a.DB.Set("Ansprechpartner", ap) == nil
	}
	aps, err := a.DB.Get("Ansprechpartner")
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[CreateAnsprechpartner] [GET(Ansprechpartner)] Fehler", err))
		dialog.Show()
		return false
	}

	for _, x := range aps.([]db.Ansprechpartner) {

		if len(*x.Mail) > 0 && len(*Mail) > 0 && x.Mail == Mail {
			dialog := application.ErrorDialog()
			dialog.SetTitle("FEHLER!")
			dialog.SetMessage(fmt.Sprintf("%s: %s", "[CreateAnsprechpartner] Ansprechpartner bereits vorhanden", err))
			dialog.Show()
			return false
		}
	}
	err = a.DB.Update("Ansprechpartner", ap)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[CreateAnsprechpartner] Fehler beim Anlegen", err))
		dialog.Show()
		return false
	}
	return true
}

func (a *App) GetAnsprechpartner(id string) *db.Ansprechpartner {
	aps := a.GetAllAnsprechpartner()

	for _, x := range aps {
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

func (a *App) GetAnsprechpartnerFromLieferant(id string) []db.Ansprechpartner {
	aps := a.GetAllAnsprechpartner()
	var res []db.Ansprechpartner
	for _, x := range aps {
		if x.LieferantenId == id {
			res = append(res, x)
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return strings.ToLower(res[i].Name) < strings.ToLower(res[j].Name)
	})

	return res
}

func (a *App) GetAllAnsprechpartner() []db.Ansprechpartner {
	aps, err := a.DB.Get("Ansprechpartner")
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[GetAllAnsprechpartner] Fehler", err))
		dialog.Show()
		return nil
	}
	return aps.([]db.Ansprechpartner)
}

func (a *App) UpdateAnsprechpartner(id, Name string, Telefon, Mobil, Mail *string) bool {
	ap := a.GetAnsprechpartner(id)
	ap.Mail = Mail
	ap.Mobil = Mobil
	ap.Telefon = Telefon
	ap.Name = Name

	err := a.DB.Update("Ansprechpartner", ap)
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
	err := a.DB.Delete("Ansprechpartner", id)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[DeleteAnsprechpartner] Fehler", err))
		dialog.Show()
		return false
	}
	return true
}
