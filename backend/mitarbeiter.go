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

	_, err = a.DB.CreateMitarbeiter(Name, Short, Gruppenwahl, InternTelefon1, InternTelefon2, FestnetzPrivat, FestnetzBusiness, HomeOffice, MobilBusiness, MobilPrivat, Email, Azubi, &geburtstag)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[CreateMitarbeiter] Fehler", err))
		dialog.Show()
		return false
	}

	return true
}

func (a *App) GetMitarbeiter(id string) *db.Mitarbeiter {
	mitarbeiter, err := a.DB.GetMitarbeiter(id)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[GetMitarbeiter] Fehler", err))
		dialog.Show()
		return nil
	}

	return mitarbeiter
}

func (a *App) GetMitarbeiterMitEinkauf(id string) *db.Mitarbeiter {
	mitarbeiter, err := a.DB.GetMitarbeiterMitEinkauf(id)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[GetMitarbeiter] Fehler", err))
		dialog.Show()
		return nil
	}

	return mitarbeiter
}

func (a *App) GetAllMitarbeiter() []db.Mitarbeiter {
	mitarbeiter, err := a.DB.GetAllMitarbeiter()
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[GetAllMitarbeiter] Fehler", err))
		dialog.Show()
		return nil
	}

	return mitarbeiter

}
func (a *App) GetAllMitarbeiterMitEinkauf() []db.Mitarbeiter {
	mitarbeiter, err := a.DB.GetAllMitarbeiterMitEinkauf()
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[GetAllMitarbeiterMitEinkauf] Fehler", err))
		dialog.Show()
		return nil
	}

	return mitarbeiter

}

func (a *App) GetEinkaufsliste() []db.Mitarbeiter {
	aps, err := a.DB.GetEinkauf()
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[GetEinkaufsliste] Fehler", err))
		dialog.Show()
		return nil
	}

	// loc, err := time.LoadLocation("Europe/Berlin")
	// if err != nil {
	// 	dialog := application.ErrorDialog()
	// 	dialog.SetTitle("FEHLER!")
	// 	dialog.SetMessage(fmt.Sprintf("%s: %s", "[GetEinkaufsliste] Fehler", err))
	// 	dialog.Show()
	// 	return nil
	// }

	// var res []db.Mitarbeiter
	// now := time.Now()
	// yesterday := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()-1, 0, 0, 0, 0, loc)
	// for _, x := range aps {
	// 	if x.Abonniert && x.Abgeschickt.Before(now) {
	// 		res = append(res, x)
	// 	}
	// 	if x.Abgeschickt.Before(now) && x.Abgeschickt.After(yesterday) {
	// 		res = append(res, x)
	// 	}
	// }

	return aps
}

func (a *App) SkipEinkauf(id string) bool {
	err := a.DB.SkipEinkauf(id)
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
	err := a.DB.DeleteEinkauf(id)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[DeleteEinkauf] Fehler", err))
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
	mitarbeiter, err := a.DB.GetAllMitarbeiter()
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[GetGeburtstagsliste] Fehler", err))
		dialog.Show()
		return nil
	}

	var z, v, h []db.Mitarbeiter
	for _, m := range mitarbeiter {
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

	err = a.DB.UpdateImage(mitarbeiterId, base64Encoding, imageNr)
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
	err := a.DB.UpdateEinkauf(id, Paypal, Abonniert, Geld, Pfand, Dinge, bild1, bild2, bild3)
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

	err = a.DB.UpdateMitarbeiter(id, Name, Short, Gruppenwahl, InternTelefon1, InternTelefon2, FestnetzPrivat, FestnetzBusiness, HomeOffice, MobilBusiness, MobilPrivat, Email, Azubi, &geburtstag)
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
	err := a.DB.DeleteMitarbeiter(id)
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
	ma, err := a.DB.GetMitarbeiter(id)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[SendPaypalMail] Fehler", err))
		dialog.Show()
		return false
	}
	var props mail.PaypalMail
	props.Benutzername = Benutzername
	props.Betrag = Betrag
	props.Mitarbeiter = *ma
	conf := a.config.Mail
	return mail.SendPaypalMail(props, conf.Server, conf.Port, conf.User, conf.Password, conf.From) == nil
}
