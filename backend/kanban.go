package backend

import (
	"fmt"
	"viktor/db"

	"github.com/wailsapp/wails/v3/pkg/application"
)

func (a *App) CreateKanban(id string, Name string) bool {
	_, err := a.DB.CreateKanban(Name, id)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[CreateKanban] Fehler beim Anlegen", err))
		dialog.Show()
		return false
	}
	return true
}

func (a *App) CreatePost(Name, Status, Importance string, Description *string, KId string) bool {
	_, err := a.DB.CreatePost(KId, Name, Description, Status, Importance)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[CreatePost] Fehler beim Anlegen", err))
		dialog.Show()
		return false
	}
	return true
}

func (a *App) GetKanbanBord(id string) *db.Kanban {
	board, err := a.DB.GetKanban(id)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[GetKanbanBord] Fehler", err))
		dialog.Show()
		return nil
	}
	return board
}

func (a *App) GetKanbanBoardsFromUser(id string) []db.Kanban {
	boards, err := a.DB.GetAllKanbans(id)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[GetKanbanBoardsFromUser] Fehler", err))
		dialog.Show()
		return nil
	}
	return boards
}

func (a *App) GetPostsFromBoard(id string) []db.Post {
	posts, err := a.DB.GetAllPostFromBoard(id)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[GetPostsFromBoard] Fehler", err))
		dialog.Show()
		return nil
	}
	return posts
}

func (a *App) UpdateKanban(id, Name string) bool {
	err := a.DB.UpdateKanban(id, Name)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[UpdateKanban] Fehler", err))
		dialog.Show()
		return false
	}
	return true
}

func (a *App) UpdatePost(id, Name, Status, Importance string, Description *string) bool {
	err := a.DB.UpdatePost(id, Name, Description, Status, Importance)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[UpdatePost] Fehler beim Speichern", err))
		dialog.Show()
		return false
	}
	return true
}

func (a *App) DeletePost(id string) bool {
	err := a.DB.DeletePost(id)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[DeletePost] Fehler", err))
		dialog.Show()
		return false
	}
	return true
}

func (a *App) DeleteBoard(id string) bool {
	err := a.DB.DeleteKanban(id)
	if err != nil {
		dialog := application.ErrorDialog()
		dialog.SetTitle("FEHLER!")
		dialog.SetMessage(fmt.Sprintf("%s: %s", "[DeleteBoard] Fehler", err))
		dialog.Show()
		return false
	}
	return true
}
