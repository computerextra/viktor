package main

import (
	"viktor/sagedb"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a *App) SearchSage(searchTerm string) (results []sagedb.SearchResult) {
	if len(searchTerm) == 0 {
		return nil
	}
	results, err := a.sage.Search(searchTerm)
	if err != nil {
		runtime.LogDebug(a.ctx, err.Error())
		return nil
	}
	return results
}

func (a *App) GetKundeWithKundennummer(kundennummer string) (user *sagedb.User) {
	user, err := a.sage.Get(kundennummer)
	if err != nil {
		return nil
	}
	return user
}
