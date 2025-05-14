package db

import (
	"database/sql"

	"github.com/lucsky/cuid"
)

type AnsprechpartnerModel struct {
	Id            string  `json:"id"`
	Name          string  `json:"Name"`
	Telefon       *string `json:"Telefon,omitempty"`
	Mobil         *string `json:"Mobil,omitempty"`
	Mail          *string `json:"Mail,omitempty"`
	LieferantenId *string `json:"lieferantId,omitempty"`
}

type AnsprechpartnerParams struct {
	Name          string  `json:"Name"`
	Telefon       *string `json:"Telefon,omitempty"`
	Mobil         *string `json:"Mobil,omitempty"`
	Mail          *string `json:"Mail,omitempty"`
	LieferantenId *string `json:"lieferantId,omitempty"`
}

// CRUD
func (d Database) createAnsprechpartner(params AnsprechpartnerParams) error {
	db, err := sql.Open("sqlite3", d.ConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Ansprechpartner VALUES(?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	x := AnsprechpartnerModel{
		Id:            cuid.New(),
		Telefon:       params.Telefon,
		Mobil:         params.Mobil,
		Mail:          params.Mail,
		LieferantenId: params.LieferantenId,
	}
	_, err = stmt.Exec(x)
	return err
}

func (d Database) readAnsprechpartner(id *string) ([]AnsprechpartnerModel, error) {
	if len(*id) > 0 {
		res, err := readOneAnsprechpartner(*id, d.ConnectionString)
		if err != nil {
			return nil, err
		}
		return []AnsprechpartnerModel{*res}, nil
	} else {
		return readAllAnsprechpartner(d.ConnectionString)
	}
}

func readOneAnsprechpartner(id string, connString string) (*AnsprechpartnerModel, error) {
	db, err := sql.Open("sqlite3", connString)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	stmt, err := db.Prepare("SELECT * FROM Ansprechpartner WHERE id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var ap AnsprechpartnerModel
	err = stmt.QueryRow(id).Scan(&ap)
	if err != nil {
		return nil, err
	}
	return &ap, nil
}

func readAllAnsprechpartner(connString string) ([]AnsprechpartnerModel, error) {
	db, err := sql.Open("sqlite3", connString)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	stmt, err := db.Prepare("SELECT * FROM Ansprechpartner ORDER BY Name ASC")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var aps []AnsprechpartnerModel
	for rows.Next() {
		var a AnsprechpartnerModel
		err = rows.Scan(&a)
		if err != nil {
			return nil, err
		}
		aps = append(aps, a)
	}
	return aps, nil
}

func (d Database) updateAnsprechpartner(id string, params AnsprechpartnerParams) error {
	db, err := sql.Open("sqlite3", d.ConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE Ansprechpartner set Name=?, Telefon=?, Mobil=?, Mail=?, lieferantenId=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(params, id)
	return err
}

func (d Database) deleteAnsprechpartner(id string) error {
	db, err := sql.Open("sqlite3", d.ConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM Ansprechpartner WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}
