package db

import (
	"time"

	"github.com/lucsky/cuid"
)

type Kanban struct {
	Id        string
	Name      string
	UserId    string
	Posts     []Post
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (d *Database) CreateKanban(
	Name string,
	UserId string,
) (*string, error) {
	query := "INSERT INTO Kanban(Id, Name, UserId) values(?,?,?);"
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
	_, err = stmt.Exec(id, Name, UserId)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (d *Database) GetKanban(id string) (*Kanban, error) {
	query := "SELECT Id, Name, UserId, CreatedAt, UpdatedAt FROM Kanban WHERE Id=?;"
	stmt, err := d.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var res Kanban
	err = stmt.QueryRow(id).Scan(&res.Id, &res.Name, &res.UserId, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return nil, err
	}

	posts, err := d.GetAllPostFromBoard(id)
	if err != nil {
		return nil, err
	}
	res.Posts = posts
	return &res, nil
}

func (d *Database) GetAllKanbans(UserId string) ([]Kanban, error) {
	query := "SELECT Id, Name, UserId, CreatedAt, UpdatedAt from Kanban WHERE UserId=? ORDER BY UpdatedAt ASC;"
	var res []Kanban

	rows, err := d.DB.Query(query, UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var x Kanban
		err = rows.Scan(&x.Id, &x.Name, &x.UserId, &x.CreatedAt, &x.UpdatedAt)
		if err != nil {
			return nil, err
		}
		posts, err := d.GetAllPostFromBoard(x.Id)
		if err != nil {
			return nil, err
		}
		x.Posts = posts
		res = append(res, x)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (d *Database) UpdateKanban(
	Id string,
	Name string,
) error {
	query := "UPDATE Kanban SET Name=?, UpdatedAt=(DATETIME('now')) WHERE Id=?;"
	tx, err := d.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(Name, Id)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) DeleteKanban(Id string) error {
	query := "DELETE FROM Kanban WHERE Id=?;"
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

	err = d.DeletePostsFromBoard(Id)
	if err != nil {
		return err
	}

	return nil
}
