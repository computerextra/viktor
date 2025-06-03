package backend

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"viktor/archive"

	"github.com/wailsapp/wails/v3/pkg/application"
)

func (a *App) SearchArchive(search string) []archive.ArchiveResult {
	res, err := a.archive.Search(search)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[SearchArchive] Fehler beim Suchen", err))
		dialog.Show()
		return nil
	}
	return res
}

func (a *App) Get(id int32) bool {
	res, err := a.archive.Get(id)
	if err != nil {
		errorDialog := application.ErrorDialog()
		errorDialog.SetTitle("FEHLER!")
		errorDialog.SetMessage(fmt.Sprintf("%s: %s", "[Get] Dokument nicht gefunden", err))
		errorDialog.Show()
		return false
	}
	dialog := application.SaveFileDialog()
	dialog.SetFilename(res.Title)

	path, err := dialog.PromptForSingleSelection()
	if err != nil {
		errorDialog := application.ErrorDialog()
		errorDialog.SetTitle("FEHLER!")
		errorDialog.SetMessage(fmt.Sprintf("%s: %s", "[Get] Pfad nicht gefunden", err))
		errorDialog.Show()
		return false
	}
	directory := filepath.Join(a.config.Folder.Archive, strings.Replace(res.Title, ":", ".", 1))
	file, err := os.ReadFile(directory)
	if err != nil {
		errorDialog := application.ErrorDialog()
		errorDialog.SetTitle("FEHLER!")
		errorDialog.SetMessage(fmt.Sprintf("%s: %s", "[Get] Datei konnte nicht gelesen werden", err))
		errorDialog.Show()
		return false
	}
	err = os.WriteFile(path, file, 0o644)
	if err != nil {
		errorDialog := application.ErrorDialog()
		errorDialog.SetTitle("FEHLER!")
		errorDialog.SetMessage(fmt.Sprintf("%s: %s", "[Get] Datei konnte nicht gespeichert werden", err))
		errorDialog.Show()
		return false
	}
	return true
}
