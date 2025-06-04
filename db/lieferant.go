package db

import (
	"github.com/lucsky/cuid"
)

type Lieferant struct {
	Id              string
	Firma           string
	Kundennummer    *string
	Webseite        *string
	Ansprechpartner []Ansprechpartner
}

func (d *Database) CreateLieferant(
	Firma string,
	Kundennummer *string,
	Webseite *string,
) (*string, error) {
	query := "INSERT INTO Lieferant(Id, Firma, Kundennummer, Webseite) values(?,?,?,?);"
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
	_, err = stmt.Exec(id, Firma, Kundennummer, Webseite)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (d *Database) GetLieferant(id string) (*Lieferant, error) {
	query := "SELECT Id, Firma, Kundennummer, Webseite from Lieferant WHERE Id=?;"
	stmt, err := d.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var res Lieferant
	err = stmt.QueryRow(id).Scan(&res.Id, &res.Firma, &res.Kundennummer, &res.Webseite)
	if err != nil {
		return nil, err
	}
	aps, err := d.GetAnsprechpartnerFromLieferant(id)
	if err != nil {
		return nil, err
	}
	res.Ansprechpartner = aps
	return &res, nil
}

func (d *Database) GetLieferanten() ([]Lieferant, error) {
	query := "SELECT Id, Firma, Kundennummer, Webseite from Lieferant ORDER BY Firma ASC;"
	var res []Lieferant

	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var x Lieferant
		err = rows.Scan(&x.Id, &x.Firma, &x.Kundennummer, &x.Webseite)
		if err != nil {
			return nil, err
		}
		aps, err := d.GetAnsprechpartnerFromLieferant(x.Id)
		if err != nil {
			return nil, err
		}
		x.Ansprechpartner = aps
		res = append(res, x)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (d *Database) UpdateLieferant(
	Id string,
	Firma string,
	Kundennummer *string,
	Webseite *string,
) error {
	query := "UPDATE Lieferant SET Firma=?, Kundennummer=?, Webseite=? WHERE Id=?;"
	tx, err := d.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(Firma, Kundennummer, Webseite, Id)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) DeleteLieferant(Id string) error {
	query := "DELETE FROM Lieferant WHERE Id=?;"
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

	err = d.DeleteAnsprechpartnerFromLieferant(Id)
	if err != nil {
		return err
	}

	return nil
}
