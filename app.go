package main

import (
	"context"
	"encoding/base64"
	"net/http"
	"os"
	"time"

	"viktor/archive"
	"viktor/db"
	"viktor/sagedb"
	"viktor/userdata"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx        context.Context
	db         *db.Database
	archive    *archive.Archive
	sage       *sagedb.SageDB
	config     *Config
	userdata   *userdata.UserData
	requestCtx context.CancelFunc
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.config = NewConfig()
	a.db = db.NewDatabase(a.config.Folder.Upload)
	a.archive = archive.NewArchive(a.config.Database.Url)
	a.sage = sagedb.NewSage(a.config.Sage.Server, a.config.Sage.Database, a.config.Sage.User, a.config.Sage.Password, a.config.Sage.Port)
	a.userdata = userdata.NewUserdata()

	runtime.EventsOn(ctx, "cancelRequest", func(_ ...any) {
		if a.requestCtx != nil {
			a.requestCtx()
			a.requestCtx = nil
		}
	})
}

func (a *App) Reload() {
	runtime.WindowReload(a.ctx)
}

func (a *App) UploadImage(mitarbeiterId string, imageNr int) bool {
	file, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Bilder",
				Pattern:     "*.jpg;*.png;*.jpeg;*.gif",
			},
		},
	})
	if err != nil {
		return false
	}
	if len(file) == 0 {
		return false
	}

	data, err := os.ReadFile(file)
	if err != nil {
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

	ma, err := a.db.Read("Mitarbeiter", &mitarbeiterId, nil)
	if err != nil {
		return false
	}
	m, ok := ma.(db.MitarbeiterModel)
	if !ok {
		return false
	}

	var params db.MitarbeiterParams

	params.Name = m.Name
	params.Short = m.Short
	params.Gruppenwahl = m.Gruppenwahl
	params.InternTelefon1 = m.InternTelefon1
	params.InternTelefon2 = m.InternTelefon2
	params.FestnetzPrivat = m.FestnetzPrivat
	params.FestnetzBusiness = m.FestnetzBusiness
	params.HomeOffice = m.HomeOffice
	params.MobilBusiness = m.MobilBusiness
	params.MobilPrivat = m.MobilPrivat
	params.Email = m.Email
	params.Azubi = m.Azubi
	params.Geburtstag = m.Geburtstag
	params.Paypal = m.Paypal
	params.Abonniert = m.Abonniert
	params.Geld = m.Geld
	params.Pfand = m.Pfand
	params.Dinge = m.Dinge
	params.Abgeschickt = m.Abgeschickt
	params.Bild1 = m.Bild1
	params.Bild2 = m.Bild2
	params.Bild3 = m.Bild3
	params.Bild1Date = m.Bild1Date
	params.Bild2Date = m.Bild2Date
	params.Bild3Date = m.Bild3Date

	var now = time.Now()

	if imageNr == 1 {
		params.Bild1 = &base64Encoding
		params.Bild1Date = &now
	} else if imageNr == 2 {
		params.Bild2 = &base64Encoding
		params.Bild2Date = &now
	} else if imageNr == 3 {
		params.Bild3 = &base64Encoding
		params.Bild3Date = &now
	}

	return a.db.Update("Mitarbeiter", params, &mitarbeiterId, nil) == nil
}

func (a *App) Login(mail, password string) *userdata.UserData {
	if len(mail) < 3 {
		return nil
	}
	if len(password) < 3 {
		return nil
	}
	user, err := a.db.GetUserByMail(mail)
	if err != nil {
		return nil
	}
	if user.Password != password {
		return nil
	}
	ma, err := a.db.Read("Mitarbeiter", &user.MitarbeiterId, nil)
	if err != nil {
		return nil
	}
	m, ok := ma.(db.MitarbeiterModel)
	if !ok {
		return nil
	}
	data, err := a.userdata.Login(m.Name, mail, m.Id)
	if err != nil {
		return nil
	}
	a.userdata = data

	return a.userdata
}

func (a *App) Logout() bool {
	a.userdata = nil
	return a.userdata.Logout() == nil
}

func (a *App) CheckSession() *userdata.UserData {
	return a.userdata
}
