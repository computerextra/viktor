package handler

import (
	"encoding/json"
	"io"
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
