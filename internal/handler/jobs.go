package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/computerextra/viktor/db"
	"github.com/computerextra/viktor/frontend"
	"github.com/computerextra/viktor/internal/util/flash"
)

type JobProps struct {
	Name   string `schema:"name,required"`
	Online bool   `schema:"online,default:false"`
}

func (h *Handler) NewJob(w http.ResponseWriter, r *http.Request) {
	uri := getPath(r.URL.Path)

	frontend.NeuerJob(uri).Render(r.Context(), w)
}

func (h *Handler) GetJobs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	res, err := h.db.Jobs.FindMany().OrderBy(db.Jobs.Name.Order(db.SortOrderAsc)).Exec(ctx)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}

	uri := getPath(r.URL.Path)
	frontend.JobsOverview(res, uri).Render(ctx, w)
}

func (h *Handler) GetJob(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	res, err := h.db.Jobs.FindUnique(db.Jobs.ID.Equals(id)).Exec(ctx)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}
	uri := getPath(r.URL.Path)
	frontend.JobBearbeiten(res, uri).Render(ctx, w)
}

func (h *Handler) CreateJob(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	r.ParseMultipartForm(10 << 20) // Max Header size (e.g. 10MB)
	var props JobProps
	err := decoder.Decode(&props, r.PostForm)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		h.logger.Error("failed to parse formdata", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}
	_, err = h.db.Jobs.CreateOne(db.Jobs.Name.Set(props.Name), db.Jobs.Online.Set(props.Online)).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	host := r.Host
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	uri := fmt.Sprintf("%s://%s/CMS/Jobs", scheme, host)
	http.Redirect(w, r, uri, http.StatusFound)
}

func (h *Handler) ToggleJob(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")

	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	status, err := h.db.Jobs.FindUnique(db.Jobs.ID.Equals(id)).Exec(ctx)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}

	_, err = h.db.Jobs.FindUnique(db.Jobs.ID.Equals(id)).Update(
		db.Jobs.Online.Set(!status.Online),
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
	uri := fmt.Sprintf("%s://%s/CMS/Jobs", scheme, host)
	http.Redirect(w, r, uri, http.StatusFound)
}

func (h *Handler) UpdateJob(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	r.ParseMultipartForm(10 << 20) // Max Header size (e.g. 10MB)
	id := r.PathValue("id")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	var props JobProps
	err := decoder.Decode(&props, r.PostForm)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		h.logger.Error("failed to parse formdata", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}
	_, err = h.db.Jobs.FindUnique(db.Jobs.ID.Equals(id)).Update(
		db.Jobs.Name.Set(props.Name),
		db.Jobs.Online.Set(props.Online),
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
	uri := fmt.Sprintf("%s://%s/CMS/Jobs", scheme, host)
	http.Redirect(w, r, uri, http.StatusFound)
}

func (h *Handler) DeleteJob(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	_, err := h.db.Jobs.FindUnique(db.Jobs.ID.Equals(id)).Delete().Exec(ctx)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}
	host := r.Host
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	uri := fmt.Sprintf("%s://%s/CMS/Jobs", scheme, host)
	http.Redirect(w, r, uri, http.StatusFound)
}
