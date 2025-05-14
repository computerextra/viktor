package db

import (
	"database/sql"

	"github.com/lucsky/cuid"
)

type UserModel struct {
	Id            string `json:"id"`
	Password      string `json:"Password"`
	Mail          string `json:"Mail"`
	Active        bool   `json:"Active"`
	MitarbeiterId string `json:"mitarbeiterId"`
}

type UserParams struct {
	Password string `json:"Password"`
	Mail     string `json:"Mail"`
	Active   bool   `json:"Active"`
}

// CRUD
func (d Database) createUser(params UserParams) error {
	db, err := sql.Open("sqlite3", d.ConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	// Find Mitarbeiter
	var ma MitarbeiterModel
	stmt, err := db.Prepare("SELECT * FROM Mitarbeiter WHERE Email=?")
	if err != nil {
		return err
	}
	err = stmt.QueryRow(params.Mail).Scan(&ma)
	if err != nil {
		return err
	}

	stmt, err = db.Prepare("INSERT INTO User VALUES(?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	x := UserModel{
		Id:            cuid.New(),
		Password:      params.Password,
		Mail:          params.Mail,
		Active:        params.Active,
		MitarbeiterId: ma.Id,
	}

	_, err = stmt.Exec(x)
	return err
}

func (d Database) readUser(id *string) ([]UserModel, error) {
	if len(*id) > 0 {
		res, err := readOneUser(*id, d.ConnectionString)
		if err != nil {
			return nil, err
		}
		return []UserModel{*res}, nil
	} else {
		return readAllUser(d.ConnectionString)
	}
}

func readOneUser(id string, connString string) (*UserModel, error) {
	db, err := sql.Open("sqlite3", connString)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	stmt, err := db.Prepare("SELECT * FROM User WHERE id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var ap UserModel
	err = stmt.QueryRow(id).Scan(&ap)
	if err != nil {
		return nil, err
	}
	return &ap, nil
}

func readAllUser(connString string) ([]UserModel, error) {
	db, err := sql.Open("sqlite3", connString)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	stmt, err := db.Prepare("SELECT * FROM User ORDER BY Mail ASC")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var aps []UserModel
	for rows.Next() {
		var a UserModel
		err = rows.Scan(&a)
		if err != nil {
			return nil, err
		}
		aps = append(aps, a)
	}
	return aps, nil
}

func (d Database) updateUser(id string, params UserParams) error {
	db, err := sql.Open("sqlite3", d.ConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE User set Password=?, Active=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(params, id)
	return err
}

func (d Database) deleteUser(id string) error {
	db, err := sql.Open("sqlite3", d.ConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM User WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}
