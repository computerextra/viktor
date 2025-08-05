package main

import (
	"sort"
	"time"
	"viktor/ent"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Mitarbeiter struct {
	*ent.Mitarbeiter
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

	mitarbeiter, err := a.db.Mitarbeiter.Query().All(a.ctx)
	if err != nil {
		runtime.LogErrorf(a.ctx, err.Error())
		return nil
	}
	var Geburtstag Geburtstag

	// Sort Mitarbeiter
	for _, ma := range mitarbeiter {

		geb := time.Date(time.Now().Year(), ma.Geburtstag.Month(), ma.Geburtstag.Day(), 0, 0, 0, 0, loc)
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

	sort.Slice(Geburtstag.Zukunft, func(i, j int) bool {
		return Geburtstag.Zukunft[i].Geburtstag.Before(Geburtstag.Zukunft[j].Geburtstag)
	})
	sort.Slice(Geburtstag.Vergangenheit, func(i, j int) bool {
		return Geburtstag.Vergangenheit[i].Geburtstag.Before(Geburtstag.Vergangenheit[j].Geburtstag)
	})
	sort.Slice(Geburtstag.Heute, func(i, j int) bool {
		return Geburtstag.Heute[i].Geburtstag.Before(Geburtstag.Heute[j].Geburtstag)
	})

	return &Geburtstag
}
