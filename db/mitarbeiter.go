package db

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/lucsky/cuid"
)

type Mitarbeiter struct {
	Id               string
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
	Azubi            bool
	Geburtstag       *time.Time
	Paypal           bool
	Abonniert        bool
	Geld             *string
	Pfand            *string
	Dinge            *string
	Abgeschickt      *time.Time
	Bild1            *string
	Bild2            *string
	Bild3            *string
	Bild1Date        *time.Time
	Bild2Date        *time.Time
	Bild3Date        *time.Time
}

func (d *Database) CreateMitarbeiter(
	Name string,
	Short *string,
	Gruppenwahl *string,
	InternTelefon1 *string,
	InternTelefon2 *string,
	FestnetzPrivat *string,
	FestnetzBusiness *string,
	HomeOffice *string,
	MobilBusiness *string,
	MobilPrivat *string,
	Email *string,
	Azubi bool,
	Geburtstag *time.Time,
) (*string, error) {
	query := `
	INSERT INTO Mitarbeiter
	(
	Id,
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
	Geburtstag
	) 
	values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
	tx, err := d.DB.Begin()
	if err != nil {
		return nil, err
	}
	stmt, err := tx.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	id := cuid.New()
	_, err = stmt.Exec(
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
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (d *Database) GetMitarbeiter(id string) (*Mitarbeiter, error) {
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		return nil, err
	}

	query := `
	SELECT Id, Name, Short, Gruppenwahl, InternTelefon1, InternTelefon2,
	FestnetzPrivat, FestnetzBusiness, HomeOffice, MobilBusiness,
	MobilPrivat, Email, Azubi, Geburtstag
	from Mitarbeiter WHERE Id=?;`
	stmt, err := d.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var res Mitarbeiter
	var geb string
	err = stmt.QueryRow(id).Scan(
		&res.Id,
		&res.Name,
		&res.Short,
		&res.Gruppenwahl,
		&res.InternTelefon1,
		&res.InternTelefon2,
		&res.FestnetzPrivat,
		&res.FestnetzBusiness,
		&res.HomeOffice,
		&res.MobilBusiness,
		&res.MobilPrivat,
		&res.Email,
		&res.Azubi,
		&geb,
	)

	if err != nil {
		return nil, err
	}
	var geburtstag time.Time
	if len(geb) > 1 {
		gebSplit := strings.Split(strings.Split(geb, "T")[0], "-")
		year, err := strconv.Atoi(gebSplit[0])
		if err != nil {
			return nil, err
		}
		month, err := strconv.Atoi(gebSplit[1])
		if err != nil {
			return nil, err
		}
		day, err := strconv.Atoi(gebSplit[2])
		if err != nil {
			return nil, err
		}
		geburtstag = time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc)
	}
	res.Geburtstag = &geburtstag
	return &res, nil
}

func (d *Database) GetMitarbeiterMitEinkauf(id string) (*Mitarbeiter, error) {
	query := `
	SELECT Id, Name, Email, Paypal, Abonniert, Geld, Pfand, Dinge, 
	Abgeschickt, Bild1, Bild2, Bild3, Bild1Date, Bild2Date, Bild3Date
	from Mitarbeiter WHERE Id=?;`
	stmt, err := d.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var res Mitarbeiter
	err = stmt.QueryRow(id).Scan(
		&res.Id,
		&res.Name,
		&res.Email,
		&res.Paypal,
		&res.Abonniert,
		&res.Geld,
		&res.Pfand,
		&res.Dinge,
		&res.Abgeschickt,
		&res.Bild1,
		&res.Bild2,
		&res.Bild3,
		&res.Bild1Date,
		&res.Bild2Date,
		&res.Bild3Date,
	)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (d *Database) GetEinkauf() ([]Mitarbeiter, error) {
	query := `
	SELECT Id, Name, Email, Paypal, Abonniert, Geld, Pfand, Dinge, 
	Abgeschickt, Bild1, Bild2, Bild3, Bild1Date, Bild2Date, Bild3Date
	from Mitarbeiter WHERE Abgeschickt > date('now', '-1 day') OR (Abonniert = true AND Abgeschickt < date('now')) 
	ORDER BY Abgeschickt ASC;`
	var res []Mitarbeiter

	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var x Mitarbeiter
		err = rows.Scan(
			&x.Id,
			&x.Name,
			&x.Email,
			&x.Paypal,
			&x.Abonniert,
			&x.Geld,
			&x.Pfand,
			&x.Dinge,
			&x.Abgeschickt,
			&x.Bild1,
			&x.Bild2,
			&x.Bild3,
			&x.Bild1Date,
			&x.Bild2Date,
			&x.Bild3Date,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, x)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (d *Database) GetMitarbeiterFromMail(Mail string) (*Mitarbeiter, error) {
	query := "SELECT Id, Name, Email from Mitarbeiter WHERE Email=? LIMIT 1;"
	var res Mitarbeiter

	stmt, err := d.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(Mail).Scan(
		&res.Id,
		&res.Name,
		&res.Email,
	)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (d *Database) GetAllMitarbeiter() ([]Mitarbeiter, error) {
	query := `
	SELECT Id, Name, Short, Gruppenwahl, InternTelefon1, InternTelefon2,
	FestnetzPrivat, FestnetzBusiness, HomeOffice, MobilBusiness,
	MobilPrivat, Email, Azubi, Geburtstag
	from Mitarbeiter ORDER BY Name ASC;`
	var res []Mitarbeiter

	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		return nil, err
	}

	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var x Mitarbeiter
		var geb string
		err = rows.Scan(
			&x.Id,
			&x.Name,
			&x.Short,
			&x.Gruppenwahl,
			&x.InternTelefon1,
			&x.InternTelefon2,
			&x.FestnetzPrivat,
			&x.FestnetzBusiness,
			&x.HomeOffice,
			&x.MobilBusiness,
			&x.MobilPrivat,
			&x.Email,
			&x.Azubi,
			&geb,
		)
		if err != nil {
			return nil, err
		}
		var geburtstag time.Time
		if len(geb) > 1 {
			gebSplit := strings.Split(strings.Split(geb, "T")[0], "-")
			year, err := strconv.Atoi(gebSplit[0])
			if err != nil {
				return nil, err
			}
			month, err := strconv.Atoi(gebSplit[1])
			if err != nil {
				return nil, err
			}
			day, err := strconv.Atoi(gebSplit[2])
			if err != nil {
				return nil, err
			}
			geburtstag = time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc)
		}
		x.Geburtstag = &geburtstag

		res = append(res, x)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (d *Database) GetAllMitarbeiterMitEinkauf() ([]Mitarbeiter, error) {
	query := `
	SELECT Id, Name, Email, Paypal, Abonniert, Geld, Pfand, Dinge, 
	Abgeschickt, Bild1, Bild2, Bild3, Bild1Date, Bild2Date, Bild3Date
	from Mitarbeiter ORDER BY Name ASC;`
	var res []Mitarbeiter

	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		return nil, err
	}

	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var x Mitarbeiter
		var geb string
		err = rows.Scan(
			&x.Id,
			&x.Name,
			&x.Email,
			&x.Paypal,
			&x.Abonniert,
			&x.Geld,
			&x.Pfand,
			&x.Dinge,
			&x.Abgeschickt,
			&x.Bild1,
			&x.Bild2,
			&x.Bild3,
			&x.Bild1Date,
			&x.Bild2Date,
			&x.Bild3Date,
		)
		if err != nil {
			return nil, err
		}
		var geburtstag time.Time
		if len(geb) > 1 {
			gebSplit := strings.Split(strings.Split(geb, "T")[0], "-")
			year, err := strconv.Atoi(gebSplit[0])
			if err != nil {
				return nil, err
			}
			month, err := strconv.Atoi(gebSplit[1])
			if err != nil {
				return nil, err
			}
			day, err := strconv.Atoi(gebSplit[2])
			if err != nil {
				return nil, err
			}
			geburtstag = time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc)
		}
		x.Geburtstag = &geburtstag

		res = append(res, x)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (d *Database) UpdateMitarbeiter(
	Id string,
	Name string,
	Short *string,
	Gruppenwahl *string,
	InternTelefon1 *string,
	InternTelefon2 *string,
	FestnetzPrivat *string,
	FestnetzBusiness *string,
	HomeOffice *string,
	MobilBusiness *string,
	MobilPrivat *string,
	Email *string,
	Azubi bool,
	Geburtstag *time.Time,
) error {
	query := `
	UPDATE Mitarbeiter SET 
	Name=?,
	Short=?,
	Gruppenwahl=?,
	InternTelefon1=?,
	InternTelefon2=?,
	FestnetzPrivat=?,
	FestnetzBusiness=?,
	HomeOffice=?,
	MobilBusiness=?,
	MobilPrivat=?,
	Email=?,
	Azubi=?,
	Geburtstag=?
	WHERE Id=?;`
	tx, err := d.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
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
		Id,
	)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) UpdateEinkauf(
	Id string,
	Paypal bool,
	Abonniert bool,
	Geld *string,
	Pfand *string,
	Dinge *string,
	Bild1 bool,
	Bild2 bool,
	Bild3 bool,
) error {
	query := `
	UPDATE Mitarbeiter SET 
	Paypal=?,
	Abonniert=?,
	Geld=?,
	Pfand=?,
	Dinge=?,
	Abgeschickt=date('now')
	WHERE Id=?;`
	tx, err := d.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		Paypal,
		Abonniert,
		Geld,
		Pfand,
		Dinge,
		Id,
	)
	if err != nil {
		return err
	}

	if !Bild1 {
		query = "UPDATE Mitarbeiter SET Bild1=null, Bild1Date=null WHERE Id=?;"
		stmt, err = tx.Prepare(query)
		if err != nil {
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(Id)
		if err != nil {
			return err
		}
	}
	if !Bild2 {
		query = "UPDATE Mitarbeiter SET Bild2=null, Bild2Date=null WHERE Id=?;"
		stmt, err = tx.Prepare(query)
		if err != nil {
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(Id)
		if err != nil {
			return err
		}
	}
	if !Bild3 {
		query = "UPDATE Mitarbeiter SET Bild3=null, Bild3Date=null WHERE Id=?;"
		stmt, err = tx.Prepare(query)
		if err != nil {
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(Id)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) UpdateImage(Id string, data string, imageNr uint) error {
	var query string = "UPDATE Mitarbeiter SET "
	switch imageNr {
	case 1:
		query = fmt.Sprintf("%sBild1=?, Bild1Date=date('now')", query)
	case 2:
		query = fmt.Sprintf("%sBild2=?, Bild2Date=date('now')", query)
	case 3:
		query = fmt.Sprintf("%sBild3=?, Bild3Date=date('now')", query)
	}
	query = fmt.Sprintf("%s WHERE Id=?;", query)

	tx, err := d.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		data,
		Id,
	)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) SkipEinkauf(Id string) error {
	query := "UPDATE Mitarbeiter SET Abgeschickt=date('now', '+1 day') WHERE Id=?;"
	tx, err := d.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		Id,
	)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) DeleteEinkauf(Id string) error {
	query := "UPDATE Mitarbeiter SET Abgeschickt=null, Dinge=null, Pfand=null, Geld=null, Paypal=false, Abonniert=false WHERE Id=?;"
	tx, err := d.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		Id,
	)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (d *Database) DeleteMitarbeiter(Id string) error {
	query := "DELETE FROM Mitarbeiter WHERE Id=?;"
	tx, err := d.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(Id)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}

	err = d.DeleteUser(Id)
	if err != nil {
		return err
	}
	return nil
}
