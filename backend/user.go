package backend

import (
	"fmt"
	"viktor/db"
	"viktor/userdata"

	"github.com/lucsky/cuid"
	"github.com/wailsapp/wails/v3/pkg/application"
)

func (a *App) CreateUser(Mail, Password string) bool {
	allMa := a.GetAllMitarbeiter()
	var ma db.Mitarbeiter
	for _, x := range allMa {
		if *x.Email == Mail {
			ma = x
		}
	}
	if len(ma.Name) == 0 {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage("[CreateUser] Mitarbeiter nicht gefunden")
		dialog.Show()
		return false
	}
	user := db.User{
		Id:            cuid.New(),
		Password:      Password,
		Mail:          Mail,
		MitarbeiterId: ma.Id,
		Mitarbeiter:   ma,
		Boards:        []db.Kanban{},
	}

	if !a.DB.HasKey("User") {
		return a.DB.Set("User", user) == nil
	}
	users, err := a.DB.Get("User")
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[CreateUser] [GET(user)] Fehler", err))
		dialog.Show()
		return false
	}

	for _, x := range users.([]db.User) {

		if x.Mail == Mail {
			dialog := application.ErrorDialog()
			dialog.SetTitle("FEHLER!")
			dialog.SetMessage(fmt.Sprintf("%s: %s", "[CreateUser] Benutzer bereits vorhanden", err))
			dialog.Show()
			return false
		}
	}
	err = a.DB.Update("User", user)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[CreateUser] Fehler beim Anlegen", err))
		dialog.Show()
		return false
	}
	return true
}

func (a *App) GetUser(id string) *db.User {
	users, err := a.DB.Get("User")
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[GetUser] [GET(User)] Fehler", err))
		dialog.Show()
		return nil
	}

	for _, x := range users.([]db.User) {
		if x.Id == id {
			return &x
		}
	}
	dialog := application.ErrorDialog()
	dialog.SetTitle("FEHLER!")
	dialog.SetMessage(fmt.Sprintf("%s: %s", "[GetUser] Bentuzer nicht gefunden", err))
	dialog.Show()
	return nil
}

func (a *App) GetUserByMail(Mail string) *db.User {
	users, err := a.DB.Get("User")
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[GetUserByMail] Fehler", err))
		dialog.Show()
		return nil
	}
	for _, x := range users.([]db.User) {
		if x.Mail == Mail {
			return &x
		}
	}
	dialog := application.ErrorDialog()
	dialog.SetTitle("FEHLER!")
	dialog.SetMessage(fmt.Sprintf("%s: %s", "[GetUserByMail] Benutzer nicht gefunden", err))
	dialog.Show()
	return nil
}

func (a *App) CheckUser(Mail, Password string) bool {
	u := a.GetUserByMail(Mail)
	return u.Password == Password
}

func (a *App) ChangePassword(id, new, old string) bool {
	if new == old {
		return true
	}
	u := a.GetUser(id)
	if u.Password != old {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage("[ChangePassword] Password ist falsch")
		dialog.Show()
		return false
	}

	u.Password = new
	err := a.DB.Update("User", u)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[ChangePassword] Fehler beim Speichern", err))
		dialog.Show()
		return false
	}
	return true
}

func (a *App) DeleteUser(id string) bool {
	err := a.DB.Delete("User", id)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[DeleteUser] Fehler", err))
		dialog.Show()
		return false
	}
	return true
}

func (a *App) Login(mail, password string) *userdata.UserData {

	if len(mail) < 3 {
		return nil
	}
	if len(password) < 3 {
		return nil
	}

	res := a.CheckUser(mail, password)
	if !res {
		return nil
	}
	user := a.GetUserByMail(mail)
	data, err := a.userdata.Login(user.Mitarbeiter.Name, user.Mail, user.Mitarbeiter.Id, user.Id)
	if err != nil {

		return nil
	}
	a.userdata = data

	return data
}

func (a *App) Logout() bool {
	empty := userdata.UserData{}
	a.userdata = &empty

	return a.userdata.Logout() == nil
}

func (a *App) CheckSession() *userdata.UserData {
	data, err := userdata.ReadFile()
	if err != nil {
		return a.userdata
	} else {
		a.userdata = data
		return data
	}
}
