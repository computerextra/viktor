package backend

import (
	"fmt"
	"viktor/archive"
	"viktor/config"
	"viktor/db"
	"viktor/sagedb"
	"viktor/userdata"
)

type App struct {
	DB       *db.Database
	archive  *archive.Archive
	sage     *sagedb.SageDB
	config   *config.Config
	userdata *userdata.UserData
}

func NewApp(config *config.Config) *App {
	return &App{
		DB:       db.NewDatabase(fmt.Sprintf("%s/viktor_test.db", config.Folder.Upload)),
		archive:  archive.NewArchive(config.Database.Url),
		sage:     sagedb.NewSage(config.Sage.Server, config.Sage.Database, config.Sage.User, config.Sage.Password, config.Sage.Port),
		config:   config,
		userdata: userdata.NewUserdata(),
	}
}
