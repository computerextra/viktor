package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/computerextra/viktor/frontend"
	"github.com/computerextra/viktor/internal"
	"github.com/computerextra/viktor/internal/util/flash"
)

func (h *Handler) GetMandate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	resp, err := http.Get("https://api.computer-extra.com")
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}
	defer resp.Body.Close()

	jsonStr, err := io.ReadAll(resp.Body)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}

	var data []internal.Sepa

	err = json.Unmarshal(jsonStr, &data)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}

	uri := getPath(r.URL.Path)
	frontend.Sepa(data, uri).Render(ctx, w)
}

type MandatProps struct {
	Kundennummer string  `schema:"kundennummer,required"`
	Firma        *string `schema:"firma"`
	Name         string  `schema:"name,required"`
	Email        string  `schema:"email,required"`
}

func (h *Handler) SetOfflineMandat(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	var props MandatProps
	err := decoder.Decode(&props, r.PostForm)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		h.logger.Error("failed to parse formdata", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}

	jsonData := []byte(fmt.Sprintf(`{"kundennummer": "%s", "firma": "%s", "name": "%s", "mail": "%s"}`, props.Kundennummer, *props.Firma, props.Name))

	// Post zur API
	_, err = http.Post("https://api.computer-extra.com", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		h.logger.Error("failed to post formdata", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}

	host := r.Host
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	uri := fmt.Sprintf("%s://%s/Sepa", scheme, host)
	http.Redirect(w, r, uri, http.StatusFound)
}
