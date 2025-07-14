package handler

import (
	"net/http"

	"github.com/computerextra/viktor/frontend"
)

func (h *Handler) GetStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	uri := getPath(r.URL.Path)

	frontend.Status(uri).Render(ctx, w)
}
