package archive

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Archive struct {
	ConnectionString string
}

type ArchiveResult struct {
	Id    int32
	Title string
	Body  string
}

func NewArchive(DatabaseUrl string) *Archive {
	return &Archive{
		ConnectionString: DatabaseUrl,
	}
}

func (a Archive) Search(searchTerm string) ([]ArchiveResult, error) {
	conn, err := sql.Open("mysql", a.ConnectionString)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	stmt, err := conn.Prepare("SELECT id, title, body FROM pdfs WHERE body LIKE ? OR title LIKE ?;")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	search := fmt.Sprintf("%%%s%%", searchTerm)

	rows, err := stmt.Query(search, search)
	if err != nil {
		return nil, err
	}

	var results []ArchiveResult

	for rows.Next() {
		var x ArchiveResult
		if err := rows.Scan(&x.Id, &x.Title, &x.Body); err != nil {
			return nil, err
		}
		results = append(results, x)
	}
	return results, nil
}

func (a Archive) Get(id int32) (*ArchiveResult, error) {
	conn, err := sql.Open("mysql", a.ConnectionString)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	stmt, err := conn.Prepare("SELECT id, title FROM pdfs WHERE id=?;")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var result ArchiveResult
	err = stmt.QueryRow(id).Scan(&result.Id, &result.Title)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
