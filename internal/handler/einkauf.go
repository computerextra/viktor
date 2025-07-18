package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/computerextra/viktor/db"
	"github.com/computerextra/viktor/frontend"
	"github.com/computerextra/viktor/internal/util/flash"
	"github.com/jlaffaye/ftp"
)

type EinkaufProps struct {
	Abo    bool    `schema:"Abo,default:false"`
	Paypal bool    `schema:"Paypal,default:false"`
	Dinge  string  `schema:"Dinge,required"`
	Geld   *string `schema:"Geld"`
	Pfand  *string `schema:"Pfand"`
	Bild1  []byte  `schema:"bild1"`
	Bild2  []byte  `schema:"bild2"`
	Bild3  []byte  `schema:"bild3"`
}

func (h *Handler) GetEinkauf(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	mitarbeiter, err := h.db.Mitarbeiter.FindUnique(db.Mitarbeiter.ID.Equals(id)).With(
		db.Mitarbeiter.Einkauf.Fetch(),
	).Exec(ctx)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}

	uri := getPath(r.URL.Path)

	frontend.EinkaufEingabe(mitarbeiter, uri).Render(ctx, w)
}

func (h *Handler) GetListe(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	loc, _ := time.LoadLocation("Europe/Berlin")
	abgeschickt := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, loc)

	einkauf, err := h.db.Einkauf.FindMany(db.Einkauf.Or(
		db.Einkauf.And(
			db.Einkauf.Abgeschickt.Lte(abgeschickt),
			db.Einkauf.Abgeschickt.Gt(abgeschickt.AddDate(0, 0, -1)),
		),
		db.Einkauf.And(
			db.Einkauf.Abonniert.Equals(true),
			db.Einkauf.Abgeschickt.Lte(abgeschickt),
		),
	)).With(
		db.Einkauf.Mitarbeiter.Fetch(),
	).OrderBy(
		db.Einkauf.Abgeschickt.Order(
			db.SortOrderDesc,
		),
	).Exec(ctx)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}

	mitarbeiter, err := h.db.Mitarbeiter.FindMany().OrderBy(db.Mitarbeiter.Name.Order(db.SortOrderAsc)).Exec(ctx)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}

	uri := getPath(r.URL.Path)

	frontend.Einkauf(einkauf, mitarbeiter, uri).Render(ctx, w)
}

func (h *Handler) SkipEinkauf(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	_, err := h.db.Einkauf.FindUnique(db.Einkauf.ID.Equals(id)).Update(
		db.Einkauf.Abgeschickt.Set(time.Now().AddDate(0, 0, 1)),
	).Exec(ctx)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}
	host := r.Host
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	uri := fmt.Sprintf("%s://%s/Einkauf", scheme, host)
	http.Redirect(w, r, uri, http.StatusFound)
}

func (h *Handler) DeleteEinkauf(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	_, err := h.db.Einkauf.FindUnique(db.Einkauf.ID.Equals(id)).Update(
		db.Einkauf.Abgeschickt.Set(time.Now().AddDate(-1, 0, 0)),
		db.Einkauf.Abonniert.Set(false),
	).Exec(ctx)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}
	host := r.Host
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	uri := fmt.Sprintf("%s://%s/Einkauf", scheme, host)
	http.Redirect(w, r, uri, http.StatusFound)
}

func (h *Handler) UpdateEinkauf(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	mitarbeiterId := r.PathValue("id")
	r.ParseMultipartForm(32 << 10) // Max Header size: 32 MB

	var einkauf EinkaufProps
	err := decoder.Decode(&einkauf, r.PostForm)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		h.logger.Error("failed to parse formdata", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}

	var file1 *string
	var file2 *string
	var file3 *string

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

	bild1FormData, bild1Handler, err := r.FormFile("bild1")
	if err == nil {
		splitted := strings.Split(bild1Handler.Filename, ".")
		fileType := splitted[len(splitted)-1]
		filename := fmt.Sprintf("%s-1.%s", mitarbeiterId, fileType)
		path := fmt.Sprintf("https://bilder.computer-extra.de/data/%s", filename)

		remotefile := FTP_UPLOAD_PATH + filename

		ftpClient.Delete(remotefile)

		if err := ftpClient.Stor(remotefile, bild1FormData); err != nil {
			flash.SetFlashMessage(w, "error", err.Error())
			h.logger.Error("failed upload file", slog.Any("error", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		file1 = &path
		bild1FormData.Close()
	}

	bild2FormData, bild2Handler, err := r.FormFile("bild2")
	if err == nil {
		splitted := strings.Split(bild2Handler.Filename, ".")
		fileType := splitted[len(splitted)-1]
		filename := fmt.Sprintf("%s-2.%s", mitarbeiterId, fileType)
		path := fmt.Sprintf("https://bilder.computer-extra.de/data/%s", filename)

		remotefile := FTP_UPLOAD_PATH + filename

		ftpClient.Delete(remotefile)

		if err := ftpClient.Stor(remotefile, bild2FormData); err != nil {
			flash.SetFlashMessage(w, "error", err.Error())
			h.logger.Error("failed upload file", slog.Any("error", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		file2 = &path
		bild2FormData.Close()
	}

	bild3FormData, bild3Handler, err := r.FormFile("bild3")
	if err == nil {
		splitted := strings.Split(bild3Handler.Filename, ".")
		fileType := splitted[len(splitted)-1]
		filename := fmt.Sprintf("%s-3.%s", mitarbeiterId, fileType)
		path := fmt.Sprintf("https://bilder.computer-extra.de/data/%s", filename)

		remotefile := FTP_UPLOAD_PATH + filename

		ftpClient.Delete(remotefile)

		if err := ftpClient.Stor(remotefile, bild3FormData); err != nil {
			flash.SetFlashMessage(w, "error", err.Error())
			h.logger.Error("failed upload file", slog.Any("error", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		file3 = &path
		bild3FormData.Close()
	}

	loc, _ := time.LoadLocation("Europe/Berlin")
	abgeschickt := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, loc)

	mitarbeiter, err := h.db.Mitarbeiter.FindUnique(db.Mitarbeiter.ID.Equals(mitarbeiterId)).Exec(ctx)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}
	einkauId, _ := mitarbeiter.EinkaufID()

	var bild1, bild2, bild3 string = "", "", ""

	if file1 != nil {
		bild1 = *file1
	}
	if file2 != nil {
		bild2 = *file2
	}
	if file3 != nil {
		bild3 = *file3
	}

	_, err = h.db.Einkauf.UpsertOne(
		db.Einkauf.ID.Equals(einkauId),
	).Create(
		db.Einkauf.Dinge.Set(einkauf.Dinge),
		db.Einkauf.Abgeschickt.Set(abgeschickt),
		db.Einkauf.Abonniert.Set(einkauf.Abo),
		db.Einkauf.Paypal.Set(einkauf.Paypal),
		db.Einkauf.Geld.SetIfPresent(einkauf.Geld),
		db.Einkauf.Pfand.SetIfPresent(einkauf.Pfand),
		db.Einkauf.Bild1.Set(bild1),
		db.Einkauf.Bild2.Set(bild2),
		db.Einkauf.Bild3.Set(bild3),
		db.Einkauf.Mitarbeiter.Link(
			db.Mitarbeiter.ID.Equals(mitarbeiterId),
		),
	).Update(
		db.Einkauf.Dinge.Set(einkauf.Dinge),
		db.Einkauf.Abgeschickt.Set(abgeschickt),
		db.Einkauf.Abonniert.Set(einkauf.Abo),
		db.Einkauf.Paypal.Set(einkauf.Paypal),
		db.Einkauf.Bild1.Set(bild1),
		db.Einkauf.Bild2.Set(bild2),
		db.Einkauf.Bild3.Set(bild3),
		db.Einkauf.Geld.SetIfPresent(einkauf.Geld),
		db.Einkauf.Pfand.SetIfPresent(einkauf.Pfand),
	).Exec(ctx)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}

	host := r.Host
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	uri := fmt.Sprintf("%s://%s/Einkauf", scheme, host)
	http.Redirect(w, r, uri, http.StatusFound)
}
