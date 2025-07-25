package main

import (
	"sort"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Mitarbeiter struct {
	db.Mitarbeiter
	Diff int
}

type Geburtstag struct {
	Heute         []Mitarbeiter
	Zukunft       []Mitarbeiter
	Vergangenheit []Mitarbeiter
}

func (a *App) GetAllMitarbeiter() *Geburtstag {
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		runtime.LogErrorf(a.ctx, err.Error())
		return nil
	}
	today := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, loc)

	mitarbeiter, err := a.db.GetAllMitarbeiter(a.ctx)
	if err != nil {
		runtime.LogErrorf(a.ctx, err.Error())
		return nil
	}
	var Geburtstag Geburtstag

	// Sort Mitarbeiter
	for _, ma := range mitarbeiter {
		if ma.Geburtstag.Valid {
			geb := time.Date(time.Now().Year(), ma.Geburtstag.Time.Month(), ma.Geburtstag.Time.Day(), 0, 0, 0, 0, loc)
			diff := today.Sub(geb)
			days := diff.Hours() / 24
			if days > 0 {
				Geburtstag.Zukunft = append(Geburtstag.Zukunft, Mitarbeiter{
					ma,
					int(days),
				})
			} else if days < 0 {
				Geburtstag.Vergangenheit = append(Geburtstag.Vergangenheit, Mitarbeiter{
					ma,
					int(days) * -1,
				})
			} else {
				Geburtstag.Heute = append(Geburtstag.Heute, Mitarbeiter{
					ma,
					0,
				})
			}
		}

	}

	sort.Slice(Geburtstag.Zukunft, func(i, j int) bool {
		return Geburtstag.Zukunft[i].Geburtstag.Time.Before(Geburtstag.Zukunft[j].Geburtstag.Time)
	})
	sort.Slice(Geburtstag.Vergangenheit, func(i, j int) bool {
		return Geburtstag.Vergangenheit[i].Geburtstag.Time.Before(Geburtstag.Vergangenheit[j].Geburtstag.Time)
	})
	sort.Slice(Geburtstag.Heute, func(i, j int) bool {
		return Geburtstag.Heute[i].Geburtstag.Time.Before(Geburtstag.Heute[j].Geburtstag.Time)
	})

	return &Geburtstag
}
