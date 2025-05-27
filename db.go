package main

import (
	"strings"
	"viktor/db"
)

// Ansprechpartner

func (a *App) CreateAnsprechpartner(Name string, Telefon, Mobil, Mail *string, LieferantenId uint) {
	a.db.CreateAnsprechpartner(Name, Telefon, Mobil, Mail, LieferantenId)
}

func (a *App) GetAnsprechpartner(id uint) *db.Ansprechpartner {
	ap, e := a.db.GetAnsprechpartner(id)
	if e != nil {
		return nil
	}
	return ap
}

func (a *App) GetAllAnsprechpartner() []db.Ansprechpartner {
	ap, err := a.db.GetAllAnsprechpartner()
	if err != nil {
		return nil
	}
	return ap
}

func (a *App) UpdateAnsprechpartner(id uint, Name string, Telefon, Mobil, Mail *string) bool {
	return a.db.UpdateAnsprechpartner(id, Name, Telefon, Mobil, Mail) != nil
}

func (a *App) DeleteAnsprechpartner(id uint) bool {
	return a.db.DeleteAnsprechpartner(id) != nil
}

// Lieferant

func (a *App) CreateLieferant(Firma string, Kundennummer, Webseite *string) bool {
	return a.db.CreateLieferant(Firma, Kundennummer, Webseite) != nil
}

func (a *App) GetLieferant(id uint) *db.Lieferant {
	l, e := a.db.GetLieferant(id)
	if e != nil {
		return nil
	}
	return l
}

func (a *App) GetLieferanten() []db.Lieferant {
	l, e := a.db.GetLieferanten()
	if e != nil {
		return nil
	}
	return l
}

func (a *App) UpdateLieferant(id uint, Firma string, Kundennummer, Webseite *string) bool {
	return a.db.UpdateLieferant(id, Firma, Kundennummer, Webseite) != nil
}

func (a *App) DeleteLieferant(id uint) bool {
	return a.db.DeleteLieferant(id) != nil
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

func (a *App) GetAllMitarbeiterEinkauf() []db.Mitarbeiter {
	m, e := a.db.GetAllMitarbeiterEinkauf()
	if e != nil {
		return nil
	}
	return m
}

func (a *App) GetMitarbeiter(id uint) *db.Mitarbeiter {
	m, e := a.db.GetMitarbeiter(id)
	if e != nil {
		return nil
	}
	return m
}

func (a *App) GetAllMitarbeiter() []db.Mitarbeiter {
	m, e := a.db.GetAllMitarbeiter()
	if e != nil {
		return nil
	}
	return m
}

func (a *App) GetEinkaufsliste() []db.Mitarbeiter {
	m, e := a.db.GetEinkaufsliste()
	if e != nil {
		return nil
	}
	return m
}

func (a *App) GetGeburtstagsliste() *db.Geburtstagsliste {
	g, e := a.db.GetGeburtstagsliste()
	if e != nil {
		return nil
	}
	return g
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
) bool {
	return a.db.UpdateMitarbeiter(
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
	) != nil
}

func (a *App) SkipEinkauf(id uint) bool {
	return a.db.SkipEinkauf(id) != nil
}

func (a *App) DeleteEinkauf(id uint) bool {
	return a.db.DeleteEinkauf(id) != nil
}

func (a *App) UpdateEinkauf(
	id uint,
	Paypal,
	Abonniert bool,
	Geld,
	Pfand,
	Dinge *string,
	bild1, bild2, bild3 bool,
) bool {
	return a.db.UpdateEinkauf(
		id,
		Paypal,
		Abonniert,
		Geld,
		Pfand,
		Dinge, bild1, bild2, bild3,
	) != nil
}

func (a *App) DeleteMitarbeiter(id uint) bool {
	return a.db.DeleteMitarbeiter(id) != nil
}

// User

func (a *App) CreateUser(Mail, Password string) string {
	// Check mail
	splittedMail := strings.Split(Mail, "@")
	if splittedMail[1] != "computer-extra.de" {
		return "Keine Firmen E-Mail Adresse angegeben"
	}
	if len(splittedMail[0]) < 3 {
		return "Firmen E-Mail Adresse darf nicht aus einem Alias bestehen"
	}
	res := a.db.CreateUser(Mail, Password)
	if res != nil {
		return res.Error()
	}
	u, err := a.db.GetUserByMail(Mail)
	if err != nil {
		return err.Error()
	}
	if len(u.Mail) < 3 {
		return "Keinen Mitarbeiter mit dieser E-Mail Adresse gefunden gefunden"
	}
	_, err = a.userdata.Login(u.Mitarbeiter.Name, u.Mail, u.Mitarbeiter.ID)
	if err != nil {
		return err.Error()
	}
	return "OK"
}

func (a *App) GetUser(id uint) *db.User {
	u, err := a.db.GetUser(id)
	if err != nil {
		return nil
	}
	return u
}

func (a *App) CheckUser(Mail, Password string) bool {
	b, e := a.db.CheckUser(Mail, Password)
	if e != nil {
		return false
	}
	return b
}

func (a *App) ChangePassword(id uint, OldPassword, NewPassword string) bool {
	return a.db.ChangePassword(id, OldPassword, NewPassword) != nil
}

func (a *App) DeleteUser(id uint) bool {
	return a.db.DeleteUser(id) != nil
}
