package handler

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"regexp"

	"github.com/computerextra/viktor/frontend"
	"github.com/computerextra/viktor/internal"
	"github.com/computerextra/viktor/internal/util/flash"
)

type SearchProps struct {
	Search string `schema:"search,required"`
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

func (h *Handler) Kunden(w http.ResponseWriter, r *http.Request) {
	uri := getPath(r.URL.Path)
	frontend.Kunden(uri, []internal.KundenResponse{}, "").Render(r.Context(), w)
}

func (h *Handler) SucheKunde(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxBodySize)
	r.ParseMultipartForm(10 << 20) // Max Header size (e.g. 10MB)
	var props SearchProps
	err := decoder.Decode(&props, r.PostForm)
	if err != nil {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		h.logger.Error("failed to parse formdata", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}

	connString := getSageConnectionString()
	conn, err := sql.Open("sqlserver", connString)
	if err != nil {
		h.logger.Error("failed to connect to sage", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}
	defer conn.Close()
	var results []internal.KundenResponse

	reverse, err := regexp.MatchString("^(\\d|[+]49)", props.Search)
	if err != nil {
		h.logger.Error("failed to matchstring", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}

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
			LIKE '%%%s%%'`, props.Search, props.Search, props.Search, props.Search,
		)

		rows, err := conn.Query(query, query, query, query)
		if err != nil {
			h.logger.Error("failed to matchstring", slog.Any("error", err))
			w.WriteHeader(http.StatusNoContent)
			return
		}
		defer rows.Close()

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
				h.logger.Error("failed to matchstring", slog.Any("error", err))
				w.WriteHeader(http.StatusNoContent)
				return
			}
			results = append(results, internal.KundenResponse{
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
		OR LiefNr LIKE @SearchWord;`, props.Search))
		if err != nil {
			h.logger.Error("failed to matchstring", slog.Any("error", err))
			w.WriteHeader(http.StatusNoContent)
			return
		}
		defer rows.Close()

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
				h.logger.Error("failed to matchstring", slog.Any("error", err))
				w.WriteHeader(http.StatusNoContent)
				return
			}
			results = append(results, internal.KundenResponse{
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

	}

	uri := getPath(r.URL.Path)
	frontend.Kunden(uri, results, props.Search).Render(r.Context(), w)
}
