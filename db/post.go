package db

import (
	"time"

	"github.com/lucsky/cuid"
)

type Post struct {
	Id          string
	KanbanId    string
	Name        string
	Description *string
	Status      string
	Importance  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (d *Database) CreatePost(
	KanbanId string,
	Name string,
	Description *string,
	Status string,
	Importance string,
) (*string, error) {
	query := "INSERT INTO Post(Id, KanbanId, Name, Description, Status, Importance) values(?,?,?,?,?,?);"
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
	_, err = stmt.Exec(id, KanbanId, Name, Description, Status, Importance)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (d *Database) GetPost(id string) (*Post, error) {
	query := "SELECT Id, KanbanId, Name, Description, Status, Importance, CreatedAt, UpdatedAt from Post WHERE Id=?;"
	stmt, err := d.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var res Post
	err = stmt.QueryRow(id).Scan(&res.Id, &res.KanbanId, &res.Name, &res.Description, &res.Status, &res.Importance, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (d *Database) GetAllPostFromBoard(BoardId string) ([]Post, error) {
	query := "SELECT Id, KanbanId, Name, Description, Status, Importance, CreatedAt, UpdatedAt from Post WHERE KanbanId=? ORDER BY UpdatedAt ASC;"
	var res []Post

	rows, err := d.DB.Query(query, BoardId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var x Post
		err = rows.Scan(&x.Id, &x.KanbanId, &x.Name, &x.Description, &x.Status, &x.Importance, &x.CreatedAt, &x.UpdatedAt)
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

func (d *Database) UpdatePost(
	Id string,
	Name string,
	Description *string,
	Status string,
	Importance string,
) error {
	query := "UPDATE Post SET Name=?, Description=?, Status=?, Importance=?, UpdatedAt=(DATETIME('now')) WHERE Id=?;"
	tx, err := d.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(Name, Description, Status, Importance, Id)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) DeletePost(Id string) error {
	query := "DELETE FROM Post WHERE Id=?;"
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

func (d *Database) DeletePostsFromBoard(BoardId string) error {
	query := "DELETE FROM Post WHERE KanbanId=?;"
	tx, err := d.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(BoardId)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
