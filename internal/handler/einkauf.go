package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/computerextra/viktor/db"
	"github.com/computerextra/viktor/frontend"
	"github.com/computerextra/viktor/internal/util/flash"
)

type EinkaufProps struct {
	Abo    string  `schema:"Abo,default:false"`
	Paypal string  `schema:"Paypal,default:false"`
	Dinge  string  `schema:"Dinge,required"`
	Geld   *string `schema:"Geld"`
	Pfand  *string `schema:"Pfand"`
}

// TODO: Image Upload

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
		sendQueryError(w, h.logger, err)
	}

	uri := getPath(r.URL.Path)

	frontend.EinkaufEingabe(mitarbeiter, uri).Render(ctx, w)
}

func (h *Handler) GetListe(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	einkauf, err := h.db.Einkauf.FindMany(db.Einkauf.Or(
		db.Einkauf.And(
			db.Einkauf.Abgeschickt.Lte(time.Now()),
			db.Einkauf.Abgeschickt.Gte(time.Now().AddDate(0, 0, -1)),
		),
		db.Einkauf.And(
			db.Einkauf.Abonniert.Equals(true),
			db.Einkauf.Abgeschickt.Lte(time.Now()),
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
	r.ParseMultipartForm(10 << 20) // Max Header size (e.g. 10MB)
	h.logger.Info("Einkauf", slog.Any("data", r.PostForm))
	var einkauf EinkaufProps
	err := decoder.Decode(&einkauf, r.PostForm)
	if err != nil {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		h.logger.Error("failed to parse formdata", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}
	h.logger.Info("Einkauf", slog.Any("data", r.PostForm))

	mitarbeiter, err := h.db.Mitarbeiter.FindUnique(db.Mitarbeiter.ID.Equals(mitarbeiterId)).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	einkauId, _ := mitarbeiter.EinkaufID()

	var Abo bool = false
	var Paypal bool = false

	if einkauf.Abo == "true" {
		Abo = true
	}
	if einkauf.Paypal == "true" {
		Paypal = true
	}

	_, err = h.db.Einkauf.UpsertOne(
		db.Einkauf.ID.Equals(einkauId),
	).Create(
		db.Einkauf.Dinge.Set(einkauf.Dinge),
		db.Einkauf.Abgeschickt.Set(time.Now()),
		db.Einkauf.Abonniert.Set(Abo),
		db.Einkauf.Paypal.Set(Paypal),
		db.Einkauf.Geld.SetIfPresent(einkauf.Geld),
		db.Einkauf.Pfand.SetIfPresent(einkauf.Pfand),
		db.Einkauf.Mitarbeiter.Link(
			db.Mitarbeiter.ID.Equals(mitarbeiterId),
		),
	).Update(
		db.Einkauf.Dinge.Set(einkauf.Dinge),
		db.Einkauf.Abgeschickt.Set(time.Now()),
		db.Einkauf.Abonniert.Set(Abo),
		db.Einkauf.Paypal.Set(Paypal),
		db.Einkauf.Geld.SetIfPresent(einkauf.Geld),
		db.Einkauf.Pfand.SetIfPresent(einkauf.Pfand),
	).Exec(ctx)
	if err != nil {
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
