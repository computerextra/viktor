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

type AngeboteProps struct {
	Title      string  `schema:"title,required"`
	SubTitle   *string `schema:"subtitle"`
	Image      string  `schema:"image,required"`
	Link       string  `schema:"link,required"`
	Date_start string  `schema:"date_start,required"`
	Date_stop  string  `schema:"date_stop,required"`
	Anzeigen   bool    `schema:"anzeigen,default:false"`
}

func (h *Handler) NewAngebot(w http.ResponseWriter, r *http.Request) {
	uri := getPath(r.URL.Path)

	frontend.NeuesAngebot(uri).Render(r.Context(), w)
}
func (h *Handler) GetAngebote(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	res, err := h.db.Angebot.FindMany().OrderBy(
		db.Angebot.Title.Order(db.SortOrderAsc),
	).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	uri := getPath(r.URL.Path)
	frontend.AngeboteOverview(res, uri).Render(ctx, w)
}

func (h *Handler) GetAngebot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	res, err := h.db.Angebot.FindUnique(db.Angebot.ID.Equals(id)).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	uri := getPath(r.URL.Path)
	frontend.AngebotBearbeiten(*res, uri).Render(ctx, w)
}

func (h *Handler) CreateAngebot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	r.ParseMultipartForm(10 << 20) // Max Header size (e.g. 10MB)
	var props AngeboteProps
	err := decoder.Decode(&props, r.PostForm)
	if err != nil {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		h.logger.Error("failed to parse formdata", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}

	Date_Start, err := time.Parse("2006-01-02", props.Date_start)
	if err != nil {
		sendError(w, h.logger, "failed to parse date", err)
	}
	Date_Stop, err := time.Parse("2006-01-02", props.Date_stop)
	if err != nil {
		sendError(w, h.logger, "failed to parse date", err)
	}

	_, err = h.db.Angebot.CreateOne(
		db.Angebot.Title.Set(props.Title),
		db.Angebot.DateStart.Set(Date_Start),
		db.Angebot.DateStop.Set(Date_Stop),
		db.Angebot.Link.Set(props.Link),
		db.Angebot.Image.Set(props.Image),
		db.Angebot.Subtitle.SetIfPresent(props.SubTitle),
		db.Angebot.Anzeigen.Set(props.Anzeigen),
	).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	host := r.Host
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	uri := fmt.Sprintf("%s://%s/CMS/Angebote", scheme, host)
	http.Redirect(w, r, uri, http.StatusFound)
}

func (h *Handler) ToggleAngebot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")

	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	status, err := h.db.Angebot.FindUnique(db.Angebot.ID.Equals(id)).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}

	_, err = h.db.Angebot.FindUnique(db.Angebot.ID.Equals(id)).Update(
		db.Angebot.Anzeigen.Set(!status.Anzeigen),
	).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	host := r.Host
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	uri := fmt.Sprintf("%s://%s/CMS/Angebote", scheme, host)
	http.Redirect(w, r, uri, http.StatusFound)
}

func (h *Handler) UpdateAngebot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	r.ParseMultipartForm(10 << 20) // Max Header size (e.g. 10MB)
	id := r.PathValue("id")
	var props AngeboteProps
	err := decoder.Decode(&props, r.PostForm)
	if err != nil {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		h.logger.Error("failed to parse formdata", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}

	Date_Start, err := time.Parse("2006-01-02", props.Date_start)
	if err != nil {
		sendError(w, h.logger, "failed to parse date", err)
	}
	Date_Stop, err := time.Parse("2006-01-02", props.Date_stop)
	if err != nil {
		sendError(w, h.logger, "failed to parse date", err)
	}

	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	_, err = h.db.Angebot.FindUnique(db.Angebot.ID.Equals(id)).Update(
		db.Angebot.Title.Set(props.Title),
		db.Angebot.DateStart.Set(Date_Start),
		db.Angebot.DateStop.Set(Date_Stop),
		db.Angebot.Link.Set(props.Link),
		db.Angebot.Image.Set(props.Image),
		db.Angebot.Subtitle.SetIfPresent(props.SubTitle),
		db.Angebot.Anzeigen.Set(props.Anzeigen),
	).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	host := r.Host
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	uri := fmt.Sprintf("%s://%s/CMS/Angebote", scheme, host)
	http.Redirect(w, r, uri, http.StatusFound)
}

func (h *Handler) DeleteAngebot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	_, err := h.db.Angebot.FindUnique(db.Angebot.ID.Equals(id)).Delete().Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	host := r.Host
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	uri := fmt.Sprintf("%s://%s/CMS/Angebote", scheme, host)
	http.Redirect(w, r, uri, http.StatusFound)
}
