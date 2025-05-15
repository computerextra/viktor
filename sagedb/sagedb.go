package sagedb

import (
	"database/sql"
	"fmt"
	"regexp"
)

type SageDB struct {
	connectionString string
}

type Sg_Adressen struct {
	SG_Adressen_PK int
	Suchbegriff    sql.NullString
	KundNr         sql.NullString
	LiefNr         sql.NullString
	Homepage       sql.NullString
	Telefon1       sql.NullString
	Telefon2       sql.NullString
	Mobiltelefon1  sql.NullString
	Mobiltelefon2  sql.NullString
	EMail1         sql.NullString
	EMail2         sql.NullString
	KundUmsatz     sql.NullFloat64
	LiefUmsatz     sql.NullFloat64
}

type SearchResult struct {
	SG_Adressen_PK int
	Suchbegriff    *string
	KundNr         *string
	LiefNr         *string
	Homepage       *string
	Telefon1       *string
	Telefon2       *string
	Mobiltelefon1  *string
	Mobiltelefon2  *string
	EMail1         *string
	EMail2         *string
	KundUmsatz     *float64
	LiefUmsatz     *float64
}

type UserSearch struct {
	Name    sql.NullString
	Vorname sql.NullString
}

type User struct {
	Name    *string
	Vorname *string
}

func NewSage(server, db, user, password string, port int) *SageDB {

	return &SageDB{
		connectionString: fmt.Sprintf("server=%s;database=%s;user id=%s;password=%s;port=%d", server, db, user, password, port),
	}
}

func (s SageDB) Get(id string) (*User, error) {
	conn, err := sql.Open("sqlserver", s.connectionString)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	stmt, err := conn.Prepare("SELECT Name, Vorname FROM sg_adressen WHERE KundNr LIKE ?;")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var res UserSearch
	err = stmt.QueryRow(fmt.Sprintf("'%s'", id)).Scan(&res)
	if err != nil {
		return nil, err
	}

	var ret User
	if res.Name.Valid {
		ret.Name = &res.Name.String
	}
	if res.Vorname.Valid {
		ret.Vorname = &res.Vorname.String
	}

	return &ret, nil
}

func (s SageDB) Search(searchTerm string) ([]SearchResult, error) {
	if len(searchTerm) == 0 {
		return nil, fmt.Errorf("no searchterm")
	}

	reverse, err := regexp.MatchString("^(\\d|[+]49)", searchTerm)
	if err != nil {
		return nil, err
	}

	conn, err := sql.Open("sqlserver", s.connectionString)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	if reverse {
		query := fmt.Sprintf("'%%%s%%'", searchTerm)
		reversed, err := conn.Prepare(`
			SELECT SG_Adressen_PK, Suchbegriff,  KundNr, LiefNr, Homepage, Telefon1, Telefon2, Mobiltelefon1, Mobiltelefon2, EMail1, EMail2, KundUmsatz, LiefUmsatz 
			FROM sg_adressen WHERE 
			REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(Telefon1, ' ',''),'/',''),'-',''),'+49','0'),'(',''),')',''),',','')
			LIKE ? 
			OR 
			REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(Telefon2, ' ',''),'/',''),'-',''),'+49','0'),'(',''),')',''),',','')
			LIKE ?
			OR 
			REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(Mobiltelefon1, ' ',''),'/',''),'-',''),'+49','0'),'(',''),')',''),',','')
			LIKE ?
			OR 
			REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(Mobiltelefon2, ' ',''),'/',''),'-',''),'+49','0'),'(',''),')',''),',','')
			LIKE ?;`)
		if err != nil {
			return nil, err
		}
		defer reversed.Close()
		rows, err := reversed.Query(query, query, query, query)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var results []SearchResult
		for rows.Next() {
			var x Sg_Adressen
			if err := rows.Scan(&x); err != nil {
				return nil, err
			}
			results = append(results, SearchResult{
				SG_Adressen_PK: x.SG_Adressen_PK,
				Suchbegriff:    &x.Suchbegriff.String,
				KundNr:         &x.KundNr.String,
				LiefNr:         &x.LiefNr.String,
				Homepage:       &x.Homepage.String,
				Telefon1:       &x.Telefon1.String,
				Telefon2:       &x.Telefon2.String,
				Mobiltelefon1:  &x.Mobiltelefon1.String,
				Mobiltelefon2:  &x.Mobiltelefon2.String,
				EMail1:         &x.EMail1.String,
				EMail2:         &x.EMail2.String,
				KundUmsatz:     &x.KundUmsatz.Float64,
				LiefUmsatz:     &x.LiefUmsatz.Float64,
			})
		}
		return results, nil
	} else {
		query := fmt.Sprintf("'%%%s%%'", searchTerm)
		normal, err := conn.Prepare(`
		DECLARE @SearchWord NVARCHAR(30) 
		SET @SearchWord = N? 
		SELECT 
		SG_Adressen_PK, 
		Suchbegriff,  
		KundNr, 
		LiefNr, 
		Homepage, 
		Telefon1, 
		Telefon2, 
		Mobiltelefon1, 
		Mobiltelefon2, 
		EMail1, 
		EMail2, 
		KundUmsatz, 
		LiefUmsatz 
		FROM sg_adressen 
		WHERE Suchbegriff LIKE @SearchWord 
		OR KundNr LIKE @SearchWord 
		OR LiefNr LIKE @SearchWord;`)
		if err != nil {
			return nil, err
		}
		defer normal.Close()
		rows, err := normal.Query(query)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var results []SearchResult
		for rows.Next() {
			var x Sg_Adressen
			if err := rows.Scan(&x); err != nil {
				return nil, err
			}
			results = append(results, SearchResult{
				SG_Adressen_PK: x.SG_Adressen_PK,
				Suchbegriff:    &x.Suchbegriff.String,
				KundNr:         &x.KundNr.String,
				LiefNr:         &x.LiefNr.String,
				Homepage:       &x.Homepage.String,
				Telefon1:       &x.Telefon1.String,
				Telefon2:       &x.Telefon2.String,
				Mobiltelefon1:  &x.Mobiltelefon1.String,
				Mobiltelefon2:  &x.Mobiltelefon2.String,
				EMail1:         &x.EMail1.String,
				EMail2:         &x.EMail2.String,
				KundUmsatz:     &x.KundUmsatz.Float64,
				LiefUmsatz:     &x.LiefUmsatz.Float64,
			})
		}
		return results, nil
	}

}
