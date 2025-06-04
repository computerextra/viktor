package db

import (
	"errors"

	"github.com/lucsky/cuid"
)

type User struct {
	Id            string
	Passwort      string
	Mail          string
	MitarbeiterId string
	Mitarbeiter   Mitarbeiter
	Boards        []Kanban
}

func (d *Database) CreateUser(
	Passwort string,
	Mail string,
) (*string, error) {
	query := "INSERT INTO User(Id, Passwort, Mail, MitarbeiterId) values(?,?,?,?);"
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
	ma, err := d.GetMitarbeiterFromMail(Mail)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(id, Passwort, Mail, ma.Id)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (d *Database) GetUser(id string) (*User, error) {
	query := "SELECT Id, Passwort, Mail, MitarbeiterId from User WHERE Id=?;"
	stmt, err := d.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var res User
	err = stmt.QueryRow(id).Scan(&res.Id, &res.Passwort, &res.Mail, &res.MitarbeiterId)
	if err != nil {
		return nil, err
	}
	ma, err := d.GetMitarbeiter(res.MitarbeiterId)
	if err != nil {
		return nil, err
	}
	res.Mitarbeiter = *ma
	return &res, nil
}

func (d *Database) GetUserByMail(Mail string) (*User, error) {
	query := "SELECT Id, Passwort, Mail, MitarbeiterId from User WHERE Mail=?;"
	stmt, err := d.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var res User
	err = stmt.QueryRow(Mail).Scan(&res.Id, &res.Passwort, &res.Mail, &res.MitarbeiterId)
	if err != nil {
		return nil, err
	}
	ma, err := d.GetMitarbeiter(res.MitarbeiterId)
	if err != nil {
		return nil, err
	}
	res.Mitarbeiter = *ma
	return &res, nil
}

func (d *Database) UpdateUser(
	Id string,
	AltesPasswort string,
	NeuesPasswort string,
) error {
	user, err := d.GetUser(Id)
	if err != nil {
		return err
	}
	if user.Passwort != AltesPasswort {
		return errors.New("Falsches Passwort angegeben")
	}

	query := "UPDATE User SET Passwort=? WHERE Id=?;"
	tx, err := d.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(NeuesPasswort, Id)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) DeleteUser(Id string) error {
	query := "DELETE FROM User WHERE Id=?;"
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
