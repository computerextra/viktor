package main

import (
	"os"
	"path/filepath"
	"strings"
	"viktor/archive"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (a App) SearchArchive(searchTerm string) (results []archive.ArchiveResult) {
	results, err := a.archive.Search(searchTerm)
	if err != nil {
		return nil
	}
	return results
}

func (a App) DownloadArchive(id int32) bool {
	res, err := a.archive.Get(id)
	if err != nil {
		return false
	}

	directory := filepath.Join(a.config.Folder.Archive, strings.Replace(res.Title, ":", ".", 1))
	file, err := os.ReadFile(directory)
	if err != nil {
		return false
	}
	pathName, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		DefaultFilename: res.Title,
	})
	if err != nil {
		return false
	}
	if len(pathName) == 0 {
		return false
	}
	err = os.WriteFile(pathName, file, 0644)
	return err == nil
}
