package backend

import (
	"fmt"
	"viktor/sagedb"

	"github.com/wailsapp/wails/v3/pkg/application"
)

func (a *App) SearchSage(search string) []sagedb.SearchResult {
	if len(search) == 0 {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage("[SearchSage] Kein Suchbegriff angegeben")
		// dialog.SetMessage(fmt.Sprintf("%s: %s", "[CreateAnsprechpartner] [GET(Ansprechpartner)] Fehler", err))
		dialog.Show()
		return nil
	}
	res, err := a.sage.Search(search)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[SearchSage] Fehler", err))
		dialog.Show()
		return nil
	}
	return res
}

func (a *App) GetKundeWithKundennummer(Kundennummer string) *sagedb.User {
	u, e := a.sage.Get(Kundennummer)
	if e != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[GetKundeWithKundennummer] Fehler", e))
		dialog.Show()
		return nil
	}
	return u
}
