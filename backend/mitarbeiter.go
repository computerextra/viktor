package backend

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"viktor/db"
	"viktor/mail"

	"github.com/lucsky/cuid"
	"github.com/wailsapp/wails/v3/pkg/application"
)

func (a *App) CreateMitarbeiter(
	Name string,
	Short,
	Gruppenwahl,
	InternTelefon1,
	InternTelefon2,
	FestnetzPrivat,
	FestnetzBusiness,
	HomeOffice,
	MobilBusiness,
	MobilPrivat,
	Email *string,
	Azubi bool,
	Geburtstag *string,
) bool {
	var geburtstag time.Time
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[CreateMitarbeiter]", err))
		dialog.Show()
		return false
	}
	if len(*Geburtstag) > 0 {
		spliites := strings.Split(*Geburtstag, ".")
		day, _ := strconv.Atoi(spliites[0])
		month, _ := strconv.Atoi(spliites[1])
		year, _ := strconv.Atoi(spliites[2])
		parsedTime := time.Date(
			year,
			time.Month(month),
			day,
			0,
			0,
			0,
			0,
			loc,
		)
		geburtstag = parsedTime
	}
	ma := db.Mitarbeiter{
		Id:               cuid.New(),
		Name:             Name,
		Short:            Short,
		Gruppenwahl:      Gruppenwahl,
		InternTelefon1:   InternTelefon1,
		InternTelefon2:   InternTelefon2,
		FestnetzPrivat:   FestnetzPrivat,
		FestnetzBusiness: FestnetzBusiness,
		HomeOffice:       HomeOffice,
		MobilBusiness:    MobilBusiness,
		MobilPrivat:      MobilPrivat,
		Email:            Email,
		Azubi:            Azubi,
		Geburtstag:       &geburtstag,
		Paypal:           false,
		Abonniert:        false,
		Geld:             nil,
		Pfand:            nil,
		Dinge:            nil,
		Abgeschickt:      nil,
		Bild1:            nil,
		Bild2:            nil,
		Bild3:            nil,
		Bild1Date:        nil,
		Bild2Date:        nil,
		Bild3Date:        nil,
	}
	if !a.DB.HasKey("Mitarbeiter") {
		return a.DB.Set("Mitarbeiter", ma) == nil
	}
	mitarbeiter, err := a.DB.Get("Mitarbeiter")
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[CreateMitarbeiter] [GET(Mitarbeiter)] Fehler", err))
		dialog.Show()
		return false
	}

	for _, x := range mitarbeiter.([]db.Mitarbeiter) {

		if len(*x.Email) > 0 && len(*Email) > 0 && x.Email == Email {
			dialog := application.ErrorDialog()
			dialog.SetTitle("FEHLER!")
			dialog.SetMessage(fmt.Sprintf("%s: %s", "[CreateMitarbeiter] Mitarbeiter bereits vorhanden", err))
			dialog.Show()
			return false
		}
	}
	err = a.DB.Update("Mitarbeiter", ma)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[CreateMitarbeiter] Fehler beim Anlegen", err))
		dialog.Show()
		return false
	}
	return true
}

func (a *App) GetMitarbeiter(id string) *db.Mitarbeiter {
	mitarbeiter, err := a.DB.Get("Mitarbeiter")
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[GetMitarbeiter] [GET(Mitarbeiter)] Fehler", err))
		dialog.Show()
		return nil
	}

	for _, x := range mitarbeiter.([]db.Mitarbeiter) {
		if x.Id == id {
			return &x
		}
	}
	dialog := application.ErrorDialog()
	dialog.SetTitle("FEHLER!")
	dialog.SetMessage(fmt.Sprintf("%s: %s", "[GetMitarbeiter] Mitarbeiter nicht gefunden", err))
	dialog.Show()
	return nil
}

func (a *App) GetAllMitarbeiter() []db.Mitarbeiter {
	aps, err := a.DB.Get("Mitarbeiter")
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[GetAllMitarbeiter] Fehler", err))
		dialog.Show()
		return nil
	}
	var res []db.Mitarbeiter
	for _, ma := range aps.([]db.Mitarbeiter) {
		res = append(res, db.Mitarbeiter{
			Id:               ma.Id,
			Name:             ma.Name,
			Short:            ma.Short,
			Gruppenwahl:      ma.Gruppenwahl,
			InternTelefon1:   ma.InternTelefon1,
			InternTelefon2:   ma.InternTelefon2,
			FestnetzPrivat:   ma.FestnetzPrivat,
			FestnetzBusiness: ma.FestnetzBusiness,
			HomeOffice:       ma.HomeOffice,
			MobilBusiness:    ma.MobilBusiness,
			MobilPrivat:      ma.MobilPrivat,
			Email:            ma.Email,
			Azubi:            ma.Azubi,
			Geburtstag:       ma.Geburtstag,
			Paypal:           ma.Paypal,
			Abonniert:        ma.Abonniert,
			Geld:             nil,
			Pfand:            nil,
			Dinge:            nil,
			Abgeschickt:      nil,
			Bild1:            nil,
			Bild2:            nil,
			Bild3:            nil,
			Bild1Date:        nil,
			Bild2Date:        nil,
			Bild3Date:        nil,
		})
	}

	return res
}

func (a *App) GetAllMitarbeiterEinkauf() []db.Mitarbeiter {
	aps, err := a.DB.Get("Mitarbeiter")
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[GetAllMitarbeiter] Fehler", err))
		dialog.Show()
		return nil
	}

	return aps.([]db.Mitarbeiter)
}

func (a *App) GetEinkaufsliste() []db.Mitarbeiter {
	aps := a.GetAllMitarbeiterEinkauf()

	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[GetEinkaufsliste] Fehler", err))
		dialog.Show()
		return nil
	}

	var res []db.Mitarbeiter
	now := time.Now()
	yesterday := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()-1, 0, 0, 0, 0, loc)
	for _, x := range aps {
		if x.Abonniert && x.Abgeschickt.Before(now) {
			res = append(res, x)
		}
		if x.Abgeschickt.Before(now) && x.Abgeschickt.After(yesterday) {
			res = append(res, x)
		}
	}

	return res
}

func (a *App) SkipEinkauf(id string) bool {
	loc, err := time.LoadLocation("Europe/Berlin")

	aps, err := a.DB.Get("Mitarbeiter")
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[SkipEinkauf] Fehler", err))
		dialog.Show()
		return false
	}

	var ma db.Mitarbeiter
	found := false
	tomorrow := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()-1, 0, 0, 0, 0, loc)
	for _, x := range aps.([]db.Mitarbeiter) {
		if x.Id == id {
			ma = x
			found = true
		}
	}
	if !found {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage("[SkipEinkauf] Mitarbeiter nicht gefunden")
		dialog.Show()
		return false
	}
	ma.Abgeschickt = &tomorrow
	err = a.DB.Update("Mitarbeiter", ma)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[SkipEinkauf] Fehler", err))
		dialog.Show()
		return false
	}
	return true
}

func (a *App) DeleteEinkauf(id string) bool {
	aps, err := a.DB.Get("Mitarbeiter")
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[DeleteEinkauf] Fehler", err))
		dialog.Show()
		return false
	}

	var ma db.Mitarbeiter
	found := false
	for _, x := range aps.([]db.Mitarbeiter) {
		if x.Id == id {
			ma = x
			found = true
		}
	}
	if !found {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage("[DeleteEinkauf] Mitarbeiter nicht gefunden")
		dialog.Show()
		return false
	}
	ma.Abgeschickt = nil
	ma.Pfand = nil
	ma.Dinge = nil
	ma.Geld = nil
	ma.Abonniert = false
	ma.Paypal = false
	err = a.DB.Update("Mitarbeiter", ma)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[SkipEinkauf] Fehler", err))
		dialog.Show()
		return false
	}
	return true
}

type Geburtstagsliste struct {
	Vergangen []db.Mitarbeiter
	Heute     []db.Mitarbeiter
	Zukunft   []db.Mitarbeiter
}

func (a *App) GetGeburtstagsliste() *Geburtstagsliste {
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[GetGeburtstagsliste] Fehler", err))
		dialog.Show()
		return nil
	}
	aps := a.GetAllMitarbeiter()

	var z, v, h []db.Mitarbeiter
	for _, m := range aps {
		newDate := time.Date(
			time.Now().Year(),
			m.Geburtstag.Month(),
			m.Geburtstag.Day(),
			time.Now().Hour(),
			time.Now().Minute(),
			time.Now().Second(),
			time.Now().Nanosecond(),
			loc,
		)
		dur := time.Since(newDate)
		days := dur.Hours() / 24

		if days < -1 {
			z = append(z, m)
		} else if days == 0 {
			h = append(h, m)
		} else {
			v = append(v, m)
		}

	}

	sort.Slice(h, func(i, j int) bool {
		later := *h[j].Geburtstag
		return h[i].Geburtstag.Before(later)
	})
	sort.Slice(v, func(i, j int) bool {
		later := *h[j].Geburtstag
		return h[i].Geburtstag.Before(later)
	})
	sort.Slice(z, func(i, j int) bool {
		later := *h[j].Geburtstag
		return h[i].Geburtstag.Before(later)
	})

	return &Geburtstagsliste{
		Vergangen: v,
		Heute:     h,
		Zukunft:   z,
	}
}

func (a *App) UploadImage(mitarbeiterId string, imageNr uint) bool {
	dialog := application.OpenFileDialog()
	dialog.SetTitle("Select Image")
	dialog.AddFilter("Bilder", "*.jpg;*.png;*.jpeg;*.gif")

	path, err := dialog.PromptForSingleSelection()
	if err != nil {
		errorDialog := application.ErrorDialog()
		errorDialog.SetTitle("FEHLER!")
		errorDialog.SetMessage(fmt.Sprintf("%s: %s", "[UploadImage] Fehler", err))
		errorDialog.Show()
		return false
	}

	if len(path) == 0 {
		errorDialog := application.ErrorDialog()
		errorDialog.SetTitle("FEHLER!")
		errorDialog.SetMessage(fmt.Sprintf("%s: %s", "[UploadImage] Fehler", err))
		errorDialog.Show()
		return false
	}

	data, err := os.ReadFile(path)
	if err != nil {
		errorDialog := application.ErrorDialog()
		errorDialog.SetTitle("FEHLER!")
		errorDialog.SetMessage(fmt.Sprintf("%s: %s", "[UploadImage] Fehler", err))
		errorDialog.Show()
		return false
	}

	var base64Encoding string
	mimeType := http.DetectContentType(data)
	switch mimeType {
	case "image/jpg":
		base64Encoding = "data:image/jpg;base64,"

	case "image/jpeg":
		base64Encoding = "data:image/jpeg;base64,"

	case "image/png":
		base64Encoding = "data:image/png;base64,"

	}
	base64Encoding += base64.StdEncoding.EncodeToString(data)

	ma := a.GetMitarbeiter(mitarbeiterId)
	now := time.Now()

	switch imageNr {
	case 1:
		ma.Bild1 = &base64Encoding
		ma.Bild1Date = &now
	case 2:
		ma.Bild2 = &base64Encoding
		ma.Bild2Date = &now
	case 3:
		ma.Bild3 = &base64Encoding
		ma.Bild3Date = &now
	}

	err = a.DB.Update("Mitarbeiter", ma)
	if err != nil {
		errorDialog := application.ErrorDialog()
		errorDialog.SetTitle("FEHLER!")
		errorDialog.SetMessage(fmt.Sprintf("%s: %s", "[UploadImage] Fehler", err))
		errorDialog.Show()
		return false
	}
	return true
}

func (a *App) UpdateEinkauf(id string, Paypal, Abonniert bool, Geld, Pfand, Dinge *string, bild1, bild2, bild3 bool) bool {
	now := time.Now()
	ma := a.GetMitarbeiter(id)

	ma.Paypal = Paypal
	ma.Abonniert = Abonniert
	ma.Geld = Geld
	ma.Pfand = Pfand
	ma.Dinge = Dinge

	if !bild1 {
		ma.Bild1 = nil
		ma.Bild1Date = nil
	}
	if !bild2 {
		ma.Bild2 = nil
		ma.Bild2Date = nil
	}
	if !bild3 {
		ma.Bild3 = nil
		ma.Bild3Date = nil
	}

	ma.Abgeschickt = &now

	err := a.DB.Update("Mitarbeiter", ma)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[UpdateEinkauf] Fehler", err))
		dialog.Show()
		return false
	}

	return true
}

func (a *App) UpdateMitarbeiter(
	id,
	Name string,
	Short,
	Gruppenwahl,
	InternTelefon1,
	InternTelefon2,
	FestnetzPrivat,
	FestnetzBusiness,
	HomeOffice,
	MobilBusiness,
	MobilPrivat,
	Email *string,
	Azubi bool,
	Geburtstag *string,
) bool {
	var geburtstag time.Time
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[UpdateMitarbeiter]", err))
		dialog.Show()
		return false
	}
	if len(*Geburtstag) > 0 {
		spliites := strings.Split(*Geburtstag, ".")
		day, _ := strconv.Atoi(spliites[0])
		month, _ := strconv.Atoi(spliites[1])
		year, _ := strconv.Atoi(spliites[2])
		parsedTime := time.Date(
			year,
			time.Month(month),
			day,
			0,
			0,
			0,
			0,
			loc,
		)
		geburtstag = parsedTime
	}
	ma := a.GetMitarbeiter(id)
	ma.Name = Name
	ma.Short = Short
	ma.Gruppenwahl = Gruppenwahl
	ma.InternTelefon1 = InternTelefon1
	ma.InternTelefon2 = InternTelefon2
	ma.FestnetzPrivat = FestnetzPrivat
	ma.FestnetzBusiness = FestnetzBusiness
	ma.HomeOffice = HomeOffice
	ma.MobilBusiness = MobilBusiness
	ma.MobilPrivat = MobilPrivat
	ma.Email = Email
	ma.Azubi = Azubi
	ma.Geburtstag = &geburtstag

	err = a.DB.Update("Mitarbeiter", ma)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[UpdateMitarbeiter] Fehler", err))
		dialog.Show()
		return false
	}
	return true
}

func (a *App) DeleteMitarbeiter(id string) bool {
	err := a.DB.Delete("Mitarbeiter", id)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[Mitarbeiter] Fehler", err))
		dialog.Show()
		return false
	}
	return true
}

func (a *App) SendPaypalMail(Benutzername, Betrag, id string) bool {
	ma := a.GetMitarbeiter(id)
	var props mail.PaypalMail
	props.Benutzername = Benutzername
	props.Betrag = Betrag
	props.Mitarbeiter = *ma
	conf := a.config.Mail
	return mail.SendPaypalMail(props, conf.Server, conf.Port, conf.User, conf.Password, conf.From) == nil
}
