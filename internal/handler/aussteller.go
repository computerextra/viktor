package handler

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/computerextra/viktor/db"
	"github.com/computerextra/viktor/frontend"
	"github.com/computerextra/viktor/internal"
	"github.com/computerextra/viktor/internal/util/flash"
	"github.com/jlaffaye/ftp"

	_ "github.com/denisenkom/go-mssqldb"
)

type AusstellerProps struct {
	Artikelnummer string `schema:"Artikelnummer,required"`
	Link          string `schema:"Link,required"`
}

func (h *Handler) Aussteller(w http.ResponseWriter, r *http.Request) {
	frontend.Aussteller(getPath(r.URL.Path), false, false, "").Render(r.Context(), w)
}

func (h *Handler) SyncAussteller(w http.ResponseWriter, r *http.Request) {
	err := sync(h.db, r.Context())
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		h.logger.Error("failed to parse formdata", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}
	frontend.Aussteller(getPath(r.URL.Path), false, true, "").Render(r.Context(), w)
}

func (h *Handler) UploadAussteller(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 10) // Max Header size: 32 MB

	file, handler, err := r.FormFile("file-input")
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		h.logger.Error("failed to parse formdata", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}
	defer file.Close()

	FTP_SERVER, ok := os.LookupEnv("FTP_SERVER")
	if !ok {
		h.logger.Error("failed to read from env: FTP_SERVER")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	FTP_USER, ok := os.LookupEnv("FTP_USER")
	if !ok {
		h.logger.Error("failed to read from env: FTP_USER")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	FTP_PASS, ok := os.LookupEnv("FTP_PASS")
	if !ok {
		h.logger.Error("failed to read from env: FTP_PASS")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	FTP_PORT, ok := os.LookupEnv("FTP_PORT")
	if !ok {
		h.logger.Error("failed to read from env: FTP_PORT")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	FTP_UPLOAD_PATH, ok := os.LookupEnv("FTP_UPLOAD_PATH")
	if !ok {
		h.logger.Error("failed to read from env: FTP_UPLOAD_PATH")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	ftpClient, err := ftp.Dial(fmt.Sprintf("%s:%s", FTP_SERVER, FTP_PORT), ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		h.logger.Error("failed connect to ftp", slog.Any("error", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer ftpClient.Quit()

	if err := ftpClient.Login(FTP_USER, FTP_PASS); err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		h.logger.Error("failed login to ftp", slog.Any("error", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	remotefile := FTP_UPLOAD_PATH + handler.Filename
	if err := ftpClient.Stor(remotefile, file); err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		h.logger.Error("failed upload file", slog.Any("error", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	domain := "https://bilder.computer-extra.de/data"

	frontend.Aussteller(getPath(r.URL.Path), false, false, fmt.Sprintf("%s/%s", domain, handler.Filename)).Render(r.Context(), w)
}

func (h *Handler) UpdateAussteller(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 10) // Max Header size: 32 MB
	var props AusstellerProps
	err := decoder.Decode(&props, r.PostForm)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		h.logger.Error("failed to parse formdata", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}

	postUrl := "https://aussteller.computer-extra.de/php/update.php"
	data, err := json.Marshal(props)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		h.logger.Error("failed marshal formdata", slog.Any("error", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, err := http.Post(postUrl, "application/json", bytes.NewBuffer(data))
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		h.logger.Error("failed post request", slog.Any("error", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	frontend.Aussteller(getPath(r.URL.Path), true, false, "").Render(r.Context(), w)
}

func sync(database *db.PrismaClient, ctx context.Context) error {
	connString := getSageConnectionString()

	conn, err := sql.Open("sqlserver", connString)
	if err != nil {
		return err
	}
	defer conn.Close()
	sage_query := "select sg_auf_artikel.SG_AUF_ARTIKEL_PK, sg_auf_artikel.ARTNR, sg_auf_artikel.SUCHBEGRIFF, sg_auf_artikel.ZUSTEXT1, sg_auf_vkpreis.PR01 FROM sg_auf_artikel INNER JOIN sg_auf_vkpreis ON sg_auf_artikel.SG_AUF_ARTIKEL_PK = sg_auf_vkpreis.SG_AUF_ARTIKEL_FK"

	rows, err := conn.Query(sage_query)
	if err != nil {
		return err
	}
	defer rows.Close()

	var Sage []internal.AusstellerArtikel
	for rows.Next() {
		var Id sql.NullInt64
		var Artikelnummer sql.NullString
		var Artikelname sql.NullString
		var Specs sql.NullString
		var Preis sql.NullFloat64
		if err := rows.Scan(&Id, &Artikelnummer, &Artikelname, &Specs, &Preis); err != nil {
			return err
		}
		if Id.Valid && Artikelnummer.Valid && Artikelname.Valid && Specs.Valid && Preis.Valid {
			var tmp internal.AusstellerArtikel
			tmp.Id = int(Id.Int64)
			tmp.Artikelnummer = Artikelnummer.String
			tmp.Artikelname = Artikelname.String
			tmp.Preis = Preis.Float64
			tmp.Specs = Specs.String
			Sage = append(Sage, tmp)
		}
	}

	rawUpsertQuery := "INSERT INTO Aussteller (id, Artikelnummer, Artikelname, Specs, Preis) VALUES"

	for idx, item := range Sage {
		if idx == len(Sage)-1 {
			rawUpsertQuery = fmt.Sprintf("%s (%d, '%s', '%s', '%s', %.2f)", rawUpsertQuery, item.Id, item.Artikelnummer, strings.ReplaceAll(item.Artikelname, "'", "\""), strings.ReplaceAll(item.Specs, "'", "\""), item.Preis)
		} else {
			rawUpsertQuery = fmt.Sprintf("%s (%d, '%s', '%s', '%s', %.2f),", rawUpsertQuery, item.Id, item.Artikelnummer, strings.ReplaceAll(item.Artikelname, "'", "\""), strings.ReplaceAll(item.Specs, "'", "\""), item.Preis)
		}
	}

	rawUpsertQuery = fmt.Sprintf("%s ON DUPLICATE KEY UPDATE Artikelnummer = VALUES(Artikelnummer), Artikelname = VALUES(Artikelname), Specs = VALUES(Specs), Preis = VALUES(Preis);", rawUpsertQuery)

	_, err = database.Prisma.ExecuteRaw(rawUpsertQuery).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
