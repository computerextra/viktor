package handler

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/computerextra/viktor/db"
	"github.com/computerextra/viktor/frontend"
	"github.com/computerextra/viktor/internal/util/flash"
)

type MitarbeiterProps struct {
	Name            string  `schema:"name,required"`       // custom name, must be supplied
	Short           *string `schema:"short"`               // custom name
	Sex             string  `schema:"sex,required"`        // custom name, must be supplied
	AbteilungId     *string `schema:"abteilungId"`         // custom name
	Image           bool    `schema:"image,default:false"` // custom name, boolean, default false
	Azubi           bool    `schema:"Azubi,default:false"` // custom name, boolean, default false
	Focus           *string `schema:"focus"`               // custom name
	Mail            *string `schema:"mail"`                // custom name
	Gruppenwahl     *string `schema:"Gruppenwahl"`         // custom name
	Geburtstag      *string `schema:"Geburtstag"`          // custom name
	HomeOffice      *string `schema:"HomeOffice"`          // custom name
	MobilPrivat     *string `schema:"Mobil_Privat"`        // custom name
	MobilBusiness   *string `schema:"Mobil_Business"`      // custom name
	TelefonBusiness *string `schema:"Telefon_Business"`    // custom name
	TelefonIntern1  *string `schema:"Telefon_Intern_1"`    // custom name
	TelefonIntern2  *string `schema:"Telefon_Intern_2"`    // custom name
	TelefonPrivat   *string `schema:"Telefon_Privat"`      // custom name
}

func (h *Handler) GetIndex(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	res, err := h.db.Mitarbeiter.FindMany().OrderBy(db.Mitarbeiter.Name.Order(db.SortOrderAsc)).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	uri := r.URL.Path

	frontend.Index(res, uri).Render(ctx, w)
}

func (h *Handler) getAbteilungen(ctx context.Context) ([]db.AbteilungModel, error) {
	res, err := h.db.Abteilung.FindMany().OrderBy(db.Abteilung.Name.Order(db.SortOrderAsc)).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (h *Handler) NewMitarbeiter(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	uri := getPath(r.URL.Path)
	abteilungen, err := h.getAbteilungen(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	frontend.NeuerMitarbeiter(abteilungen, uri).Render(r.Context(), w)
}

func (h *Handler) GetMitarbeitersWithAbteilung(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	res, err := h.db.Mitarbeiter.FindMany().With(
		db.Mitarbeiter.Abteilung.Fetch(),
	).OrderBy(db.Mitarbeiter.Name.Order(db.SortOrderAsc)).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	uri := getPath(r.URL.Path)
	frontend.MitarbeiterOverview(res, uri).Render(ctx, w)
}

func (h *Handler) GetMitarbeiter(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	res, err := h.db.Mitarbeiter.FindUnique(db.Mitarbeiter.ID.Equals(id)).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	abteilungen, err := h.getAbteilungen(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	uri := getPath(r.URL.Path)
	frontend.MitarbeiterBearbeiten(res, abteilungen, uri).Render(ctx, w)
}

func (h *Handler) CreateMitarbeiter(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	r.ParseMultipartForm(10 << 20) // Max Header size (e.g. 10MB)
	var mitarbeiter MitarbeiterProps
	err := decoder.Decode(&mitarbeiter, r.PostForm)
	if err != nil {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		h.logger.Error("failed to parse formdata", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}
	var day time.Time
	var bday *db.DateTime
	if mitarbeiter.Geburtstag != nil && len(*mitarbeiter.Geburtstag) > 0 {
		day, err = time.Parse("2006-01-02", *mitarbeiter.Geburtstag)
		if err != nil {
			sendError(w, h.logger, "failed to parse date", err)
		}
		bday = &day
	}

	var abteilungId *string
	if len(*mitarbeiter.AbteilungId) > 0 {
		abteilungId = mitarbeiter.AbteilungId
	}

	_, err = h.db.Mitarbeiter.CreateOne(
		db.Mitarbeiter.Name.Set(mitarbeiter.Name),
		db.Mitarbeiter.Short.SetIfPresent(mitarbeiter.Short),
		db.Mitarbeiter.AbteilungID.SetIfPresent(abteilungId),
		db.Mitarbeiter.Image.Set(mitarbeiter.Image),
		db.Mitarbeiter.Azubi.Set(mitarbeiter.Azubi),
		db.Mitarbeiter.Focus.SetIfPresent(mitarbeiter.Focus),
		db.Mitarbeiter.Mail.SetIfPresent(mitarbeiter.Mail),
		db.Mitarbeiter.Gruppenwahl.SetIfPresent(mitarbeiter.Gruppenwahl),
		db.Mitarbeiter.Geburtstag.SetIfPresent(bday),
		db.Mitarbeiter.HomeOffice.SetIfPresent(mitarbeiter.HomeOffice),
		db.Mitarbeiter.MobilPrivat.SetIfPresent(mitarbeiter.MobilPrivat),
		db.Mitarbeiter.MobilBusiness.SetIfPresent(mitarbeiter.MobilBusiness),
		db.Mitarbeiter.TelefonIntern1.SetIfPresent(mitarbeiter.TelefonIntern1),
		db.Mitarbeiter.TelefonIntern2.SetIfPresent(mitarbeiter.TelefonIntern2),
		db.Mitarbeiter.TelefonPrivat.SetIfPresent(mitarbeiter.TelefonPrivat),
		db.Mitarbeiter.TelefonBusiness.SetIfPresent(mitarbeiter.TelefonBusiness),
		db.Mitarbeiter.Sex.Set(mitarbeiter.Sex),
	).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	host := r.Host
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	uri := fmt.Sprintf("%s://%s/CMS/Mitarbeiter", scheme, host)
	http.Redirect(w, r, uri, http.StatusFound)
}

func (h *Handler) UpdateMitarbeiter(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	r.ParseMultipartForm(10 << 20) // Max Header size (e.g. 10MB)
	id := r.PathValue("id")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	var mitarbeiter MitarbeiterProps
	err := decoder.Decode(&mitarbeiter, r.PostForm)
	if err != nil {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		h.logger.Error("failed to parse formdata", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}
	var day time.Time
	var bday *db.DateTime
	if mitarbeiter.Geburtstag != nil && len(*mitarbeiter.Geburtstag) > 0 {
		day, err = time.Parse("2006-01-02", *mitarbeiter.Geburtstag)
		if err != nil {
			sendError(w, h.logger, "failed to parse date", err)
		}
		bday = &day
	}

	var abteilungId *string
	if len(*mitarbeiter.AbteilungId) > 0 {
		abteilungId = mitarbeiter.AbteilungId
	}

	_, err = h.db.Mitarbeiter.FindUnique(db.Mitarbeiter.ID.Equals(id)).Update(
		db.Mitarbeiter.Name.Set(mitarbeiter.Name),
		db.Mitarbeiter.Short.SetIfPresent(mitarbeiter.Short),
		db.Mitarbeiter.AbteilungID.SetIfPresent(abteilungId),
		db.Mitarbeiter.Image.Set(mitarbeiter.Image),
		db.Mitarbeiter.Azubi.Set(mitarbeiter.Azubi),
		db.Mitarbeiter.Focus.SetIfPresent(mitarbeiter.Focus),
		db.Mitarbeiter.Mail.SetIfPresent(mitarbeiter.Mail),
		db.Mitarbeiter.Gruppenwahl.SetIfPresent(mitarbeiter.Gruppenwahl),
		db.Mitarbeiter.Geburtstag.SetIfPresent(bday),
		db.Mitarbeiter.HomeOffice.SetIfPresent(mitarbeiter.HomeOffice),
		db.Mitarbeiter.MobilPrivat.SetIfPresent(mitarbeiter.MobilPrivat),
		db.Mitarbeiter.MobilBusiness.SetIfPresent(mitarbeiter.MobilBusiness),
		db.Mitarbeiter.TelefonIntern1.SetIfPresent(mitarbeiter.TelefonIntern1),
		db.Mitarbeiter.TelefonIntern2.SetIfPresent(mitarbeiter.TelefonIntern2),
		db.Mitarbeiter.TelefonPrivat.SetIfPresent(mitarbeiter.TelefonPrivat),
		db.Mitarbeiter.TelefonBusiness.SetIfPresent(mitarbeiter.TelefonBusiness),
		db.Mitarbeiter.Sex.Set(mitarbeiter.Sex),
	).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	host := r.Host
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	uri := fmt.Sprintf("%s://%s/CMS/Mitarbeiter", scheme, host)
	http.Redirect(w, r, uri, http.StatusFound)
}

func (h *Handler) DeleteMitarbeiter(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	_, err := h.db.Mitarbeiter.FindUnique(db.Mitarbeiter.ID.Equals(id)).Delete().Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	host := r.Host
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	uri := fmt.Sprintf("%s://%s/CMS/Mitarbeiter", scheme, host)
	http.Redirect(w, r, uri, http.StatusFound)
}
