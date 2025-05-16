package main

// TODO: Datenbank abfragen gehen alle nicht!!!!!!!

import (
	"fmt"
	"viktor/db"
)

func (a *App) Create(model db.Model, params any) bool {
	err := a.db.Create(model, params)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (a *App) Update(model db.Model, params any, idString *string, idInt *int) bool {
	return a.db.Update(model, params, idString, idInt) == nil
}

func (a *App) Read(model db.Model, idString *string, idInt *int) (results any) {
	results, err := a.db.Read(model, idString, idInt)
	if err != nil {
		return nil
	}
	return results
}

func (a *App) Einkaufsliste() []db.MitarbeiterModel {
	results, err := a.db.GetEinkaufsListe()
	if err != nil {
		return nil
	}
	return results
}

func (a *App) Geburtstagsliste() *db.GeburtstagsListe {
	result, err := a.db.GetGeburtstagsListe()
	if err != nil {
		return nil
	}

	return result
}

func (a *App) Delete(model db.Model, idString *string, idInt *int) bool {
	return a.db.Delete(model, idString, idInt) == nil
}
