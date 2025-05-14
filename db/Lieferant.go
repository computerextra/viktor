package db

import (
	"database/sql"

	"github.com/lucsky/cuid"
)

type LieferantModel struct {
	Id              string  `json:"id"`
	Firma           string  `json:"Firma"`
	Kundennummer    *string `json:"Kundennummer,omitempty"`
	Webseite        *string `json:"Webseite,omitempty"`
	Ansprechpartner []AnsprechpartnerModel
}

type LieferantParams struct {
	Firma        string  `json:"Firma"`
	Kundennummer *string `json:"Kundennummer,omitempty"`
	Webseite     *string `json:"Webseite,omitempty"`
}

// CRUD
func (d Database) createLieferant(params LieferantParams) error {
	db, err := sql.Open("sqlite3", d.ConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Lieferant VALUES(?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	x := LieferantModel{
		Id:           cuid.New(),
		Firma:        params.Firma,
		Kundennummer: params.Kundennummer,
		Webseite:     params.Webseite,
	}

	_, err = stmt.Exec(x)
	return err
}

func (d Database) readLieferant(id *string) ([]LieferantModel, error) {
	if len(*id) > 0 {
		res, err := readOneLieferant(*id, d.ConnectionString)
		if err != nil {
			return nil, err
		}
		return []LieferantModel{*res}, nil
	} else {
		return readAllLieferant(d.ConnectionString)
	}
}

func readOneLieferant(id string, connString string) (*LieferantModel, error) {
	db, err := sql.Open("sqlite3", connString)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	stmt, err := db.Prepare("SELECT * FROM Lieferant WHERE id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var ap LieferantModel
	err = stmt.QueryRow(id).Scan(&ap)
	if err != nil {
		return nil, err
	}

	stmt, err = db.Prepare("SELECT * FROM Ansprechpartner WHERE lieferantId=?")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(ap.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var x AnsprechpartnerModel
		err = rows.Scan(&x)
		if err != nil {
			return nil, err
		}
		ap.Ansprechpartner = append(ap.Ansprechpartner, x)
	}

	return &ap, nil
}

func readAllLieferant(connString string) ([]LieferantModel, error) {
	db, err := sql.Open("sqlite3", connString)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	stmt, err := db.Prepare("SELECT * FROM Lieferant ORDER BY Firma ASC")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lieferanten []LieferantModel
	for rows.Next() {
		var a LieferantModel
		err = rows.Scan(&a)
		if err != nil {
			return nil, err
		}
		lieferanten = append(lieferanten, a)
	}

	stmt, err = db.Prepare("SELECT * FROM Ansprechpartner WHERE lieferantId=?")
	if err != nil {
		return nil, err
	}
	for _, a := range lieferanten {
		rows, err = stmt.Query(a.Id)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var x AnsprechpartnerModel
			err = rows.Scan(&x)
			if err != nil {
				return nil, err
			}
			a.Ansprechpartner = append(a.Ansprechpartner, x)
		}
	}

	return lieferanten, nil
}

func (d Database) updateLieferant(id string, params LieferantParams) error {
	db, err := sql.Open("sqlite3", d.ConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE Lieferant set Firma=?, Kundennummer=?, Webseite=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(params, id)
	return err
}

func (d Database) deleteLieferant(id string) error {
	db, err := sql.Open("sqlite3", d.ConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM Lieferant WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}
