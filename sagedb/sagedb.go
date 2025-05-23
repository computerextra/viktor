package sagedb

import (
	"database/sql"
	"fmt"
	"regexp"

	_ "github.com/denisenkom/go-mssqldb"
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

	rows, err := conn.Query(fmt.Sprintf("SELECT Name, Vorname FROM sg_adressen WHERE KundNr LIKE '%s';", id))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res UserSearch
	for rows.Next() {
		var name sql.NullString
		var vorname sql.NullString
		if err := rows.Scan(&name, &vorname); err != nil {
			return nil, err
		}
		if name.Valid {
			res.Name = name
		}
		if vorname.Valid {
			res.Vorname = vorname
		}
		if err := rows.Err(); err != nil {
			return nil, err
		}
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
		query := fmt.Sprintf(`
			SELECT SG_Adressen_PK, Suchbegriff,  KundNr, LiefNr, Homepage, Telefon1, Telefon2, Mobiltelefon1, Mobiltelefon2, EMail1, EMail2, KundUmsatz, LiefUmsatz 
			FROM sg_adressen WHERE 
			REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(Telefon1, ' ',''),'/',''),'-',''),'+49','0'),'(',''),')',''),',','')
			LIKE '%%%s%%' 
			OR 
			REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(Telefon2, ' ',''),'/',''),'-',''),'+49','0'),'(',''),')',''),',','')
			LIKE '%%%s%%' 
			OR 
			REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(Mobiltelefon1, ' ',''),'/',''),'-',''),'+49','0'),'(',''),')',''),',','')
			LIKE '%%%s%%' 
			OR 
			REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(REPLACE(Mobiltelefon2, ' ',''),'/',''),'-',''),'+49','0'),'(',''),')',''),',','')
			LIKE '%%%s%%'`, searchTerm, searchTerm, searchTerm, searchTerm,
		)

		rows, err := conn.Query(query, query, query, query)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var results []SearchResult
		for rows.Next() {
			var x Sg_Adressen
			if err := rows.Scan(
				&x.SG_Adressen_PK,
				&x.Suchbegriff,
				&x.KundNr,
				&x.LiefNr,
				&x.Homepage,
				&x.Telefon1,
				&x.Telefon2,
				&x.Mobiltelefon1,
				&x.Mobiltelefon2,
				&x.EMail1,
				&x.EMail2,
				&x.KundUmsatz,
				&x.LiefUmsatz,
			); err != nil {
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
		rows, err := conn.Query(fmt.Sprintf(`
		DECLARE @SearchWord NVARCHAR(30) 
		SET @SearchWord = N'%%%s%%' 
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
		OR LiefNr LIKE @SearchWord;`, searchTerm))
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var results []SearchResult
		for rows.Next() {
			var x Sg_Adressen
			if err := rows.Scan(
				&x.SG_Adressen_PK,
				&x.Suchbegriff,
				&x.KundNr,
				&x.LiefNr,
				&x.Homepage,
				&x.Telefon1,
				&x.Telefon2,
				&x.Mobiltelefon1,
				&x.Mobiltelefon2,
				&x.EMail1,
				&x.EMail2,
				&x.KundUmsatz,
				&x.LiefUmsatz,
			); err != nil {
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
