package main

import (
	"embed"
	"reflect"
	"strconv"
	"strings"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/build
var assets embed.FS

//go:embed .env
var envStr string

type Env struct {
	DATABASE_URL    string
	SAGE_URL        string
	ACCESS_DB       string
	SMTP_HOST       string
	SMTP_PORT       int
	SMTP_USER       string
	SMTP_PASS       string
	SMTP_FROM       string
	ArchivePath     string
	FTP_UPLOAD_PATH string
	FTP_SERVER      string
	FTP_USER        string
	FTP_PASS        string
	FTP_PORT        int
	AUTH_USER       string
	AUTH_PASS       string
}

func main() {
	_instance := &Env{}
	v := reflect.ValueOf(_instance).Elem()

	lines := strings.Split(envStr, "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}
		line = strings.SplitN(line, "#", 2)[0]
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		f := v.FieldByName(key)
		if f.IsValid() {
			if f.CanInt() {
				val, _ := strconv.Atoi(value)
				f.SetInt(int64(val))
			} else {
				f.SetString(value)
			}
		}
	}

	// Create an instance of the app structure
	app := NewApp(_instance)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Viktor",
		Width:  1500,
		Height: 800,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}

}
