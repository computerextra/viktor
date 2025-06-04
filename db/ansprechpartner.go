package db

import (
	"github.com/lucsky/cuid"
)

type Ansprechpartner struct {
	Id            string
	Name          string
	Telefon       *string
	Mobil         *string
	Mail          *string
	LieferantenId string
}

func (d *Database) CreateAnsprechpartner(
	Name string,
	Telefon,
	Mobil,
	Mail *string,
	LieferantenId string,
) (*string, error) {
	query := "INSERT INTO Ansprechpartner(Id, Name, Telefon, Mobil, Mail, LieferantenId) values(?,?,?,?,?,?);"
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
	_, err = stmt.Exec(id, Name, Telefon, Mobil, Mail, LieferantenId)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (d *Database) GetAnsprechpartner(id string) (*Ansprechpartner, error) {
	query := "SELECT Id, Name, Telefon, Mobil, Mail, LieferantenId from Ansprechpartner WHERE Id=?;"
	stmt, err := d.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var res Ansprechpartner
	err = stmt.QueryRow(id).Scan(&res.Id, &res.Name, &res.Telefon, &res.Mobil, &res.Mail, &res.LieferantenId)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (d *Database) GetAnsprechpartnerFromLieferant(LieferantenId string) ([]Ansprechpartner, error) {
	query := "SELECT Id, Name, Telefon, Mobil, Mail, LieferantenId from Ansprechpartner WHERE LieferantenId=? ORDER BY Name ASC;"
	var res []Ansprechpartner

	rows, err := d.DB.Query(query, LieferantenId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var x Ansprechpartner
		err = rows.Scan(&x.Id, &x.Name, &x.Telefon, &x.Mobil, &x.Mail, &x.LieferantenId)
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

func (d *Database) GetAllAnsprechpartner() ([]Ansprechpartner, error) {
	query := "SELECT Id, Name, Telefon, Mobil, Mail, LieferantenId from Ansprechpartner ORDER BY Name ASC;"
	var res []Ansprechpartner

	rows, err := d.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var x Ansprechpartner
		err = rows.Scan(&x.Id, &x.Name, &x.Telefon, &x.Mobil, &x.Mail, &x.LieferantenId)
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

func (d *Database) UpdateAnsprechpartner(
	Id string,
	Name string,
	Telefon,
	Mobil,
	Mail *string,
) error {
	query := "UPDATE Ansprechpartner SET Name=?, Telefon=?, Mobil=?, Mail=? WHERE Id=?;"
	tx, err := d.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(Name, Telefon, Mobil, Mail, Id)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) DeleteAnsprechpartner(Id string) error {
	query := "DELETE FROM Ansprechpartner WHERE Id=?;"
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
	return nil
}

func (d *Database) DeleteAnsprechpartnerFromLieferant(LieferantenId string) error {
	query := "DELETE FROM Ansprechpartner WHERE LieferantenId=?;"
	tx, err := d.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(LieferantenId)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
