package main

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
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
	a.db = db.NewDatabase(fmt.Sprintf("%s/viktor.db", a.config.Folder.Upload))
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

func (a *App) UploadImage(mitarbeiterId uint, imageNr uint) bool {
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

	m := a.db.GetMitarbeiter(mitarbeiterId)

	now := time.Now()

	if imageNr == 1 {

		m.Bild1 = &base64Encoding
		m.Bild1Date = sql.NullTime{
			Valid: true,
			Time:  now,
		}
	} else if imageNr == 2 {
		m.Bild2 = &base64Encoding
		m.Bild2Date = sql.NullTime{
			Valid: true,
			Time:  now,
		}
	} else if imageNr == 3 {
		m.Bild3 = &base64Encoding
		m.Bild3Date = sql.NullTime{
			Valid: true,
			Time:  now,
		}
	}
	a.db.UpdateMitarbeiterImages(m)
	return true
}

func (a *App) Login(mail, password string) *userdata.UserData {
	if len(mail) < 3 {
		return nil
	}
	if len(password) < 3 {
		return nil
	}
	res := a.db.CheckUser(mail, password)
	if !res {
		return nil
	}
	user := a.db.GetUserByMail(mail)

	data, err := a.userdata.Login(user.Mitarbeiter.Name, user.Mail, user.Mitarbeiter.ID)
	if err != nil {
		return nil
	}
	a.userdata = data

	return data
}

func (a *App) Logout() bool {
	a.userdata = nil
	return a.userdata.Logout() == nil
}

func (a *App) CheckSession() *userdata.UserData {
	return a.userdata
}
