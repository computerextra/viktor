package handler

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/computerextra/viktor/db"
	"github.com/computerextra/viktor/frontend"
	"github.com/computerextra/viktor/internal/util/flash"
	_ "github.com/denisenkom/go-mssqldb"

	gomail "gopkg.in/mail.v2"
)

type Warenlieferung struct {
	Name          string
	Angelegt      time.Time
	Geliefert     sql.NullTime
	Alterpreis    sql.NullFloat64
	Neuerpreis    sql.NullFloat64
	Preis         sql.NullTime
	Artikelnummer string
	ID            int32
}
type Artikel struct {
	Id            int
	Artikelnummer string
	Suchbegriff   string
	Preis         float64
}

type Leichen struct {
	Artikelnummer string
	Artikelname   string
	Bestand       int16
	Verfügbar     int16
	EK            float64
	LetzterUmsatz string
}

type SummenArtikel struct {
	Artikelnummer string
	Artikelname   string
	Bestand       int16
	EK            float64
	Summe         float64
}

type VerfArtikel struct {
	SummenArtikel
	Verfügbar int16
}

type WertArtikel struct {
	Bestand   int16
	Verfügbar int16
	EK        float64
}

type History struct {
	Id     int
	Action string
}

type Price struct {
	Id     int
	Action string
	Price  float64
}

type AlteSeriennummer struct {
	ArtNr       string
	Suchbegriff string
	Bestand     int
	Verfügbar   int
	GeBeginn    string
}

type AccessArtikel struct {
	Id            int
	Artikelnummer string
	Artikeltext   string
	Preis         float64
}
type AusstellerArtikel struct {
	Id            int
	Artikelnummer string
	Artikelname   string
	Specs         string
	Preis         float64
}
type SageArtikel struct {
	Id            int
	Artikelnummer string
	Suchbegriff   string
	Preis         float64
}
type sg_auf_fschrift struct {
	ERFART    sql.NullString
	USERNAME  sql.NullString
	DATUM     sql.NullTime
	AUFNR     sql.NullString
	ALTAUFNR  sql.NullString
	ENDPRB    sql.NullFloat64
	VERTRETER sql.NullString
	NAME      sql.NullString
}

type Response struct {
	Auftrag    string
	Wert       float64
	Verbrecher string
	Datum      time.Time
	Kunde      string
}

func getSageConnectionString() string {
	sageurl, ok := os.LookupEnv("SAGE_URL")
	if !ok {
		return "nil"
	}
	return sageurl
}

func (h *Handler) Warenlieferung(w http.ResponseWriter, r *http.Request) {
	uri := getPath(r.URL.Path)
	frontend.Warenlieferung(uri, false, false).Render(r.Context(), w)
}

func (h *Handler) GenerateWarenlieferung(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	AlleArtikel, err := h.db.Warenlieferung.FindMany().Exec(ctx)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}

	neueArtikel, geliefert, neuePreise, err := sortProducts(AlleArtikel)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		h.logger.Error("failed to connect to sage", slog.Any("error", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, item := range neueArtikel {
		_, err := h.db.Warenlieferung.CreateOne(
			db.Warenlieferung.Name.Set(item.Name),
			db.Warenlieferung.Artikelnummer.Set(item.Artikelnummer),
			db.Warenlieferung.ID.Set(int(item.ID)),
		).Exec(ctx)

		if err != nil {
			flash.SetFlashMessage(w, "error", err.Error())
			sendQueryError(w, h.logger, err)
		}
	}
	for _, item := range geliefert {
		_, err := h.db.Warenlieferung.FindUnique(db.Warenlieferung.ID.Equals(int(item.ID))).Update(
			db.Warenlieferung.Name.Set(item.Name),
			db.Warenlieferung.Geliefert.Set(time.Now()),
		).Exec(ctx)

		if err != nil {
			flash.SetFlashMessage(w, "error", err.Error())
			sendQueryError(w, h.logger, err)
		}
	}
	for _, item := range neuePreise {
		var altFloat *float64
		var neuFloat *float64
		if item.Alterpreis.Valid {
			altFloat = &item.Alterpreis.Float64
		}
		if item.Neuerpreis.Valid {
			neuFloat = &item.Neuerpreis.Float64
		}
		if (neuFloat != nil && *neuFloat > 0) || (altFloat != nil && *altFloat > 0) {
			_, err := h.db.Warenlieferung.FindUnique(
				db.Warenlieferung.ID.Equals(int(item.ID)),
			).Update(
				db.Warenlieferung.AlterPreis.SetIfPresent(altFloat),
				db.Warenlieferung.NeuerPreis.SetIfPresent(neuFloat),
				db.Warenlieferung.Preis.Set(time.Now()),
			).Exec(ctx)
			if err != nil {
				flash.SetFlashMessage(w, "error", err.Error())
				sendQueryError(w, h.logger, err)
			}
		}
	}
	uri := getPath(r.URL.Path)
	frontend.Warenlieferung(uri, true, false).Render(r.Context(), w)
}

// TODO: Preisberechnung in Prozent geht nicht!
// •	1101375 - VER KYOCERA TK-5270K, schwarz, ~8000 Seiten: 99.90 ➡️ 124.90 (0.00 % // 25.00 €)
func (h *Handler) SendWarenlieferung(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	Mitarbeiter, err := h.db.Mitarbeiter.FindMany(
		db.Mitarbeiter.Not(
			db.Mitarbeiter.Mail.IsNull(),
		),
	).Exec(ctx)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}
	var NeueArtikel []db.WarenlieferungModel
	err = h.db.Prisma.QueryRaw("SELECT * FROM Warenlieferung WHERE DATE_FORMAT(angelegt, '%Y-%m-%d') = DATE_FORMAT(NOW(), '%Y-%m-%d') ORDER BY Artikelnummer ASC").Exec(ctx, &NeueArtikel)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}
	var GelieferteArtikel []db.WarenlieferungModel
	err = h.db.Prisma.QueryRaw("SELECT * FROM Warenlieferung WHERE DATE_FORMAT(geliefert, '%Y-%m-%d') = DATE_FORMAT(NOW(), '%Y-%m-%d') AND DATE_FORMAT(angelegt, '%Y-%m-%d') != DATE_FORMAT(NOW(), '%Y-%m-%d') ORDER BY Artikelnummer ASC").Exec(ctx, &GelieferteArtikel)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}
	var NeuePreise []db.WarenlieferungModel
	err = h.db.Prisma.QueryRaw("SELECT * FROM Warenlieferung WHERE DATE_FORMAT(Preis, '%Y-%m-%d') = DATE_FORMAT(NOW(), '%Y-%m-%d') AND DATE_FORMAT(angelegt, '%Y-%m-%d') != DATE_FORMAT(NOW(), '%Y-%m-%d') ORDER BY Artikelnummer ASC").Exec(ctx, &NeuePreise)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}
	wertBestand, wertVerfügbar, err := getLagerWert()
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}
	teureArtikel, err := getHighestSum()
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}
	teureVerfArtikel, err := getHighestVerfSum()
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}
	leichen, err := getLeichen()
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}
	SN, err := getAlteSeriennummern()
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}
	Verbrecher, gesamtWert, err := getOldAuftraege()
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		sendQueryError(w, h.logger, err)
	}

	// Mail Versand
	from, ok := os.LookupEnv("SMTP_FROM")
	if !ok {
		flash.SetFlashMessage(w, "error", "failed to parse env")
		h.logger.Error("failed to parse env", slog.String("key", "SMTP_FROM"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	server, ok := os.LookupEnv("SMTP_HOST")
	if !ok {
		flash.SetFlashMessage(w, "error", "failed to parse env")
		h.logger.Error("failed to parse env", slog.String("key", "SMTP_HOST"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user, ok := os.LookupEnv("SMTP_USER")
	if !ok {
		flash.SetFlashMessage(w, "error", "failed to parse env")
		h.logger.Error("failed to parse env", slog.String("key", "SMTP_USER"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	pass, ok := os.LookupEnv("SMTP_PASS")
	if !ok {
		flash.SetFlashMessage(w, "error", "failed to parse env")
		h.logger.Error("failed to parse env", slog.String("key", "SMTP_PASS"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	portStr, ok := os.LookupEnv("SMTP_PORT")
	if !ok {
		flash.SetFlashMessage(w, "error", "failed to parse env")
		h.logger.Error("failed to parse env", slog.String("key", "SMTP_PORT"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		h.logger.Error("failed to parse env", slog.Any("error", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var body string
	if len(NeueArtikel) > 0 {
		body = fmt.Sprintf("%s<h2>Neue Artikel</h2><ul>", body)

		for i := range NeueArtikel {
			body = fmt.Sprintf("%s<li><b>%s</b> - %s</li>", body, NeueArtikel[i].Artikelnummer, NeueArtikel[i].Name)
		}
		body = fmt.Sprintf("%s</ul>", body)
	}

	if len(GelieferteArtikel) > 0 {
		body = fmt.Sprintf("%s<br><br><h2>Gelieferte Artikel</h2><ul>", body)

		for i := range GelieferteArtikel {
			body = fmt.Sprintf("%s<li><b>%s</b> - %s</li>", body, GelieferteArtikel[i].Artikelnummer, GelieferteArtikel[i].Name)
		}
		body = fmt.Sprintf("%s</ul>", body)
	}

	if len(NeuePreise) > 0 {
		body = fmt.Sprintf("%s<br><br><h2>Preisänderungen</h2><ul>", body)

		for i := range NeuePreise {
			var altFloat float64
			var neuFloat float64
			altFloat, ok_alt := NeuePreise[i].AlterPreis()
			neuFloat, ok_neu := NeuePreise[i].NeuerPreis()

			if ok_alt && ok_neu && neuFloat != altFloat {
				body = fmt.Sprintf("%s<li><b>%s</b> - %s: %.2f ➡️ %.2f ", body, NeuePreise[i].Artikelnummer, NeuePreise[i].Name, altFloat, neuFloat)

				absolute := neuFloat - altFloat
				prozent := ((altFloat / altFloat) * 100) - 100
				body = fmt.Sprintf("%s(%.2f %% // %.2f €)</li>", body, prozent, absolute)
			}

		}
		body = fmt.Sprintf("%s</ul>", body)
	}

	body = fmt.Sprintf("%s<h2>Aktuelle Lagerwerte</h2><p><b>Lagerwert Verfügbare Artikel:</b> %.2f €</p><p><b>Lagerwert alle lagernde Artikel:</b> %.2f €</p>", body, wertVerfügbar, wertBestand)
	body = fmt.Sprintf("%s<p>Wert in aktuellen Aufträgen: %.2f €", body, wertBestand-wertVerfügbar)
	body = fmt.Sprintf("%s<p>Offene Posten laut Sage: %.2f €* (Hier kann nicht nach bereits lagernder Ware gesucht werden!)", body, gesamtWert)

	if len(SN) > 0 {
		body = fmt.Sprintf("%s<h2>Artikel mit alten Seriennummern</h2><p>Nachfolgende Artikel sollten mit erhöhter Prioriät verkauf werden, da die Seriennummern bereits sehr alt sind. Gegebenenfalls sind die Artikel bereits außerhalb der Herstellergarantie!</p>", body)
		body = fmt.Sprintf("%s<p>Folgende Werte gelten:</p>", body)
		body = fmt.Sprintf("%s<p>Wortmann: Angebene Garantielaufzeit + 2 Monate ab Kaufdatum CompEx</p>", body)
		body = fmt.Sprintf("%s<p>Lenovo: Angegebene Garantielaufzeit ab Kauf CompEx</p>", body)
		body = fmt.Sprintf("%s<p>Bei allen anderen Herstellern gilt teilweise das Kaufdatum des Kunden. <br>Falls sich dies ändern sollte, wird es in der Aufzählung ergänzt.</p>", body)

		body = fmt.Sprintf("%s<p>Erklärungen der Farben:</p>", body)
		body = fmt.Sprintf("%s<p><span style='background-color: \"#f45865\"'>ROT:</span> Artikel ist bereits seit mehr als 2 Jahren lagernd und sollte schnellstens Verkauft werden!</p>", body)
		body = fmt.Sprintf("%s<p><span style='background-color: \"#fff200\"'>Gelb:</span> Artikel ist bereits seit mehr als 1 Jahr lagernd!</p>", body)

		body = fmt.Sprintf("%s<table><thead>", body)
		body = fmt.Sprintf("%s<tr>", body)
		body = fmt.Sprintf("%s<th>Artikelnummer</th>", body)
		body = fmt.Sprintf("%s<th>Name</th>", body)
		body = fmt.Sprintf("%s<th>Bestand</th>", body)
		body = fmt.Sprintf("%s<th>Verfügbar</th>", body)
		body = fmt.Sprintf("%s<th>Garantiebeginn des ältesten Artikels</th>", body)
		body = fmt.Sprintf("%s</tr>", body)
		body = fmt.Sprintf("%s</thead>", body)
		body = fmt.Sprintf("%s</thbody>", body)
		for i := range SN {
			year, _, _ := time.Now().Date()
			tmp := strings.Split(strings.Replace(strings.Split(SN[i].GeBeginn, "T")[0], "-", ".", -1), ".")
			year_tmp, err := strconv.Atoi(tmp[0])
			if err != nil {
				flash.SetFlashMessage(w, "error", err.Error())
				h.logger.Error("SendMail: Fehler beim voncertieren von string zu int (year) in GetAlteSeriennummern!", slog.Any("error", err))
			}

			GarantieBeginn := fmt.Sprintf("%s.%s.%s", tmp[2], tmp[1], tmp[0])
			diff := year - year_tmp
			if diff >= 2 {
				body = fmt.Sprintf("%s<tr style='background-color: \"#f45865\"'>", body)
			} else if diff >= 1 {
				body = fmt.Sprintf("%s<tr style='background-color: \"#fff200\"'>", body)
			} else {
				body = fmt.Sprintf("%s<tr>", body)
			}
			body = fmt.Sprintf("%s<td>%s</td>", body, SN[i].ArtNr)
			body = fmt.Sprintf("%s<td>%s</td>", body, SN[i].Suchbegriff)
			body = fmt.Sprintf("%s<td>%v</td>", body, SN[i].Bestand)
			body = fmt.Sprintf("%s<td>%v</td>", body, SN[i].Verfügbar)
			body = fmt.Sprintf("%s<td>%s</td>", body, GarantieBeginn)
			body = fmt.Sprintf("%s</tr>", body)

		}
		body = fmt.Sprintf("%s</tbody></table>", body)
	}

	if len(teureArtikel) > 0 {
		body = fmt.Sprintf("%s<h2>Top 10: Die teuersten Artikel inkl. aktive Aufträge</h2><table><thead><tr><th>Artikelnummer</th><th>Name</th><th>Bestand</th><th>Einzelpreis</th><th>Summe</th></tr></thead><tbody>", body)

		for i := range teureArtikel {
			body = fmt.Sprintf("%s<tr><td>%s</td><td>%s</td><td>%d</td><td>%.2f €</td><td>%.2f €</td></tr>", body, teureArtikel[i].Artikelnummer, teureArtikel[i].Artikelname, teureArtikel[i].Bestand, teureArtikel[i].EK, teureArtikel[i].Summe)
		}
		body = fmt.Sprintf("%s</tbody></table>", body)
	}

	if len(teureVerfArtikel) > 0 {
		body = fmt.Sprintf("%s<h2>Top 10: Die teuersten Artikel exkl. aktive Aufträge</h2><table><thead><tr><th>Artikelnummer</th><th>Name</th><th>Bestand</th><th>Einzelpreis</th><th>Summe</th></tr></thead><tbody>", body)

		for i := range teureVerfArtikel {
			body = fmt.Sprintf("%s<tr><td>%s</td><td>%s</td><td>%d</td><td>%.2f €</td><td>%.2f €</td></tr>", body, teureVerfArtikel[i].Artikelnummer, teureVerfArtikel[i].Artikelname, teureVerfArtikel[i].Bestand, teureVerfArtikel[i].EK, teureVerfArtikel[i].Summe)

		}
		body = fmt.Sprintf("%s</tbody></table>", body)
	}

	if len(leichen) > 0 {
		body = fmt.Sprintf("%s<h2>Top 20: Leichen bei CE</h2><table><thead><tr><th>Artikelnummer</th><th>Name</th><th>Bestand</th><th>Verfügbar</th><th>Letzter Umsatz:</th><th>Wert im Lager:</th></tr></thead><tbody>", body)
		for i := range leichen {
			summe := float64(leichen[i].Verfügbar) * leichen[i].EK
			var LetzterUmsatz string
			if leichen[i].LetzterUmsatz == "1899-12-30T00:00:00Z" {
				LetzterUmsatz = "nie"
			} else {
				tmp := strings.Split(strings.Replace(strings.Split(leichen[i].LetzterUmsatz, "T")[0], "-", ".", -1), ".")
				LetzterUmsatz = fmt.Sprintf("%s.%s.%s", tmp[2], tmp[1], tmp[0])
			}
			bestand := leichen[i].Bestand
			verf := leichen[i].Verfügbar
			artNr := leichen[i].Artikelnummer
			name := leichen[i].Artikelname
			body = fmt.Sprintf("%s<tr><td>%s</td><td>%s</td><td>%d</td><td>%d</td><td>%s</td><td>%.2f€</td></tr>", body, artNr, name, bestand, verf, LetzterUmsatz, summe)
		}
		body = fmt.Sprintf("%s</tbody></table>", body)
	}

	if len(Verbrecher) > 0 {
		body = fmt.Sprintf("%s<h2>Aktuell offene Aufträge & Lieferscheine (Es können Leichen dabei sein, das lässt sich leider nicht korrekt filtern)</h2><table><thead><tr><th>Auftrag</th><th>Summe Brutto</th><th>Vertreter</th><th>Kundenname</th><th>Datum</th></tr></thead><tbody>", body)

		for _, i := range Verbrecher {
			body = fmt.Sprintf("%s<tr>", body)
			body = fmt.Sprintf("%s<td>%s</td>", body, i.Auftrag)
			body = fmt.Sprintf("%s<td>%.2f €</td>", body, i.Wert)
			body = fmt.Sprintf("%s<td>%s</td>", body, i.Verbrecher)
			body = fmt.Sprintf("%s<td>%s</td>", body, i.Kunde)
			date := strings.Split(i.Datum.Format(time.DateOnly), "-")
			body = fmt.Sprintf("%s<td>%s.%s.%s</td>", body, date[2], date[1], date[0])
			body = fmt.Sprintf("%s</tr>", body)

		}
		body = fmt.Sprintf("%s</tbody></table>", body)
	}

	d := gomail.NewDialer(server, port, user, pass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	s, err := d.Dial()
	if err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		h.logger.Error("failed to dial smtp server", slog.Any("error", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	m := gomail.NewMessage()

	// TODO: Testen nach Warenlieferung!
	m.SetHeader("From", from)
	m.SetHeader("To", "johannes.kirchner@computer-extra.de")
	m.SetHeader("Subject", fmt.Sprintf("Warenlieferung vom %v", time.Now().Format(time.DateOnly)))
	m.SetBody("text/html", body)
	if err := gomail.Send(s, m); err != nil {
		flash.SetFlashMessage(w, "error", err.Error())
		h.logger.Error("failed to send mail", slog.Any("error", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	uri := getPath(r.URL.Path)
	frontend.Warenlieferung(uri, true, true).Render(r.Context(), w)
	return

	for _, ma := range Mitarbeiter {
		mail, ok := ma.Mail()
		if ok && len(mail) > 1 {
			m.SetHeader("From", from)
			m.SetHeader("To", mail)
			m.SetHeader("Subject", fmt.Sprintf("Warenlieferung vom %v", time.Now().Format(time.DateOnly)))
			m.SetBody("text/html", body)
			if err := gomail.Send(s, m); err != nil {
				flash.SetFlashMessage(w, "error", err.Error())
				h.logger.Error("failed to send mail", slog.Any("error", err))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			m.Reset()
		}
	}
	// uri := getPath(r.URL.Path)
	// frontend.Warenlieferung(uri, true, true).Render(r.Context(), w)
}

func sortProducts(Products []db.WarenlieferungModel) ([]Warenlieferung, []Warenlieferung, []Warenlieferung, error) {
	Sage, err := getAllProductsFromSage()
	if err != nil {

		return nil, nil, nil, err
	}
	History, err := getLagerHistory()
	if err != nil {
		return nil, nil, nil, err
	}
	Prices, err := getPrices()
	if err != nil {
		return nil, nil, nil, err
	}

	var neueArtikel []Warenlieferung
	var gelieferteArtikel []Warenlieferung
	var geliefert []int
	var neuePreise []Warenlieferung

	if len(Products) <= 0 {
		for i := range Sage {
			neu := Warenlieferung{
				ID:            int32(Sage[i].Id),
				Artikelnummer: Sage[i].Artikelnummer,
				Name:          Sage[i].Suchbegriff,
			}
			neueArtikel = append(neueArtikel, neu)
		}
	} else {
		for i := range History {
			if History[i].Action == "Insert" {
				geliefert = append(geliefert, History[i].Id)
			}
		}
		for i := range Sage {
			var found bool
			found = false
			for y := 0; y < len(geliefert); y++ {
				if Sage[i].Id == geliefert[y] {
					prod := Warenlieferung{
						ID:   int32(Sage[i].Id),
						Name: Sage[i].Suchbegriff,
					}
					gelieferteArtikel = append(gelieferteArtikel, prod)
				}
			}
			for x := 0; x < len(Products); x++ {
				if Sage[i].Id == int(Products[x].ID) {
					found = true
					break
				}
			}
			if !found {
				neu := Warenlieferung{
					ID:            int32(Sage[i].Id),
					Artikelnummer: Sage[i].Artikelnummer,
					Name:          Sage[i].Suchbegriff,
				}
				neueArtikel = append(neueArtikel, neu)
			}
		}
		for i := range Prices {
			var temp Warenlieferung
			var found bool
			idx := 0
			if len(neuePreise) > 0 {
				for x := 0; x < len(neuePreise); x++ {
					if neuePreise[x].ID == int32(Prices[i].Id) {
						found = true
						temp = neuePreise[x]
						idx = x
					}
				}
			}
			if !found {
				temp.ID = int32(Prices[i].Id)
				temp.Preis = sql.NullTime{Time: time.Now(), Valid: true}
			}
			if Prices[i].Action == "Insert" {
				temp.Neuerpreis = sql.NullFloat64{
					Float64: Prices[i].Price,
					Valid:   true,
				}
			}
			if Prices[i].Action == "Delete" {
				temp.Alterpreis = sql.NullFloat64{
					Float64: Prices[i].Price,
					Valid:   true,
				}
			}
			if idx > 0 {
				var altFloat float64
				var neuFloat float64
				if temp.Alterpreis.Valid {
					altFloat = temp.Alterpreis.Float64
				}
				if altFloat > 0 {
					neuePreise[idx].Alterpreis = temp.Alterpreis
				}
				if temp.Neuerpreis.Valid {
					neuFloat = temp.Neuerpreis.Float64
				}
				if neuFloat > 0 {
					neuePreise[idx].Neuerpreis = temp.Neuerpreis
				}
			} else {
				neuePreise = append(neuePreise, temp)
			}
		}
	}

	return neueArtikel, gelieferteArtikel, neuePreise, nil
}
func getAllProductsFromSage() ([]Artikel, error) {
	flag.Parse()

	var artikel []Artikel

	connString := getSageConnectionString()

	conn, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT sg_auf_artikel.SG_AUF_ARTIKEL_PK, sg_auf_artikel.ARTNR, sg_auf_artikel.SUCHBEGRIFF, sg_auf_vkpreis.PR01 FROM sg_auf_artikel INNER JOIN sg_auf_vkpreis ON (sg_auf_artikel.SG_AUF_ARTIKEL_PK = sg_auf_vkpreis.SG_AUF_ARTIKEL_FK)")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var art Artikel
		var Artikelnummer sql.NullString
		var Suchbegriff sql.NullString
		var Price sql.NullFloat64

		if err := rows.Scan(&art.Id, &Artikelnummer, &Suchbegriff, &Price); err != nil {
			return nil,
				err
		}
		if Artikelnummer.Valid {
			art.Artikelnummer = Artikelnummer.String
		}
		if Suchbegriff.Valid {
			art.Suchbegriff = Suchbegriff.String
		}
		if Price.Valid {
			art.Preis = Price.Float64
		}
		if Suchbegriff.Valid && Artikelnummer.Valid {
			artikel = append(artikel, art)
		}

	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return artikel, nil
}

func getLagerHistory() ([]History, error) {
	var history []History

	queryString := fmt.Sprintf("SELECT SG_AUF_ARTIKEL_FK, Hist_Action FROM sg_auf_lager_history WHERE BEWEGUNG >= 0 AND BEMERKUNG LIKE 'Warenlieferung:%%' AND convert(varchar, Hist_Datetime, 105) = convert(varchar, getdate(), 105)")

	connString := getSageConnectionString()

	conn, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var hist History
		var Action sql.NullString

		if err := rows.Scan(&hist.Id, &Action); err != nil {
			return nil, err
		}
		if Action.Valid {
			hist.Action = Action.String
			history = append(history, hist)
		}
	}

	return history, nil
}

func getPrices() ([]Price, error) {
	var prices []Price

	queryString := "SELECT Hist_Action, SG_AUF_ARTIKEL_FK, PR01 FROM sg_auf_vkpreis_history WHERE convert(varchar, Hist_Datetime, 105) = convert(varchar, getdate(), 105)"

	connString := getSageConnectionString()

	conn, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var price Price
		var Action sql.NullString
		var p sql.NullFloat64

		if err := rows.Scan(&Action, &price.Id, &p); err != nil {
			return nil, fmt.Errorf("GetPrices: Row Error: %s", err)
		}
		if p.Valid {
			price.Price = p.Float64
		}
		if Action.Valid {
			price.Action = Action.String
		}
		if p.Valid && Action.Valid {
			prices = append(prices, price)
		}
	}

	return prices, nil
}

func getLeichen() ([]Leichen, error) {
	var artikel []Leichen

	conn, err := sql.Open("sqlserver", getSageConnectionString())
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	rows, err := conn.Query("SELECT TOP 20 ARTNR, SUCHBEGRIFF, BESTAND, VERFUEGBAR, LetzterUmsatz, EKPR01 FROM sg_auf_artikel WHERE VERFUEGBAR > 0 ORDER BY LetzterUmsatz ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var art Leichen
		var Artikelnummer sql.NullString
		var Artikelname sql.NullString
		var Bestand sql.NullInt16
		var Verfügbar sql.NullInt16
		var EK sql.NullFloat64
		var LetzerUmsatz sql.NullString

		if err := rows.Scan(&Artikelnummer, &Artikelname, &Bestand, &Verfügbar, &LetzerUmsatz, &EK); err != nil {
			return nil,
				err
		}
		if Artikelnummer.Valid {
			art.Artikelnummer = Artikelnummer.String
		}
		if Artikelname.Valid {
			art.Artikelname = Artikelname.String
		}
		if Bestand.Valid {
			art.Bestand = Bestand.Int16
		}
		if EK.Valid {
			art.EK = EK.Float64
		}
		if Verfügbar.Valid {
			art.Verfügbar = Verfügbar.Int16
		}
		if LetzerUmsatz.Valid {
			art.LetzterUmsatz = LetzerUmsatz.String
		}

		if Artikelnummer.Valid && Artikelname.Valid && Bestand.Valid && EK.Valid && LetzerUmsatz.Valid && Verfügbar.Valid {
			artikel = append(artikel, art)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return artikel, nil
}

func getHighestVerfSum() ([]VerfArtikel, error) {
	var artikel []VerfArtikel
	database, err := sql.Open("sqlserver", getSageConnectionString())
	if err != nil {
		return nil, err
	}
	defer database.Close()
	rows, err := database.Query("SELECT TOP 10 ARTNR, SUCHBEGRIFF, BESTAND, VERFUEGBAR, EKPR01, VERFUEGBAR * EKPR01 as Summe FROM sg_auf_artikel WHERE VERFUEGBAR > 0 ORDER BY Summe DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var art VerfArtikel
		var Artikelnummer sql.NullString
		var Artikelname sql.NullString
		var Bestand sql.NullInt16
		var Verfügbar sql.NullInt16
		var EK sql.NullFloat64
		var Summe sql.NullFloat64

		if err := rows.Scan(&Artikelnummer, &Artikelname, &Bestand, &Verfügbar, &EK, &Summe); err != nil {
			return nil,
				err
		}
		if Artikelnummer.Valid {
			art.Artikelnummer = Artikelnummer.String
		}
		if Artikelname.Valid {
			art.Artikelname = Artikelname.String
		}
		if Bestand.Valid {
			art.Bestand = Bestand.Int16
		}
		if EK.Valid {
			art.EK = EK.Float64
		}
		if Summe.Valid {
			art.Summe = Summe.Float64
		}
		if Verfügbar.Valid {
			art.Verfügbar = Verfügbar.Int16
		}

		if Artikelnummer.Valid && Artikelname.Valid && Bestand.Valid && EK.Valid && Summe.Valid && Verfügbar.Valid {
			artikel = append(artikel, art)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return artikel, nil
}

func getHighestSum() ([]SummenArtikel, error) {
	var artikel []SummenArtikel

	conn, err := sql.Open("sqlserver", getSageConnectionString())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	rows, err := conn.Query("SELECT TOP 10 ARTNR, SUCHBEGRIFF, BESTAND, EKPR01, BESTAND * EKPR01 as Summe FROM sg_auf_artikel WHERE BESTAND > 0 ORDER BY Summe DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var art SummenArtikel
		var Artikelnummer sql.NullString
		var Artikelname sql.NullString
		var Bestand sql.NullInt16
		var EK sql.NullFloat64
		var Summe sql.NullFloat64

		if err := rows.Scan(&Artikelnummer, &Artikelname, &Bestand, &EK, &Summe); err != nil {
			return nil,
				err
		}
		if Artikelnummer.Valid {
			art.Artikelnummer = Artikelnummer.String
		}
		if Artikelname.Valid {
			art.Artikelname = Artikelname.String
		}
		if Bestand.Valid {
			art.Bestand = Bestand.Int16
		}
		if EK.Valid {
			art.EK = EK.Float64
		}
		if Summe.Valid {
			art.Summe = Summe.Float64
		}

		if Artikelnummer.Valid && Artikelname.Valid && Bestand.Valid && EK.Valid && Summe.Valid {
			artikel = append(artikel, art)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return artikel, nil
}

func getLagerWert() (float64, float64, error) {
	var wertBestand float64
	var wertVerfügbar float64
	wertBestand = 0
	wertVerfügbar = 0

	conn, err := sql.Open("sqlserver", getSageConnectionString())
	if err != nil {
		return 0, 0, err
	}
	defer conn.Close()
	rows, err := conn.Query("SELECT BESTAND, VERFUEGBAR, EKPR01 FROM sg_auf_artikel WHERE BESTAND > 0")
	if err != nil {
		return 0, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var Bestand sql.NullInt16
		var Verfügbar sql.NullInt16
		var Ek sql.NullFloat64

		if err := rows.Scan(&Bestand, &Verfügbar, &Ek); err != nil {
			return 0, 0,
				err
		}
		if Bestand.Valid && Ek.Valid {
			wertBestand = wertBestand + (float64(Bestand.Int16) * Ek.Float64)
		}
		if Verfügbar.Valid && Ek.Valid {
			wertVerfügbar = wertVerfügbar + (float64(Verfügbar.Int16) * Ek.Float64)
		}
	}
	if err := rows.Err(); err != nil {
		return 0, 0, err
	}

	return wertBestand, wertVerfügbar, nil
}

func getAlteSeriennummern() ([]AlteSeriennummer, error) {
	var artikel []AlteSeriennummer

	conn, err := sql.Open("sqlserver", getSageConnectionString())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	rows, err := conn.Query("SELECT sg_auf_artikel.ARTNR, sg_auf_artikel.SUCHBEGRIFF, sg_auf_artikel.BESTAND, sg_auf_artikel.VERFUEGBAR, sg_auf_snr.GE_Beginn FROM sg_auf_artikel INNER JOIN sg_auf_snr ON sg_auf_artikel.SG_AUF_ARTIKEL_PK = sg_auf_snr.SG_AUF_ARTIKEL_FK  WHERE sg_auf_artikel.VERFUEGBAR > 0 AND sg_auf_snr.SNR_STATUS != 2 AND sg_auf_snr.GE_Beginn <= DATEADD(month, DATEDIFF(month, 0, DATEADD(MONTH,-1,GETDATE())), 0) ORDER BY sg_auf_snr.GE_Beginn ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var art AlteSeriennummer
		var Artikelnummer sql.NullString
		var Suchbegriff sql.NullString
		var Bestand sql.NullInt16
		var Verfügbar sql.NullInt16
		var Garantie sql.NullString

		if err := rows.Scan(&Artikelnummer, &Suchbegriff, &Bestand, &Verfügbar, &Garantie); err != nil {
			return nil,
				err
		}

		if Artikelnummer.Valid && Suchbegriff.Valid && Bestand.Valid && Verfügbar.Valid && Garantie.Valid {
			art.ArtNr = Artikelnummer.String
			art.Suchbegriff = Suchbegriff.String
			art.Bestand = int(Bestand.Int16)
			art.Verfügbar = int(Verfügbar.Int16)
			art.GeBeginn = Garantie.String
			if !slices.Contains(artikel, art) {
				artikel = append(artikel, art)
			}
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return artikel, nil
}

func If[T any](cond bool, vtrue, vfalse T) T {
	if cond {
		return vtrue
	}
	return vfalse
}

func getOldAuftraege() ([]Response, float64, error) {
	var verbrecher []Response
	var gesamtWert float64 = 0
	conn, err := sql.Open("sqlserver", getSageConnectionString())
	if err != nil {
		return nil, 0, err
	}
	defer conn.Close()
	rows, err := conn.Query("SELECT ERFART, USERNAME, DATUM, AUFNR, ALTAUFNR, ENDPRB, VERTRETER, NAME FROM sg_auf_fschrift WHERE FORTGEFUEHRT != 1 AND (ERFART like '02AU' OR ERFART like '03LI');")
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	var res []sg_auf_fschrift
	for rows.Next() {
		var x sg_auf_fschrift
		if err := rows.Scan(&x.ERFART, &x.USERNAME, &x.DATUM, &x.AUFNR, &x.ALTAUFNR, &x.ENDPRB, &x.VERTRETER, &x.NAME); err != nil {
			return nil, 0, err
		}
		res = append(res, x)
	}

	for _, x := range res {
		if x.AUFNR.Valid {
			query := fmt.Sprintf("SELECT ERFART FROM sg_auf_fschrift WHERE ALTAUFNR LIKE '%s';", x.AUFNR.String)
			row, err := conn.Query(query)
			if err != nil {
				return nil, 0, err
			}
			defer row.Close()
			if row.Next() {
				var y sql.NullString
				if err := row.Scan(&y); err != nil {
					return nil, 0, err
				}
				if y.Valid {
					if y.String != "04RE" {
						if x.ENDPRB.Valid {
							verbrecher = append(verbrecher, Response{
								Auftrag:    x.AUFNR.String,
								Wert:       x.ENDPRB.Float64,
								Verbrecher: x.VERTRETER.String,
								Datum:      x.DATUM.Time,
								Kunde:      x.NAME.String,
							})
							gesamtWert += x.ENDPRB.Float64
						}
					}
				}
			} else {
				if x.ENDPRB.Valid {
					verbrecher = append(verbrecher, Response{
						Auftrag:    x.AUFNR.String,
						Wert:       x.ENDPRB.Float64,
						Verbrecher: x.VERTRETER.String,
						Datum:      x.DATUM.Time,
						Kunde:      x.NAME.String,
					})
					gesamtWert += x.ENDPRB.Float64
				}
			}
		}
	}

	// Verbrecher sortieren nach Datum
	sort.Slice(verbrecher, func(i, j int) bool {
		return verbrecher[i].Datum.Before(verbrecher[j].Datum)
	})

	return verbrecher, gesamtWert, nil
}
