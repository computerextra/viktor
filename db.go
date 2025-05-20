package main

import (
	"viktor/db"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// Ansprechpartner

func (a *App) CreateAnsprechpartner(Name string, Telefon, Mobil, Mail *string, LieferantenId uint) {
	a.db.CreateAnsprechpartner(Name, Telefon, Mobil, Mail, LieferantenId)
}

func (a *App) GetAnsprechpartner(id uint) db.Ansprechpartner {
	return a.db.GetAnsprechpartner(id)
}

func (a *App) GetAllAnsprechpartner() []db.Ansprechpartner {
	return a.db.GetAllAnsprechpartner()
}

func (a *App) UpdateAnsprechpartner(id uint, Name string, Telefon, Mobil, Mail *string) {
	a.db.UpdateAnsprechpartner(id, Name, Telefon, Mobil, Mail)
}

func (a *App) DeleteAnsprechpartner(id uint) {
	a.db.DeleteAnsprechpartner(id)
}

// Lieferant

func (a *App) CreateLieferant(Firma string, Kundennummer, Webseite *string) {
	a.db.CreateLieferant(Firma, Kundennummer, Webseite)
}

func (a *App) GetLieferant(id uint) db.Lieferant {
	return a.db.GetLieferant(id)
}

func (a *App) GetLieferanten() []db.Lieferant {
	return a.db.GetLieferanten()
}

func (a *App) UpdateLieferant(id uint, Firma string, Kundennummer, Webseite *string) {
	a.db.UpdateLieferant(id, Firma, Kundennummer, Webseite)
}

func (a *App) DeleteLieferant(id uint) {
	a.db.DeleteLieferant(id)
}

// Mitarbeiter

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
) {
	a.db.CreateMitarbeiter(
		Name,
		Short,
		Gruppenwahl,
		InternTelefon1,
		InternTelefon2,
		FestnetzPrivat,
		FestnetzBusiness,
		HomeOffice,
		MobilBusiness,
		MobilPrivat,
		Email,
		Azubi,
		Geburtstag,
	)
}

func (a *App) GetMitarbeiter(id uint) db.Mitarbeiter {
	return a.db.GetMitarbeiter(id)
}

func (a *App) GetAllMitarbeiter() []db.Mitarbeiter {
	return a.db.GetAllMitarbeiter()
}

func (a *App) GetEinkaufsliste() []db.Mitarbeiter {
	return a.db.GetEinkaufsliste()
}

func (a *App) GetGeburtstagsliste() db.Geburtstagsliste {
	return a.db.GetGeburtstagsliste()
}

func (a *App) UpdateMitarbeiter(
	id uint,
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
) {
	runtime.LogDebug(a.ctx, *Geburtstag)
	a.db.UpdateMitarbeiter(
		id,
		Name,
		Short,
		Gruppenwahl,
		InternTelefon1,
		InternTelefon2,
		FestnetzPrivat,
		FestnetzBusiness,
		HomeOffice,
		MobilBusiness,
		MobilPrivat,
		Email,
		Azubi,
		Geburtstag,
	)
}

func (a *App) UpdateEinkauf(
	id uint,
	Paypal,
	Abonniert bool,
	Geld,
	Pfand,
	Dinge *string,
) {
	a.db.UpdateEinkauf(
		id,
		Paypal,
		Abonniert,
		Geld,
		Pfand,
		Dinge,
	)
}

func (a *App) DeleteMitarbeiter(id uint) {
	a.db.DeleteMitarbeiter(id)
}

// User

func (a *App) CreateUser(Mail, Password string) {
	a.db.CreateUser(Mail, Password)
	u := a.db.GetUserByMail(Mail)
	a.userdata.Login(u.Mitarbeiter.Name, u.Mail, u.Mitarbeiter.ID)
}

func (a *App) GetUser(id uint) db.User {
	return a.db.GetUser(id)
}

func (a *App) CheckUser(Mail, Password string) bool {
	return a.db.CheckUser(Mail, Password)
}

func (a *App) ChangePassword(id uint, OldPassword, NewPassword string) {
	a.db.ChangePassword(id, OldPassword, NewPassword)
}

func (a *App) DeleteUser(id uint) {
	a.db.DeleteUser(id)
}
