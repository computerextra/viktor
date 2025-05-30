package db

import (
	"database/sql"
	"sort"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Mitarbeiter struct {
	gorm.Model
	Name             string
	Short            *string
	Gruppenwahl      *string
	InternTelefon1   *string
	InternTelefon2   *string
	FestnetzPrivat   *string
	FestnetzBusiness *string
	HomeOffice       *string
	MobilBusiness    *string
	MobilPrivat      *string
	Email            *string
	Azubi            bool `gorm:"default:false"`
	Geburtstag       sql.NullTime
	Paypal           bool `gorm:"default:false"`
	Abonniert        bool `gorm:"default:false"`
	Geld             *string
	Pfand            *string
	Dinge            *string
	Abgeschickt      sql.NullTime
	Bild1            *string
	Bild2            *string
	Bild3            *string
	Bild1Date        sql.NullTime
	Bild2Date        sql.NullTime
	Bild3Date        sql.NullTime
}

func (d Database) CreateMitarbeiter(
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
) error {
	var b sql.NullTime
	var day, month, year int
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		return err
	}

	if len(*Geburtstag) > 0 {
		spliites := strings.Split(*Geburtstag, ".")
		day, _ = strconv.Atoi(spliites[0])
		month, _ = strconv.Atoi(spliites[1])
		year, _ = strconv.Atoi(spliites[2])
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
		b.Valid = true
		b.Time = parsedTime
	} else {
		b.Valid = false
	}

	return d.db.Create(&Mitarbeiter{
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
		Geburtstag:       b,
		Paypal:           false,
		Abonniert:        false,
		Geld:             nil,
		Pfand:            nil,
		Dinge:            nil,
		Abgeschickt:      sql.NullTime{Valid: false},
		Bild1:            nil,
		Bild2:            nil,
		Bild3:            nil,
		Bild1Date:        sql.NullTime{Valid: false},
		Bild2Date:        sql.NullTime{Valid: false},
		Bild3Date:        sql.NullTime{Valid: false},
	}).Error
}

func (d Database) GetMitarbeiter(id uint) (*Mitarbeiter, error) {
	var m Mitarbeiter
	err := d.db.Select(
		"ID",
		"Name",
		"Short",
		"Gruppenwahl",
		"InternTelefon1",
		"InternTelefon2",
		"FestnetzPrivat",
		"FestnetzBusiness",
		"HomeOffice",
		"MobilBusiness",
		"MobilPrivat",
		"Email",
		"Azubi",
		"Geburtstag",
	).First(&m, id).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (d Database) GetAllMitarbeiter() ([]Mitarbeiter, error) {
	var m []Mitarbeiter
	err := d.db.Select(
		"ID",
		"Name",
		"Short",
		"Gruppenwahl",
		"InternTelefon1",
		"InternTelefon2",
		"FestnetzPrivat",
		"FestnetzBusiness",
		"HomeOffice",
		"MobilBusiness",
		"MobilPrivat",
		"Email",
		"Azubi",
		"Geburtstag",
	).Order("Name asc").Find(&m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (d Database) GetAllMitarbeiterEinkauf() ([]Mitarbeiter, error) {
	var m []Mitarbeiter
	err := d.db.Select(
		"ID",
		"Name",
		"Email",
		"Paypal",
		"Abonniert",
		"Geld",
		"Pfand",
		"Dinge",
	).Order("Name asc").Find(&m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (d Database) GetEinkaufsliste() ([]Mitarbeiter, error) {
	var m []Mitarbeiter
	err := d.db.Select(
		"ID",
		"Name",
		"Email",
		"Paypal",
		"Abonniert",
		"Geld",
		"Pfand",
		"Dinge",
		"Abgeschickt",
		"Bild1",
		"Bild2",
		"Bild3",
		"Bild1Date",
		"Bild2Date",
		"Bild3Date",
	).Where("DATE(Abgeschickt)=DATE('now')").
		Or("DATE(Abgeschickt)<=DATE('now') AND Abonniert=?", true).
		Order("Name asc").Find(&m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (d Database) SkipEinkauf(id uint) error {
	var ma Mitarbeiter
	err := d.db.First(&ma, id).Error
	if err != nil {
		return err
	}
	tomorrow := time.Now().AddDate(0, 0, 1)
	ma.Abgeschickt = sql.NullTime{
		Valid: true,
		Time:  tomorrow,
	}
	return d.db.Save(&ma).Error
}

func (d Database) DeleteEinkauf(id uint) error {
	var ma Mitarbeiter
	err := d.db.First(&ma, id).Error
	if err != nil {
		return err
	}
	ma.Abgeschickt = sql.NullTime{
		Valid: false,
	}
	ma.Abonniert = false
	return d.db.Save(&ma).Error
}

type Geburtstagsliste struct {
	Vergangen []Mitarbeiter
	Heute     []Mitarbeiter
	Zukunft   []Mitarbeiter
}

func (d Database) GetGeburtstagsliste() (*Geburtstagsliste, error) {
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		return nil, err
	}

	var ms []Mitarbeiter
	err = d.db.
		Select("ID", "Name", "Geburtstag").
		Not(&Mitarbeiter{Geburtstag: sql.NullTime{Valid: false}}).
		Find(&ms).
		Error
	if err != nil {
		return nil, err
	}

	var z, v, h []Mitarbeiter
	for _, m := range ms {
		var geb time.Time
		if m.Geburtstag.Valid {
			geb = m.Geburtstag.Time
		}
		newDate := time.Date(
			time.Now().Year(),
			geb.Month(),
			geb.Day(),
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
		var one, two time.Time
		if h[i].Geburtstag.Valid {
			one = h[i].Geburtstag.Time
		}
		if h[j].Geburtstag.Valid {
			two = h[j].Geburtstag.Time
		}
		return one.Before(two)
	})
	sort.Slice(v, func(i, j int) bool {
		var one, two time.Time
		if v[i].Geburtstag.Valid {
			one = v[i].Geburtstag.Time
		}
		if v[j].Geburtstag.Valid {
			two = v[j].Geburtstag.Time
		}
		return one.Before(two)
	})
	sort.Slice(z, func(i, j int) bool {
		var one, two time.Time
		if z[i].Geburtstag.Valid {
			one = z[i].Geburtstag.Time
		}
		if z[j].Geburtstag.Valid {
			two = z[j].Geburtstag.Time
		}
		return one.Before(two)
	})

	return &Geburtstagsliste{
		Vergangen: v,
		Heute:     h,
		Zukunft:   z,
	}, nil
}

func (d Database) UpdateMitarbeiterImages(m Mitarbeiter) error {
	return d.db.Save(&m).Error
}

func (d Database) UpdateMitarbeiter(
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
) error {
	var b sql.NullTime
	var day, month, year int
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		return err
	}

	if len(*Geburtstag) > 0 {
		spliites := strings.Split(*Geburtstag, ".")
		day, err = strconv.Atoi(spliites[0])
		if err != nil {
			return err
		}
		month, err = strconv.Atoi(spliites[1])
		if err != nil {
			return err
		}
		year, err = strconv.Atoi(spliites[2])
		if err != nil {
			return err
		}
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
		b.Valid = true
		b.Time = parsedTime
	}

	var m Mitarbeiter
	err = d.db.First(&m, id).Error
	if err != nil {
		return err
	}
	m.Name = Name
	if len(*Short) > 0 {
		m.Short = Short
	} else {
		m.Short = nil
	}
	if len(*Gruppenwahl) > 0 {
		m.Gruppenwahl = Gruppenwahl
	} else {
		m.Gruppenwahl = nil
	}

	if len(*InternTelefon1) > 0 {
		m.InternTelefon1 = InternTelefon1
	} else {
		m.InternTelefon1 = nil
	}

	if len(*InternTelefon2) > 0 {
		m.InternTelefon2 = InternTelefon2
	} else {
		m.InternTelefon2 = nil
	}

	if len(*FestnetzBusiness) > 0 {
		m.FestnetzBusiness = FestnetzBusiness
	} else {
		m.FestnetzBusiness = nil
	}

	if len(*FestnetzPrivat) > 0 {
		m.FestnetzPrivat = FestnetzPrivat
	} else {
		m.FestnetzPrivat = nil
	}

	if len(*HomeOffice) > 0 {
		m.HomeOffice = HomeOffice
	} else {
		m.HomeOffice = nil
	}

	if len(*MobilBusiness) > 0 {
		m.MobilBusiness = MobilBusiness
	} else {
		m.MobilBusiness = nil
	}

	if len(*MobilPrivat) > 0 {
		m.MobilPrivat = MobilPrivat
	} else {
		m.MobilPrivat = nil
	}

	if len(*Email) > 0 {
		m.Email = Email
	} else {
		m.Email = nil
	}

	m.Azubi = Azubi
	m.Geburtstag = b
	return d.db.Save(&m).Error
}

func (d Database) UpdateEinkauf(
	id uint,
	Paypal,
	Abonniert bool,
	Geld,
	Pfand,
	Dinge *string,
	bild1, bild2, bild3 bool,
) error {
	var m Mitarbeiter
	err := d.db.First(&m, id).Error
	if err != nil {
		return err
	}
	m.Paypal = Paypal
	m.Abonniert = Abonniert
	m.Geld = Geld
	m.Pfand = Pfand
	m.Dinge = Dinge
	m.Abgeschickt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	if !bild1 {
		m.Bild1 = nil
		m.Bild1Date.Valid = false
	}
	if !bild2 {
		m.Bild2 = nil
		m.Bild2Date.Valid = false
	}
	if !bild3 {
		m.Bild3 = nil
		m.Bild3Date.Valid = false
	}
	return d.db.Save(&m).Error
}

func (d Database) DeleteMitarbeiter(id uint) error {
	return d.db.Delete(&Mitarbeiter{}, id).Error
}
