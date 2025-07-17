package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/computerextra/viktor/frontend"
	"github.com/computerextra/viktor/internal"
	"github.com/computerextra/viktor/internal/util/flash"

	_ "github.com/mattn/go-adodb"
)

func (h *Handler) Label(w http.ResponseWriter, r *http.Request) {
	frontend.Label(getPath(r.URL.Path), false).Render(r.Context(), w)
}

func (h *Handler) SyncLabel(w http.ResponseWriter, r *http.Request) {
	sageItems, err := getAllProductsFromSage()
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}
	label, err := readAccessDb()
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}
	err = syncDB(sageItems, label)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}

	frontend.Label(getPath(r.URL.Path), true).Render(r.Context(), w)
}

func readAccessDb() ([]internal.AccessArtikel, error) {
	accessDb, ok := os.LookupEnv("ACCESS_DB")
	if !ok {
		return nil, fmt.Errorf("could not load ACCESS_DB from env")
	}

	conn, err := sql.Open("adodb", fmt.Sprintf("Provider=Microsoft.ACE.OLEDB.12.0;Data Source=%s;", accessDb))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var artikel []internal.AccessArtikel

	rows, err := conn.Query("SELECT ID, Artikelnummer, Artikeltext, Preis FROM Artikel")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var art internal.AccessArtikel
		if err := rows.Scan(&art.Id, &art.Artikelnummer, &art.Artikeltext, &art.Preis); err != nil {
			return nil, err
		}
		artikel = append(artikel, art)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return artikel, nil
}

func syncDB(sage []Artikel, label []internal.AccessArtikel) error {
	var updates []internal.AccessArtikel
	var create []internal.AccessArtikel

	for _, sageItem := range sage {
		var found bool = false
		for _, labelItem := range label {
			if sageItem.Id == labelItem.Id {
				found = true
				break
			}
		}
		art := internal.AccessArtikel{
			Id:            sageItem.Id,
			Artikelnummer: sageItem.Artikelnummer,
			Preis:         sageItem.Preis,
			Artikeltext:   sageItem.Suchbegriff,
		}
		if found {
			updates = append(updates, art)
		} else {
			create = append(create, art)
		}
	}

	err := insert(create)
	if err != nil {
		return err
	}
	err = update(updates)
	if err != nil {
		return err
	}
	return nil
}

func insert(items []internal.AccessArtikel) error {
	accessDb, ok := os.LookupEnv("ACCESS_DB")
	if !ok {
		return fmt.Errorf("could not load ACCESS_DB from env")
	}

	conn, err := sql.Open("adodb", fmt.Sprintf("Provider=Microsoft.ACE.OLEDB.12.0;Data Source=%s;", accessDb))
	if err != nil {
		return err
	}
	defer conn.Close()

	stmt, err := conn.Prepare("INSERT INTO Artikel (ID, Artikelnummer, Artikeltext, Preis) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, item := range items {
		if _, err := stmt.Exec(item.Id, item.Artikelnummer, strings.ReplaceAll(item.Artikeltext, "'", "\""), item.Preis); err != nil {
			return err
		}
	}

	return nil
}

func update(items []internal.AccessArtikel) error {
	accessDb, ok := os.LookupEnv("ACCESS_DB")
	if !ok {
		return fmt.Errorf("could not load ACCESS_DB from env")
	}

	conn, err := sql.Open("adodb", fmt.Sprintf("Provider=Microsoft.ACE.OLEDB.12.0;Data Source=%s;", accessDb))
	if err != nil {
		return err
	}
	defer conn.Close()

	stmt, err := conn.Prepare("UPDATE Artikel SET Artikelnummer=?, Artikeltext=?, Preis=? where ID=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, item := range items {
		if _, err := stmt.Exec(item.Artikelnummer, strings.ReplaceAll(item.Artikeltext, "'", "\""), item.Preis, item.Id); err != nil {
			return err
		}
	}

	return nil
}
