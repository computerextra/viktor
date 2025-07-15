package handler

import (
	"net/http"

	"github.com/computerextra/viktor/db"
	"github.com/computerextra/viktor/frontend"
)

type LieferantProps struct {
	Name string `schema:"name,required"`
}

func (h *Handler) GetLieferanten(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	res, err := h.db.Lieferant.FindMany().With(
		db.Lieferant.Ansprechpartner.Fetch(),
	).OrderBy(db.Lieferant.Firma.Order(db.SortOrderAsc)).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}

	uri := getPath(r.URL.Path)
	frontend.LieferantenOverview(res, uri).Render(ctx, w)
}

func (h *Handler) NewLieferant(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement
}

func (h *Handler) CreateLieferant(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement
}

func (h *Handler) GetLieferant(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement
}

func (h *Handler) UpdateLieferant(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement
}

func (h *Handler) DeleteLieferant(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement
}
