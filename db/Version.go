package db

import "database/sql"

type VersionModel struct {
	Id      int     `json:"id"`
	Current float32 `json:"current_version"`
}

type VersionParams struct {
	Current float32 `json:"current_version"`
}

// CRUD
func (d Database) createVersion(params VersionParams) error {
	db, err := sql.Open("sqlite3", d.ConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO Version VALUES(?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(params)
	return err
}

func (d Database) readVersion(id *int) ([]VersionModel, error) {
	if *id > 0 {
		res, err := readOneVersion(*id, d.ConnectionString)
		if err != nil {
			return nil, err
		}
		return []VersionModel{*res}, nil
	} else {
		return readAllVersion(d.ConnectionString)
	}
}

func readOneVersion(id int, connString string) (*VersionModel, error) {
	db, err := sql.Open("sqlite3", connString)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	stmt, err := db.Prepare("SELECT * FROM Version WHERE id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var ap VersionModel
	err = stmt.QueryRow(id).Scan(&ap)
	if err != nil {
		return nil, err
	}
	return &ap, nil
}

func readAllVersion(connString string) ([]VersionModel, error) {
	db, err := sql.Open("sqlite3", connString)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	stmt, err := db.Prepare("SELECT * FROM Version ORDER BY Name ASC")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var aps []VersionModel
	for rows.Next() {
		var a VersionModel
		err = rows.Scan(&a)
		if err != nil {
			return nil, err
		}
		aps = append(aps, a)
	}
	return aps, nil
}

func (d Database) updateVersion(id int, params VersionParams) error {
	db, err := sql.Open("sqlite3", d.ConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE Version set curent_version=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(params, id)
	return err
}

func (d Database) deleteVersion(id int) error {
	db, err := sql.Open("sqlite3", d.ConnectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM Version WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}
