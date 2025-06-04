package backend

import (
	"fmt"
	"viktor/db"
	"viktor/userdata"

	"github.com/wailsapp/wails/v3/pkg/application"
)

func (a *App) CreateUser(Mail, Password string) bool {
	_, err := a.DB.CreateUser(Password, Mail)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[CreateUser] Fehler", err))
		dialog.Show()
		return false
	}

	return true
}

func (a *App) GetUser(id string) *db.User {
	user, err := a.DB.GetUser(id)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[GetUser] Fehler", err))
		dialog.Show()
		return nil
	}
	return user
}

func (a *App) ChangePassword(id, new, old string) bool {
	err := a.DB.UpdateUser(id, old, new)
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
	err := a.DB.DeleteUser(id)
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
	user, err := a.DB.GetUserByMail(mail)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[Login] Fehler", err))
		dialog.Show()
		return nil
	}
	if user.Passwort != password {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage("[Login] Falscher Benutzername oder Passwort")
		dialog.Show()
		return nil
	}

	data, err := a.userdata.Login(user.Mitarbeiter.Name, user.Mail, user.Mitarbeiter.Id, user.Id)
	if err != nil {
		return nil
	}
	a.userdata = data

	return data
}

func (a *App) Logout() bool {
	a.userdata = &userdata.UserData{}
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
