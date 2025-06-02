package main

import (
	"strings"

	"viktor/db"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// Ansprechpartner

func (a *App) CreateAnsprechpartner(Name string, Telefon, Mobil, Mail *string, LieferantenId uint) bool {
	err := a.db.CreateAnsprechpartner(Name, Telefon, Mobil, Mail, LieferantenId)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return false
	}
	return true
}

func (a *App) GetAnsprechpartner(id uint) *db.Ansprechpartner {
	ap, e := a.db.GetAnsprechpartner(id)
	if e != nil {
		runtime.LogError(a.ctx, e.Error())
		return nil
	}
	return ap
}

func (a *App) GetAllAnsprechpartner() []db.Ansprechpartner {
	ap, err := a.db.GetAllAnsprechpartner()
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return nil
	}
	return ap
}

func (a *App) UpdateAnsprechpartner(id uint, Name string, Telefon, Mobil, Mail *string) bool {
	err := a.db.UpdateAnsprechpartner(id, Name, Telefon, Mobil, Mail)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return false
	}
	return true
}

func (a *App) DeleteAnsprechpartner(id uint) bool {
	err := a.db.DeleteAnsprechpartner(id)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return false
	}
	return true
}

// Lieferant

func (a *App) CreateLieferant(Firma string, Kundennummer, Webseite *string) bool {
	err := a.db.CreateLieferant(Firma, Kundennummer, Webseite)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return false
	}
	return true
}

func (a *App) GetLieferant(id uint) *db.Lieferant {
	l, e := a.db.GetLieferant(id)
	if e != nil {
		runtime.LogError(a.ctx, e.Error())
		return nil
	}
	return l
}

func (a *App) GetLieferanten() []db.Lieferant {
	l, e := a.db.GetLieferanten()
	if e != nil {
		runtime.LogError(a.ctx, e.Error())
		return nil
	}
	return l
}

func (a *App) UpdateLieferant(id uint, Firma string, Kundennummer, Webseite *string) bool {
	err := a.db.UpdateLieferant(id, Firma, Kundennummer, Webseite)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return false
	}
	return true
}

func (a *App) DeleteLieferant(id uint) bool {
	err := a.db.DeleteLieferant(id)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return false
	}
	return true
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
) bool {
	err := a.db.CreateMitarbeiter(
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
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return false
	}
	return true
}

func (a *App) GetAllMitarbeiterEinkauf() []db.Mitarbeiter {
	m, e := a.db.GetAllMitarbeiterEinkauf()
	if e != nil {
		runtime.LogError(a.ctx, e.Error())
		return nil
	}
	return m
}

func (a *App) GetMitarbeiter(id uint) *db.Mitarbeiter {
	m, e := a.db.GetMitarbeiter(id)
	if e != nil {
		runtime.LogError(a.ctx, e.Error())
		return nil
	}
	return m
}

func (a *App) GetAllMitarbeiter() []db.Mitarbeiter {
	m, e := a.db.GetAllMitarbeiter()
	if e != nil {
		runtime.LogError(a.ctx, e.Error())
		return nil
	}
	return m
}

func (a *App) GetEinkaufsliste() []db.Mitarbeiter {
	m, e := a.db.GetEinkaufsliste()
	if e != nil {
		runtime.LogError(a.ctx, e.Error())
		return nil
	}
	return m
}

func (a *App) GetGeburtstagsliste() *db.Geburtstagsliste {
	g, e := a.db.GetGeburtstagsliste()
	if e != nil {
		runtime.LogError(a.ctx, e.Error())
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
	err := a.db.UpdateMitarbeiter(
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
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return false
	}
	return true
}

func (a *App) SkipEinkauf(id uint) bool {
	err := a.db.SkipEinkauf(id)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return false
	}
	return true
}

func (a *App) DeleteEinkauf(id uint) bool {
	err := a.db.DeleteEinkauf(id)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return false
	}
	return true
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
	err := a.db.UpdateEinkauf(
		id,
		Paypal,
		Abonniert,
		Geld,
		Pfand,
		Dinge, bild1, bild2, bild3,
	)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return false
	}
	return true
}

func (a *App) DeleteMitarbeiter(id uint) bool {
	err := a.db.DeleteMitarbeiter(id)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return false
	}
	return true
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
	err := a.db.CreateUser(Mail, Password)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return err.Error()
	}
	u, err := a.db.GetUserByMail(Mail)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return err.Error()
	}
	if len(u.Mail) < 3 {
		return "Keinen Mitarbeiter mit dieser E-Mail Adresse gefunden gefunden"
	}
	_, err = a.userdata.Login(u.Mitarbeiter.Name, u.Mail, u.Mitarbeiter.ID, u.ID)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return err.Error()
	}
	return "OK"
}

func (a *App) GetUser(id uint) *db.User {
	u, err := a.db.GetUser(id)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return nil
	}
	return u
}

func (a *App) CheckUser(Mail, Password string) bool {
	b, e := a.db.CheckUser(Mail, Password)
	if e != nil {
		runtime.LogError(a.ctx, e.Error())
		return false
	}
	return b
}

func (a *App) ChangePassword(id uint, OldPassword, NewPassword string) bool {
	err := a.db.ChangePassword(id, OldPassword, NewPassword)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return false
	}
	return true
}

func (a *App) DeleteUser(id uint) bool {
	err := a.db.DeleteUser(id)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return false
	}
	return true
}

// Kanban

func (a *App) CreateBoard(userID uint, BoardName string) bool {
	err := a.db.CreateKanban(userID, BoardName)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return false
	}
	return true
}

func (a *App) CreatePost(BoardId uint, PostName string, Description *string, Importance, status string) bool {
	err := a.db.CreatePost(BoardId, PostName, Description, status, Importance)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return false
	}
	return true
}

func (a *App) GetBoardFromUser(UserId uint) []db.Kanban {
	boards, err := a.db.GetKanbanBoardsFromUser(UserId)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return nil
	}
	return boards
}

func (a *App) GetBoard(BoardId uint) *db.Kanban {
	board, err := a.db.GetKanbanBord(BoardId)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return nil
	}
	return board
}

func (a *App) UpdateBoard(BoardId uint, Name string) bool {
	if err := a.db.UpdateKanban(BoardId, Name); err != nil {
		runtime.LogError(a.ctx, err.Error())
		return false
	}
	return true
}

func (a *App) UpdatePost(PostId uint, Name string, Description *string, status, Importance string) bool {
	if err := a.db.UpdatePost(PostId, Name, Description, status, Importance); err != nil {
		runtime.LogError(a.ctx, err.Error())
		return false
	}
	return true
}

func (a *App) DeletePost(PostId uint) bool {
	if err := a.db.DeletePost(PostId); err != nil {
		runtime.LogError(a.ctx, err.Error())
		return false
	}
	return true
}

func (a *App) DeleteBoard(BoardId uint) bool {
	if err := a.db.DeleteBoard(BoardId); err != nil {
		runtime.LogError(a.ctx, err.Error())
		return false
	}
	return true
}
