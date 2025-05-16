package db

// TODO: Datenbank abfragen gehen alle nicht!!!!!!!

import (
	"database/sql"
	"sort"
	"time"

	"github.com/lucsky/cuid"
)

type MitarbeiterModel struct {
	Id               string     `json:"id"`
	Name             string     `json:"Name"`
	Short            *string    `json:"Short,omitempty"`
	Gruppenwahl      *string    `json:"Gruppenwahl,omitempty"`
	InternTelefon1   *string    `json:"Intern_telefon1,omitempty"`
	InternTelefon2   *string    `json:"Intern_telefon2,omitempty"`
	FestnetzPrivat   *string    `json:"Festnetz_privat,omitempty"`
	FestnetzBusiness *string    `json:"Festnetz_busines,omitempty"`
	HomeOffice       *string    `json:"Home_office,omitempty"`
	MobilBusiness    *string    `json:"Mobil_buiness,omitempty"`
	MobilPrivat      *string    `json:"Mobil_privat,omitempty"`
	Email            *string    `json:"Email,omitempty"`
	Azubi            *bool      `json:"Azubi,omitempty"`
	Geburtstag       *time.Time `json:"Geburtstag,omitempty"`
	Paypal           *bool      `json:"Paypal,omitempty"`
	Abonniert        *bool      `json:"Abonniert,omitempty"`
	Geld             *string    `json:"Geld,omitempty"`
	Pfand            *string    `json:"Pfand,omitempty"`
	Dinge            *string    `json:"Dinge,omitempty"`
	Abgeschickt      *time.Time `json:"Abgeschickt,omitempty"`
	Bild1            *string    `json:"Bild1,omitempty"`
	Bild2            *string    `json:"Bild2,omitempty"`
	Bild3            *string    `json:"Bild3,omitempty"`
	Bild1Date        *time.Time `json:"Bild1Date,omitempty"`
	Bild2Date        *time.Time `json:"Bild2Date,omitempty"`
	Bild3Date        *time.Time `json:"Bild3Date,omitempty"`
}

type MitarbeiterParams struct {
	Name             string     `json:"Name"`
	Short            *string    `json:"Short,omitempty"`
	Gruppenwahl      *string    `json:"Gruppenwahl,omitempty"`
	InternTelefon1   *string    `json:"Intern_telefon1,omitempty"`
	InternTelefon2   *string    `json:"Intern_telefon2,omitempty"`
	FestnetzPrivat   *string    `json:"Festnetz_privat,omitempty"`
	FestnetzBusiness *string    `json:"Festnetz_busines,omitempty"`
	HomeOffice       *string    `json:"Home_office,omitempty"`
	MobilBusiness    *string    `json:"Mobil_buiness,omitempty"`
	MobilPrivat      *string    `json:"Mobil_privat,omitempty"`
	Email            *string    `json:"Email,omitempty"`
	Azubi            *bool      `json:"Azubi,omitempty"`
	Geburtstag       *time.Time `json:"Geburtstag,omitempty"`
	Paypal           *bool      `json:"Paypal,omitempty"`
	Abonniert        *bool      `json:"Abonniert,omitempty"`
	Geld             *string    `json:"Geld,omitempty"`
	Pfand            *string    `json:"Pfand,omitempty"`
	Dinge            *string    `json:"Dinge,omitempty"`
	Abgeschickt      *time.Time `json:"Abgeschickt,omitempty"`
	Bild1            *string    `json:"Bild1,omitempty"`
	Bild2            *string    `json:"Bild2,omitempty"`
	Bild3            *string    `json:"Bild3,omitempty"`
	Bild1Date        *time.Time `json:"Bild1Date,omitempty"`
	Bild2Date        *time.Time `json:"Bild2Date,omitempty"`
	Bild3Date        *time.Time `json:"Bild3Date,omitempty"`
}

// CRUD
func (d Database) createMitarbeiter(params MitarbeiterParams) error {
	db, err := sql.Open("sqlite3", d.ConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Mitarbeiter VALUES(?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	x := MitarbeiterModel{
		Id:               cuid.New(),
		Name:             params.Name,
		Short:            params.Short,
		Gruppenwahl:      params.Gruppenwahl,
		InternTelefon1:   params.InternTelefon1,
		InternTelefon2:   params.InternTelefon2,
		FestnetzPrivat:   params.FestnetzPrivat,
		FestnetzBusiness: params.FestnetzBusiness,
		HomeOffice:       params.HomeOffice,
		MobilBusiness:    params.MobilBusiness,
		MobilPrivat:      params.MobilPrivat,
		Email:            params.Email,
		Azubi:            params.Azubi,
		Geburtstag:       params.Geburtstag,
		Paypal:           params.Paypal,
		Abonniert:        params.Abonniert,
		Geld:             params.Geld,
		Pfand:            params.Pfand,
		Dinge:            params.Dinge,
		Abgeschickt:      params.Abgeschickt,
		Bild1:            params.Bild1,
		Bild2:            params.Bild2,
		Bild3:            params.Bild3,
		Bild1Date:        params.Bild1Date,
		Bild2Date:        params.Bild2Date,
		Bild3Date:        params.Bild3Date,
	}
	_, err = stmt.Exec(x)
	return err
}

func (d Database) readMitarbeiter(id *string) ([]MitarbeiterModel, error) {
	if len(*id) > 0 {
		res, err := readOneMitarbeiter(*id, d.ConnectionString)
		if err != nil {
			return nil, err
		}
		return []MitarbeiterModel{*res}, nil
	} else {
		return readAllMitarbeiter(d.ConnectionString)
	}
}

func readOneMitarbeiter(id string, connString string) (*MitarbeiterModel, error) {
	db, err := sql.Open("sqlite3", connString)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	stmt, err := db.Prepare("SELECT * FROM Mitarbeiter WHERE id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var ap MitarbeiterModel
	err = stmt.QueryRow(id).Scan(&ap)
	if err != nil {
		return nil, err
	}
	return &ap, nil
}

func readAllMitarbeiter(connString string) ([]MitarbeiterModel, error) {
	db, err := sql.Open("sqlite3", connString)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	stmt, err := db.Prepare("SELECT * FROM Mitarbeiter ORDER BY Name ASC")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var aps []MitarbeiterModel
	for rows.Next() {
		var a MitarbeiterModel
		err = rows.Scan(&a)
		if err != nil {
			return nil, err
		}
		aps = append(aps, a)
	}
	return aps, nil
}

func (d Database) updateMitarbeiter(id string, params MitarbeiterParams) error {
	db, err := sql.Open("sqlite3", d.ConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(
		`UPDATE Mitarbeiter SET 
		Name=?, 
		Short=?, 
		Gruppenwahl=?, 
		Intern_telefon1=?,
		Intern_telefon2  =?,
		Festnetz_privat=?,
		Festnetz_busines=?,
		Home_office=?,
		Mobil_buiness=?,
		Mobil_privat=?,
		Email=?,
		Azubi=?,
		Geburtstag=?,
		Paypal=?,
		Abonniert=?,
		Geld=?,
		Pfand=?,
		Dinge=?,
		Abgeschickt=?,
		Bild1=?,
		Bild2=?,
		Bild3=?,
		Bild1Date=?,
		Bild2Date=?,
		Bild3Date=?
		WHERE id=?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(params, id)
	return err
}

func (d Database) deleteMitarbeiter(id string) error {
	db, err := sql.Open("sqlite3", d.ConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM Mitarbeiter WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}

type Geburtstag struct {
	Name       string
	Geburtstag time.Time
	Diff       float64
}

type GeburtstagsListe struct {
	Vergangen []Geburtstag
	Heute     []Geburtstag
	Zukunft   []Geburtstag
}

func (d Database) GetGeburtstagsListe() (*GeburtstagsListe, error) {
	var vergangen, heute, zukunft []Geburtstag
	location, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", d.ConnectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	stmt, err := db.Prepare("SELECT * FROM Mitarbeiter WHERE Geburtstag IS NOT NULL")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var a MitarbeiterModel
		err = rows.Scan(&a)
		if err != nil {
			return nil, err
		}
		newDate := time.Date(
			time.Now().Year(),
			a.Geburtstag.Month(),
			a.Geburtstag.Day(),
			time.Now().Hour(),
			time.Now().Minute(),
			time.Now().Second(),
			time.Now().Nanosecond(),
			location,
		)
		duration := time.Since(newDate)
		days := duration.Hours() / 24

		if days < -1 {
			zukunft = append(zukunft, Geburtstag{
				Name:       a.Name,
				Geburtstag: newDate,
				Diff:       days,
			})

		} else if days == 0 {
			heute = append(heute, Geburtstag{
				Name:       a.Name,
				Geburtstag: newDate,
				Diff:       days,
			})
		} else {
			vergangen = append(vergangen, Geburtstag{
				Name:       a.Name,
				Geburtstag: newDate,
				Diff:       days,
			})
		}
	}

	sort.Slice(heute, func(i, j int) bool {
		return heute[i].Geburtstag.Before(heute[j].Geburtstag)
	})
	sort.Slice(vergangen, func(i, j int) bool {
		return vergangen[i].Geburtstag.Before(vergangen[j].Geburtstag)
	})
	sort.Slice(zukunft, func(i, j int) bool {
		return zukunft[i].Geburtstag.Before(zukunft[j].Geburtstag)
	})

	GeburtstagsListe := GeburtstagsListe{
		Vergangen: vergangen,
		Zukunft:   zukunft,
		Heute:     heute,
	}

	return &GeburtstagsListe, nil
}

func (d Database) GetEinkaufsListe() ([]MitarbeiterModel, error) {
	db, err := sql.Open("sqlite3", d.ConnectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	stmt, err := db.Prepare("SELECT * FROM Mitarbeiter WHERE Abonniert  = TRUE OR (Abgeschickt > date('now', '-1 day') AND Abgeschickt <= date('now'))")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ma []MitarbeiterModel
	for rows.Next() {
		var a MitarbeiterModel
		err = rows.Scan(&a)
		if err != nil {
			return nil, err
		}

		ma = append(ma, a)
	}

	return ma, nil
}

// TODO: Skip Einkauf
// TODO: Delete Einkauf
// TODO: Mail Senden für Abrechnung
