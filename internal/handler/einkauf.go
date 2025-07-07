package handler

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/computerextra/viktor/db"
	"github.com/computerextra/viktor/internal/util/flash"
)

type EinkaufProps struct {
	Abo    bool    `schema:"Abo,default:false"`
	Paypal bool    `schema:"Paypal,default:false"`
	Dinge  string  `schema:"Dinge,required"`
	Geld   *string `schema:"Geld"`
	Pfand  *string `schema:"Pfand"`
}

func (h *Handler) GetImage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	mitarbeiter, err := h.db.Mitarbeiter.FindUnique(db.Mitarbeiter.ID.Equals(id)).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	einkaufID, ok := mitarbeiter.EinkaufID()
	if !ok {
		sendJsonData(marshalData("", w, h.logger), w)
	}
	images, err := h.db.Einkauf.FindFirst(db.Einkauf.ID.Equals(einkaufID)).Select(
		db.Einkauf.Bild1.Field(),
		db.Einkauf.Bild2.Field(),
		db.Einkauf.Bild3.Field(),
	).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	sendJsonData(marshalData(images, w, h.logger), w)
}

func (h *Handler) UploadImage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	r.ParseForm()
	imageNr := r.FormValue("imagenr")
	imageName := r.FormValue("imagename")
	url, ok := os.LookupEnv("UPLOADTHING_URL")
	if !ok {
		h.logger.Error("failed to parse env", slog.String("key", "UPLOADTHING_URL"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	path := fmt.Sprintf("%s%s", url, imageName)
	mitarbeiter, err := h.db.Mitarbeiter.FindUnique(db.Mitarbeiter.ID.Equals(id)).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	einkaufID, ok := mitarbeiter.EinkaufID()
	if !ok {
		einkaufNew, err := h.db.Einkauf.CreateOne(
			db.Einkauf.Dinge.Set(""),
			db.Einkauf.Abgeschickt.Set(time.Now().AddDate(1, 0, 0)),
		).Exec(ctx)
		if err != nil {
			sendQueryError(w, h.logger, err)
		}
		einkaufID = einkaufNew.ID
	}
	switch imageNr {
	case "1":
		res, err := h.db.Einkauf.FindUnique(db.Einkauf.ID.Equals(einkaufID)).Update(db.Einkauf.Bild1.Set(path)).Exec(ctx)
		if err != nil {
			sendQueryError(w, h.logger, err)
		}
		sendJsonData(marshalData(res, w, h.logger), w)
	case "2":
		res, err := h.db.Einkauf.FindUnique(db.Einkauf.ID.Equals(einkaufID)).Update(db.Einkauf.Bild2.Set(path)).Exec(ctx)
		if err != nil {
			sendQueryError(w, h.logger, err)
		}
		sendJsonData(marshalData(res, w, h.logger), w)
	case "3":
		res, err := h.db.Einkauf.FindUnique(db.Einkauf.ID.Equals(einkaufID)).Update(db.Einkauf.Bild3.Set(path)).Exec(ctx)
		if err != nil {
			sendQueryError(w, h.logger, err)
		}
		sendJsonData(marshalData(res, w, h.logger), w)

	}
}

func (h *Handler) DeleteImage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	r.ParseForm()
	imageNr := r.FormValue("imagenr")
	mitarbeiter, err := h.db.Mitarbeiter.FindUnique(db.Mitarbeiter.ID.Equals(id)).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	einkaufID, ok := mitarbeiter.EinkaufID()
	if !ok {
		einkaufNew, err := h.db.Einkauf.CreateOne(
			db.Einkauf.Dinge.Set(""),
			db.Einkauf.Abgeschickt.Set(time.Now().AddDate(-1, 0, 0)),
		).Exec(ctx)
		if err != nil {
			sendQueryError(w, h.logger, err)
		}
		einkaufID = einkaufNew.ID
	}
	var null *string
	switch imageNr {
	case "1":
		res, err := h.db.Einkauf.FindUnique(db.Einkauf.ID.Equals(einkaufID)).Update(db.Einkauf.Bild1.SetIfPresent(null)).Exec(ctx)
		if err != nil {
			sendQueryError(w, h.logger, err)
		}
		sendJsonData(marshalData(res, w, h.logger), w)
	case "2":
		res, err := h.db.Einkauf.FindUnique(db.Einkauf.ID.Equals(einkaufID)).Update(db.Einkauf.Bild2.SetIfPresent(null)).Exec(ctx)
		if err != nil {
			sendQueryError(w, h.logger, err)
		}
		sendJsonData(marshalData(res, w, h.logger), w)
	case "3":
		res, err := h.db.Einkauf.FindUnique(db.Einkauf.ID.Equals(einkaufID)).Update(db.Einkauf.Bild3.SetIfPresent(null)).Exec(ctx)
		if err != nil {
			sendQueryError(w, h.logger, err)
		}
		sendJsonData(marshalData(res, w, h.logger), w)

	}
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
		sendQueryError(w, h.logger, err)
	}
	sendJsonData(marshalData(mitarbeiter, w, h.logger), w)
}

func (h *Handler) GetListe(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	res, err := h.db.Einkauf.FindMany(db.Einkauf.Or(
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
		sendQueryError(w, h.logger, err)
	}
	sendJsonData(marshalData(res, w, h.logger), w)
}

func (h *Handler) SkipEinkauf(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	res, err := h.db.Einkauf.FindUnique(db.Einkauf.ID.Equals(id)).Update(
		db.Einkauf.Abgeschickt.Set(time.Now().AddDate(0, 0, 1)),
	).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	sendJsonData(marshalData(res, w, h.logger), w)
}

func (h *Handler) DeleteEinkauf(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	if id == "" {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	res, err := h.db.Einkauf.FindUnique(db.Einkauf.ID.Equals(id)).Update(
		db.Einkauf.Abgeschickt.Set(time.Now().AddDate(-1, 0, 0)),
		db.Einkauf.Abonniert.Set(false),
	).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	sendJsonData(marshalData(res, w, h.logger), w)
}

func (h *Handler) UpdateEinkauf(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	mitarbeiterId := r.PathValue("id")
	r.ParseForm()
	var einkauf EinkaufProps
	err := decoder.Decode(&einkauf, r.PostForm)
	if err != nil {
		flash.SetFlashMessage(w, "error", "content cannot be empty")
		h.logger.Error("failed to parse formdata", slog.Any("error", err))
		w.WriteHeader(http.StatusNoContent)
		return
	}
	mitarbeiter, err := h.db.Mitarbeiter.FindUnique(db.Mitarbeiter.ID.Equals(mitarbeiterId)).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	einkauId, _ := mitarbeiter.EinkaufID()
	res, err := h.db.Einkauf.UpsertOne(
		db.Einkauf.ID.EqualsIfPresent(&einkauId),
	).Create(
		db.Einkauf.Dinge.Set(einkauf.Dinge),
		db.Einkauf.Abgeschickt.Set(time.Now()),
		db.Einkauf.Abonniert.Set(einkauf.Abo),
		db.Einkauf.Paypal.Set(einkauf.Paypal),
		db.Einkauf.Geld.SetIfPresent(einkauf.Geld),
		db.Einkauf.Pfand.SetIfPresent(einkauf.Pfand),
	).Update(
		db.Einkauf.Dinge.Set(einkauf.Dinge),
		db.Einkauf.Abgeschickt.Set(time.Now()),
		db.Einkauf.Abonniert.Set(einkauf.Abo),
		db.Einkauf.Paypal.Set(einkauf.Paypal),
		db.Einkauf.Geld.SetIfPresent(einkauf.Geld),
		db.Einkauf.Pfand.SetIfPresent(einkauf.Pfand),
	).Exec(ctx)
	if err != nil {
		sendQueryError(w, h.logger, err)
	}
	sendJsonData(marshalData(res, w, h.logger), w)
}
