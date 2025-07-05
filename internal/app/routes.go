package app

import (
	"fmt"
	"io/fs"
	"net/http"

	"github.com/computerextra/viktor/internal/handler"
)

func (a *App) loadRoutes() (http.Handler, error) {
	router := http.NewServeMux()

	reactApp, err := fs.Sub(a.frontend, "frontend/dist")
	if err != nil {
		return nil, fmt.Errorf("error finding dist folder: %w", err)
	}

	router.Handle("GET /", http.FileServerFS(reactApp))

	a.loadPages(router)

	return router, nil
}

func (a *App) loadPages(router *http.ServeMux) {
	h := handler.New(a.logger)

	router.HandleFunc("GET /api/something", h.GetSomething)
	router.HandleFunc("POST /api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}
